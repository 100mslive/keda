package scalers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/go-logr/logr"
	"github.com/redis/go-redis/v9"
	v2 "k8s.io/api/autoscaling/v2"
	"k8s.io/metrics/pkg/apis/external_metrics"

	"github.com/kedacore/keda/v2/pkg/util"
)

type redisHmsScaler struct {
	metricType     v2.MetricTargetType
	metadata       *redisHmsMetadata
	closeFn        func() error
	calculateValue func(context.Context) (int64, error)
	logger         logr.Logger
}

type redisHmsMetadata struct {
	mapValue           int64
	activationMapValue int64
	mapName            string
	databaseIndex      int
	connectionInfo     redisConnectionInfo
	triggerIndex       int
}

// NewRedisHMSScaler creates a new redisHMSScaler
func NewRedisHMSScaler(ctx context.Context, config *ScalerConfig) (Scaler, error) {
	luaScript := `
		local mapName = KEYS[1]

		-- Retrieve the map from Redis
		local map = redis.call(hgetall, mapName)
		
		-- Convert the flat list returned by hgetall to a proper table
		local mapTable = {}
		for i = 1, #map, 2 do
			mapTable[map[i]] = tonumber(map[i+1])
		end
	
		-- Ensure the map has the necessary keys
		if mapTable.current == nil or mapTable.desired == nil then
			return nil, "Error: The map must contain 'current' and 'desired' keys."
		end
	
		-- Calculate the difference
		local difference = mapTable.desired - mapTable.current
	
		-- Return the difference
		return difference
	`

	metricType, err := GetMetricTargetType(config)
	if err != nil {
		return nil, fmt.Errorf("error getting scaler metric type: %w", err)
	}

	logger := InitializeLogger(config, "redis_hms_scaler")

	isCluster := config.TriggerMetadata["cluster"] == "true"
	isSentinel := config.TriggerMetadata["sentinel"] == "true"
	if isCluster {
		meta, err := parseRedisHmsMetadata(config, parseRedisClusterAddress)
		if err != nil {
			return nil, fmt.Errorf("error parsing redis metadata: %w", err)
		}
		return createClusteredRedisHmsScaler(ctx, meta, luaScript, metricType, logger)
	} else if isSentinel {
		meta, err := parseRedisHmsMetadata(config, parseRedisSentinelAddress)
		if err != nil {
			return nil, fmt.Errorf("error parsing redis metadata: %w", err)
		}
		return createSentinelRedisHmsScaler(ctx, meta, luaScript, metricType, logger)
	}

	meta, err := parseRedisHmsMetadata(config, parseRedisAddress)
	if err != nil {
		return nil, fmt.Errorf("error parsing redis metadata: %w", err)
	}

	return createRedisHmsScaler(ctx, meta, luaScript, metricType, logger)
}

func createClusteredRedisHmsScaler(ctx context.Context, meta *redisHmsMetadata, script string, metricType v2.MetricTargetType, logger logr.Logger) (Scaler, error) {
	client, err := getRedisClusterClient(ctx, meta.connectionInfo)
	if err != nil {
		return nil, fmt.Errorf("connection to redis cluster failed: %w", err)
	}

	closeFn := func() error {
		if err := client.Close(); err != nil {
			logger.Error(err, "error closing redis client")
			return err
		}
		return nil
	}

	calcFn := func(ctx context.Context) (int64, error) {
		cmd := client.Eval(ctx, script, []string{meta.mapName})
		if cmd.Err() != nil {
			return -1, cmd.Err()
		}

		return cmd.Int64()
	}

	return &redisHmsScaler{
		metricType:     metricType,
		metadata:       meta,
		closeFn:        closeFn,
		calculateValue: calcFn,
		logger:         logger,
	}, nil
}

func createSentinelRedisHmsScaler(ctx context.Context, meta *redisHmsMetadata, script string, metricType v2.MetricTargetType, logger logr.Logger) (Scaler, error) {
	client, err := getRedisSentinelClient(ctx, meta.connectionInfo, meta.databaseIndex)
	if err != nil {
		return nil, fmt.Errorf("connection to redis sentinel failed: %w", err)
	}

	return createRedisHmsScalerWithClient(client, meta, script, metricType, logger), nil
}

func createRedisHmsScaler(ctx context.Context, meta *redisHmsMetadata, script string, metricType v2.MetricTargetType, logger logr.Logger) (Scaler, error) {
	client, err := getRedisClient(ctx, meta.connectionInfo, meta.databaseIndex)
	if err != nil {
		return nil, fmt.Errorf("connection to redis failed: %w", err)
	}

	return createRedisHmsScalerWithClient(client, meta, script, metricType, logger), nil
}

func createRedisHmsScalerWithClient(client *redis.Client, meta *redisHmsMetadata, script string, metricType v2.MetricTargetType, logger logr.Logger) Scaler {
	closeFn := func() error {
		if err := client.Close(); err != nil {
			logger.Error(err, "error closing redis client")
			return err
		}
		return nil
	}

	calcValFn := func(ctx context.Context) (int64, error) {
		cmd := client.Eval(ctx, script, []string{meta.mapName})
		if cmd.Err() != nil {
			return -1, cmd.Err()
		}

		return cmd.Int64()
	}

	return &redisHmsScaler{
		metricType:     metricType,
		metadata:       meta,
		closeFn:        closeFn,
		calculateValue: calcValFn,
		logger:         logger,
	}
}

func parseRedisHmsMetadata(config *scalersconfig.ScalerConfig, parserFn redisAddressParser) (*redisHmsMetadata, error) {
	connInfo, err := parserFn(config.TriggerMetadata, config.ResolvedEnv, config.AuthParams)
	if err != nil {
		return nil, err
	}
	meta := redisHmsMetadata{
		connectionInfo: connInfo,
	}

	err = parseTLSConfigIntoConnectionInfo(config, &meta.connectionInfo)
	if err != nil {
		return nil, err
	}

	meta.mapValue = defaultListLength
	if val, ok := config.TriggerMetadata["mapValue"]; ok {
		v, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("mapValue parsing error: %w", err)
		}
		meta.mapValue = v
	}

	meta.activationMapValue = defaultActivationListLength
	if val, ok := config.TriggerMetadata["activationMapValue"]; ok {
		v, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("activationMapValue parsing error %w", err)
		}
		meta.activationMapValue = v
	}

	if val, ok := config.TriggerMetadata["mapName"]; ok {
		meta.mapName = val
	} else {
		return nil, ErrRedisNoListName
	}

	meta.databaseIndex = defaultDBIdx
	if val, ok := config.TriggerMetadata["databaseIndex"]; ok {
		dbIndex, err := strconv.ParseInt(val, 10, 32)
		if err != nil {
			return nil, fmt.Errorf("databaseIndex: parsing error %w", err)
		}
		meta.databaseIndex = int(dbIndex)
	}
	meta.triggerIndex = config.TriggerIndex
	return &meta, nil
}

func (s *redisHmsScaler) Close(context.Context) error {
	return s.closeFn()
}

// GetMetricSpecForScaling returns the metric spec for the HPA
func (s *redisHmsScaler) GetMetricSpecForScaling(context.Context) []v2.MetricSpec {
	metricName := util.NormalizeString(fmt.Sprintf("redis-hms-%s", s.metadata.mapName))
	externalMetric := &v2.ExternalMetricSource{
		Metric: v2.MetricIdentifier{
			Name: GenerateMetricNameWithIndex(s.metadata.triggerIndex, metricName),
		},
		Target: GetMetricTarget(s.metricType, s.metadata.mapValue),
	}
	metricSpec := v2.MetricSpec{
		External: externalMetric, Type: externalMetricType,
	}
	return []v2.MetricSpec{metricSpec}
}

// GetMetricsAndActivity connects to Redis and finds the length of the list
func (s *redisHmsScaler) GetMetricsAndActivity(ctx context.Context, metricName string) ([]external_metrics.ExternalMetricValue, bool, error) {
	diff, err := s.calculateValue(ctx)

	if err != nil {
		s.logger.Error(err, "error getting calculated value")
		return []external_metrics.ExternalMetricValue{}, false, err
	}

	metric := GenerateMetricInMili(metricName, float64(diff))

	return []external_metrics.ExternalMetricValue{metric}, diff > 0, nil
}

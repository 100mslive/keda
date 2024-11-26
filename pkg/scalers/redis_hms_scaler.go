package scalers

import (
	"context"
	"fmt"

	"github.com/go-logr/logr"
	"github.com/redis/go-redis/v9"
	v2 "k8s.io/api/autoscaling/v2"
	"k8s.io/metrics/pkg/apis/external_metrics"

	"github.com/kedacore/keda/v2/pkg/scalers/scalersconfig"
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
	MapValue           int64               `keda:"name=mapValue,       order=triggerMetadata, optional, default=5"`
	ActivationMapValue int64               `keda:"name=activationMapValue,       order=triggerMetadata, optional"`
	MapName            string              `keda:"name=mapName,       order=triggerMetadata"`
	DatabaseIndex      int                 `keda:"name=databaseIndex,       order=triggerMetadata, optional"`
	ConnectionInfo     redisConnectionInfo `keda:"optional"`
	triggerIndex       int
}

// NewRedisHMSScaler creates a new redisHMSScaler
func NewRedisHMSScaler(ctx context.Context, config *scalersconfig.ScalerConfig) (Scaler, error) {
	luaScript := `
local mapName = KEYS[1]

-- Retrieve the map from Redis
local map = redis.call('HGETALL', mapName)

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

-- If the difference is positive, return it
if difference > 0 then
  return difference
end

-- Return the difference
return 0
	`

	metricType, err := GetMetricTargetType(config)
	if err != nil {
		return nil, fmt.Errorf("error getting scaler metric type: %w", err)
	}

	logger := InitializeLogger(config, "redis_hms_scaler")

	meta, err := parseRedisHmsMetadata(config)
	if err != nil {
		return nil, fmt.Errorf("error parsing redis metadata: %w", err)
	}

	isCluster := config.TriggerMetadata["cluster"] == "true"
	isSentinel := config.TriggerMetadata["sentinel"] == "true"
	if isCluster {
		return createClusteredRedisHmsScaler(ctx, meta, luaScript, metricType, logger)
	} else if isSentinel {
		return createSentinelRedisHmsScaler(ctx, meta, luaScript, metricType, logger)
	}

	return createRedisHmsScaler(ctx, meta, luaScript, metricType, logger)
}

func createClusteredRedisHmsScaler(ctx context.Context, meta *redisHmsMetadata, script string, metricType v2.MetricTargetType, logger logr.Logger) (Scaler, error) {
	client, err := getRedisClusterClient(ctx, meta.ConnectionInfo)
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
		cmd := client.Eval(ctx, script, []string{meta.MapName})
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
	client, err := getRedisSentinelClient(ctx, meta.ConnectionInfo, meta.DatabaseIndex)
	if err != nil {
		return nil, fmt.Errorf("connection to redis sentinel failed: %w", err)
	}

	return createRedisHmsScalerWithClient(client, meta, script, metricType, logger), nil
}

func createRedisHmsScaler(ctx context.Context, meta *redisHmsMetadata, script string, metricType v2.MetricTargetType, logger logr.Logger) (Scaler, error) {
	client, err := getRedisClient(ctx, meta.ConnectionInfo, meta.DatabaseIndex)
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
		cmd := client.Eval(ctx, script, []string{meta.MapName})
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

func parseRedisHmsMetadata(config *scalersconfig.ScalerConfig) (*redisHmsMetadata, error) {
	meta := &redisHmsMetadata{}
	if err := config.TypedConfig(meta); err != nil {
		return nil, fmt.Errorf("error parsing redis metadata: %w", err)
	}

	meta.triggerIndex = config.TriggerIndex
	return meta, nil
}

func (s *redisHmsScaler) Close(context.Context) error {
	return s.closeFn()
}

// GetMetricSpecForScaling returns the metric spec for the HPA
func (s *redisHmsScaler) GetMetricSpecForScaling(context.Context) []v2.MetricSpec {
	metricName := util.NormalizeString(fmt.Sprintf("redis-hms-%s", s.metadata.MapName))
	externalMetric := &v2.ExternalMetricSource{
		Metric: v2.MetricIdentifier{
			Name: GenerateMetricNameWithIndex(s.metadata.triggerIndex, metricName),
		},
		Target: GetMetricTarget(s.metricType, s.metadata.MapValue),
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

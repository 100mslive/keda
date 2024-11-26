// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatch

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an anomaly detection model for a CloudWatch metric. You can use the
// model to display a band of expected normal values when the metric is graphed.
//
// If you have enabled unified cross-account observability, and this account is a
// monitoring account, the metric can be in the same account or a source account.
// You can specify the account ID in the object you specify in the
// SingleMetricAnomalyDetector parameter.
//
// For more information, see [CloudWatch Anomaly Detection].
//
// [CloudWatch Anomaly Detection]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Anomaly_Detection.html
func (c *Client) PutAnomalyDetector(ctx context.Context, params *PutAnomalyDetectorInput, optFns ...func(*Options)) (*PutAnomalyDetectorOutput, error) {
	if params == nil {
		params = &PutAnomalyDetectorInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutAnomalyDetector", params, optFns, c.addOperationPutAnomalyDetectorMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutAnomalyDetectorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutAnomalyDetectorInput struct {

	// The configuration specifies details about how the anomaly detection model is to
	// be trained, including time ranges to exclude when training and updating the
	// model. You can specify as many as 10 time ranges.
	//
	// The configuration can also include the time zone to use for the metric.
	Configuration *types.AnomalyDetectorConfiguration

	// The metric dimensions to create the anomaly detection model for.
	//
	// Deprecated: Use SingleMetricAnomalyDetector.
	Dimensions []types.Dimension

	// Use this object to include parameters to provide information about your metric
	// to CloudWatch to help it build more accurate anomaly detection models.
	// Currently, it includes the PeriodicSpikes parameter.
	MetricCharacteristics *types.MetricCharacteristics

	// The metric math anomaly detector to be created.
	//
	// When using MetricMathAnomalyDetector , you cannot include the following
	// parameters in the same operation:
	//
	//   - Dimensions
	//
	//   - MetricName
	//
	//   - Namespace
	//
	//   - Stat
	//
	//   - the SingleMetricAnomalyDetector parameters of PutAnomalyDetectorInput
	//
	// Instead, specify the metric math anomaly detector attributes as part of the
	// property MetricMathAnomalyDetector .
	MetricMathAnomalyDetector *types.MetricMathAnomalyDetector

	// The name of the metric to create the anomaly detection model for.
	//
	// Deprecated: Use SingleMetricAnomalyDetector.
	MetricName *string

	// The namespace of the metric to create the anomaly detection model for.
	//
	// Deprecated: Use SingleMetricAnomalyDetector.
	Namespace *string

	// A single metric anomaly detector to be created.
	//
	// When using SingleMetricAnomalyDetector , you cannot include the following
	// parameters in the same operation:
	//
	//   - Dimensions
	//
	//   - MetricName
	//
	//   - Namespace
	//
	//   - Stat
	//
	//   - the MetricMathAnomalyDetector parameters of PutAnomalyDetectorInput
	//
	// Instead, specify the single metric anomaly detector attributes as part of the
	// property SingleMetricAnomalyDetector .
	SingleMetricAnomalyDetector *types.SingleMetricAnomalyDetector

	// The statistic to use for the metric and the anomaly detection model.
	//
	// Deprecated: Use SingleMetricAnomalyDetector.
	Stat *string

	noSmithyDocumentSerde
}

type PutAnomalyDetectorOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutAnomalyDetectorMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpPutAnomalyDetector{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpPutAnomalyDetector{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutAnomalyDetector"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addSpanRetryLoop(stack, options); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addOpPutAnomalyDetectorValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutAnomalyDetector(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	if err = addSpanInitializeStart(stack); err != nil {
		return err
	}
	if err = addSpanInitializeEnd(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestStart(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestEnd(stack); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opPutAnomalyDetector(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutAnomalyDetector",
	}
}

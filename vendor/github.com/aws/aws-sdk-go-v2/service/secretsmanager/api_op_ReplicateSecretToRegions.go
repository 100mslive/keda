// Code generated by smithy-go-codegen DO NOT EDIT.

package secretsmanager

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Replicates the secret to a new Regions. See [Multi-Region secrets].
//
// Secrets Manager generates a CloudTrail log entry when you call this action. Do
// not include sensitive information in request parameters because it might be
// logged. For more information, see [Logging Secrets Manager events with CloudTrail].
//
// Required permissions: secretsmanager:ReplicateSecretToRegions . If the primary
// secret is encrypted with a KMS key other than aws/secretsmanager , you also need
// kms:Decrypt permission to the key. To encrypt the replicated secret with a KMS
// key other than aws/secretsmanager , you need kms:GenerateDataKey and kms:Encrypt
// to the key. For more information, see [IAM policy actions for Secrets Manager]and [Authentication and access control in Secrets Manager].
//
// [Authentication and access control in Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access.html
// [Logging Secrets Manager events with CloudTrail]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/retrieve-ct-entries.html
// [Multi-Region secrets]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/create-manage-multi-region-secrets.html
// [IAM policy actions for Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/reference_iam-permissions.html#reference_iam-permissions_actions
func (c *Client) ReplicateSecretToRegions(ctx context.Context, params *ReplicateSecretToRegionsInput, optFns ...func(*Options)) (*ReplicateSecretToRegionsOutput, error) {
	if params == nil {
		params = &ReplicateSecretToRegionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ReplicateSecretToRegions", params, optFns, c.addOperationReplicateSecretToRegionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ReplicateSecretToRegionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ReplicateSecretToRegionsInput struct {

	// A list of Regions in which to replicate the secret.
	//
	// This member is required.
	AddReplicaRegions []types.ReplicaRegionType

	// The ARN or name of the secret to replicate.
	//
	// This member is required.
	SecretId *string

	// Specifies whether to overwrite a secret with the same name in the destination
	// Region. By default, secrets aren't overwritten.
	ForceOverwriteReplicaSecret bool

	noSmithyDocumentSerde
}

type ReplicateSecretToRegionsOutput struct {

	// The ARN of the primary secret.
	ARN *string

	// The status of replication.
	ReplicationStatus []types.ReplicationStatusType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationReplicateSecretToRegionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpReplicateSecretToRegions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpReplicateSecretToRegions{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ReplicateSecretToRegions"); err != nil {
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
	if err = addOpReplicateSecretToRegionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opReplicateSecretToRegions(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opReplicateSecretToRegions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ReplicateSecretToRegions",
	}
}

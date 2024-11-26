// Code generated by smithy-go-codegen DO NOT EDIT.

package secretsmanager

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Attaches a resource-based permission policy to a secret. A resource-based
// policy is optional. For more information, see [Authentication and access control for Secrets Manager]
//
// For information about attaching a policy in the console, see [Attach a permissions policy to a secret].
//
// Secrets Manager generates a CloudTrail log entry when you call this action. Do
// not include sensitive information in request parameters because it might be
// logged. For more information, see [Logging Secrets Manager events with CloudTrail].
//
// Required permissions: secretsmanager:PutResourcePolicy . For more information,
// see [IAM policy actions for Secrets Manager]and [Authentication and access control in Secrets Manager].
//
// [Authentication and access control in Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access.html
// [Logging Secrets Manager events with CloudTrail]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/retrieve-ct-entries.html
// [Attach a permissions policy to a secret]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access_resource-based-policies.html
// [IAM policy actions for Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/reference_iam-permissions.html#reference_iam-permissions_actions
// [Authentication and access control for Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access.html
func (c *Client) PutResourcePolicy(ctx context.Context, params *PutResourcePolicyInput, optFns ...func(*Options)) (*PutResourcePolicyOutput, error) {
	if params == nil {
		params = &PutResourcePolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutResourcePolicy", params, optFns, c.addOperationPutResourcePolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutResourcePolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutResourcePolicyInput struct {

	// A JSON-formatted string for an Amazon Web Services resource-based policy. For
	// example policies, see [Permissions policy examples].
	//
	// [Permissions policy examples]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access_examples.html
	//
	// This member is required.
	ResourcePolicy *string

	// The ARN or name of the secret to attach the resource-based policy.
	//
	// For an ARN, we recommend that you specify a complete ARN rather than a partial
	// ARN. See [Finding a secret from a partial ARN].
	//
	// [Finding a secret from a partial ARN]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/troubleshoot.html#ARN_secretnamehyphen
	//
	// This member is required.
	SecretId *string

	// Specifies whether to block resource-based policies that allow broad access to
	// the secret, for example those that use a wildcard for the principal. By default,
	// public policies aren't blocked.
	//
	// Resource policy validation and the BlockPublicPolicy parameter help protect
	// your resources by preventing public access from being granted through the
	// resource policies that are directly attached to your secrets. In addition to
	// using these features, carefully inspect the following policies to confirm that
	// they do not grant public access:
	//
	//   - Identity-based policies attached to associated Amazon Web Services
	//   principals (for example, IAM roles)
	//
	//   - Resource-based policies attached to associated Amazon Web Services
	//   resources (for example, Key Management Service (KMS) keys)
	//
	// To review permissions to your secrets, see [Determine who has permissions to your secrets].
	//
	// [Determine who has permissions to your secrets]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/determine-acccess_examine-iam-policies.html
	BlockPublicPolicy *bool

	noSmithyDocumentSerde
}

type PutResourcePolicyOutput struct {

	// The ARN of the secret.
	ARN *string

	// The name of the secret.
	Name *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutResourcePolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutResourcePolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutResourcePolicy{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutResourcePolicy"); err != nil {
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
	if err = addOpPutResourcePolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutResourcePolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutResourcePolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutResourcePolicy",
	}
}

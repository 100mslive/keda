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

// Creates a new secret. A secret can be a password, a set of credentials such as
// a user name and password, an OAuth token, or other secret information that you
// store in an encrypted form in Secrets Manager. The secret also includes the
// connection information to access a database or other service, which Secrets
// Manager doesn't encrypt. A secret in Secrets Manager consists of both the
// protected secret data and the important information needed to manage the secret.
//
// For secrets that use managed rotation, you need to create the secret through
// the managing service. For more information, see [Secrets Manager secrets managed by other Amazon Web Services services].
//
// For information about creating a secret in the console, see [Create a secret].
//
// To create a secret, you can provide the secret value to be encrypted in either
// the SecretString parameter or the SecretBinary parameter, but not both. If you
// include SecretString or SecretBinary then Secrets Manager creates an initial
// secret version and automatically attaches the staging label AWSCURRENT to it.
//
// For database credentials you want to rotate, for Secrets Manager to be able to
// rotate the secret, you must make sure the JSON you store in the SecretString
// matches the [JSON structure of a database secret].
//
// If you don't specify an KMS encryption key, Secrets Manager uses the Amazon Web
// Services managed key aws/secretsmanager . If this key doesn't already exist in
// your account, then Secrets Manager creates it for you automatically. All users
// and roles in the Amazon Web Services account automatically have access to use
// aws/secretsmanager . Creating aws/secretsmanager can result in a one-time
// significant delay in returning the result.
//
// If the secret is in a different Amazon Web Services account from the
// credentials calling the API, then you can't use aws/secretsmanager to encrypt
// the secret, and you must create and use a customer managed KMS key.
//
// Secrets Manager generates a CloudTrail log entry when you call this action. Do
// not include sensitive information in request parameters except SecretBinary or
// SecretString because it might be logged. For more information, see [Logging Secrets Manager events with CloudTrail].
//
// Required permissions: secretsmanager:CreateSecret . If you include tags in the
// secret, you also need secretsmanager:TagResource . To add replica Regions, you
// must also have secretsmanager:ReplicateSecretToRegions . For more information,
// see [IAM policy actions for Secrets Manager]and [Authentication and access control in Secrets Manager].
//
// To encrypt the secret with a KMS key other than aws/secretsmanager , you need
// kms:GenerateDataKey and kms:Decrypt permission to the key.
//
// When you enter commands in a command shell, there is a risk of the command
// history being accessed or utilities having access to your command parameters.
// This is a concern if the command includes the value of a secret. Learn how to [Mitigate the risks of using command-line tools to store Secrets Manager secrets].
//
// [Authentication and access control in Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access.html
// [Logging Secrets Manager events with CloudTrail]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/retrieve-ct-entries.html
// [Secrets Manager secrets managed by other Amazon Web Services services]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/service-linked-secrets.html
// [Mitigate the risks of using command-line tools to store Secrets Manager secrets]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/security_cli-exposure-risks.html
// [Create a secret]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/manage_create-basic-secret.html
// [IAM policy actions for Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/reference_iam-permissions.html#reference_iam-permissions_actions
// [JSON structure of a database secret]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/reference_secret_json_structure.html
func (c *Client) CreateSecret(ctx context.Context, params *CreateSecretInput, optFns ...func(*Options)) (*CreateSecretOutput, error) {
	if params == nil {
		params = &CreateSecretInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateSecret", params, optFns, c.addOperationCreateSecretMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateSecretOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateSecretInput struct {

	// The name of the new secret.
	//
	// The secret name can contain ASCII letters, numbers, and the following
	// characters: /_+=.@-
	//
	// Do not end your secret name with a hyphen followed by six characters. If you do
	// so, you risk confusion and unexpected results when searching for a secret by
	// partial ARN. Secrets Manager automatically adds a hyphen and six random
	// characters after the secret name at the end of the ARN.
	//
	// This member is required.
	Name *string

	// A list of Regions and KMS keys to replicate secrets.
	AddReplicaRegions []types.ReplicaRegionType

	// If you include SecretString or SecretBinary , then Secrets Manager creates an
	// initial version for the secret, and this parameter specifies the unique
	// identifier for the new version.
	//
	// If you use the Amazon Web Services CLI or one of the Amazon Web Services SDKs
	// to call this operation, then you can leave this parameter empty. The CLI or SDK
	// generates a random UUID for you and includes it as the value for this parameter
	// in the request.
	//
	// If you generate a raw HTTP request to the Secrets Manager service endpoint,
	// then you must generate a ClientRequestToken and include it in the request.
	//
	// This value helps ensure idempotency. Secrets Manager uses this value to prevent
	// the accidental creation of duplicate versions if there are failures and retries
	// during a rotation. We recommend that you generate a [UUID-type]value to ensure uniqueness
	// of your versions within the specified secret.
	//
	//   - If the ClientRequestToken value isn't already associated with a version of
	//   the secret then a new version of the secret is created.
	//
	//   - If a version with this value already exists and the version SecretString and
	//   SecretBinary values are the same as those in the request, then the request is
	//   ignored.
	//
	//   - If a version with this value already exists and that version's SecretString
	//   and SecretBinary values are different from those in the request, then the
	//   request fails because you cannot modify an existing version. Instead, use PutSecretValueto
	//   create a new version.
	//
	// This value becomes the VersionId of the new version.
	//
	// [UUID-type]: https://wikipedia.org/wiki/Universally_unique_identifier
	ClientRequestToken *string

	// The description of the secret.
	Description *string

	// Specifies whether to overwrite a secret with the same name in the destination
	// Region. By default, secrets aren't overwritten.
	ForceOverwriteReplicaSecret bool

	// The ARN, key ID, or alias of the KMS key that Secrets Manager uses to encrypt
	// the secret value in the secret. An alias is always prefixed by alias/ , for
	// example alias/aws/secretsmanager . For more information, see [About aliases].
	//
	// To use a KMS key in a different account, use the key ARN or the alias ARN.
	//
	// If you don't specify this value, then Secrets Manager uses the key
	// aws/secretsmanager . If that key doesn't yet exist, then Secrets Manager creates
	// it for you automatically the first time it encrypts the secret value.
	//
	// If the secret is in a different Amazon Web Services account from the
	// credentials calling the API, then you can't use aws/secretsmanager to encrypt
	// the secret, and you must create and use a customer managed KMS key.
	//
	// [About aliases]: https://docs.aws.amazon.com/kms/latest/developerguide/alias-about.html
	KmsKeyId *string

	// The binary data to encrypt and store in the new version of the secret. We
	// recommend that you store your binary data in a file and then pass the contents
	// of the file as a parameter.
	//
	// Either SecretString or SecretBinary must have a value, but not both.
	//
	// This parameter is not available in the Secrets Manager console.
	//
	// Sensitive: This field contains sensitive information, so the service does not
	// include it in CloudTrail log entries. If you create your own log entries, you
	// must also avoid logging the information in this field.
	SecretBinary []byte

	// The text data to encrypt and store in this new version of the secret. We
	// recommend you use a JSON structure of key/value pairs for your secret value.
	//
	// Either SecretString or SecretBinary must have a value, but not both.
	//
	// If you create a secret by using the Secrets Manager console then Secrets
	// Manager puts the protected secret text in only the SecretString parameter. The
	// Secrets Manager console stores the information as a JSON structure of key/value
	// pairs that a Lambda rotation function can parse.
	//
	// Sensitive: This field contains sensitive information, so the service does not
	// include it in CloudTrail log entries. If you create your own log entries, you
	// must also avoid logging the information in this field.
	SecretString *string

	// A list of tags to attach to the secret. Each tag is a key and value pair of
	// strings in a JSON text string, for example:
	//
	//     [{"Key":"CostCenter","Value":"12345"},{"Key":"environment","Value":"production"}]
	//
	// Secrets Manager tag key names are case sensitive. A tag with the key "ABC" is a
	// different tag from one with key "abc".
	//
	// If you check tags in permissions policies as part of your security strategy,
	// then adding or removing a tag can change permissions. If the completion of this
	// operation would result in you losing your permissions for this secret, then
	// Secrets Manager blocks the operation and returns an Access Denied error. For
	// more information, see [Control access to secrets using tags]and [Limit access to identities with tags that match secrets' tags].
	//
	// For information about how to format a JSON parameter for the various command
	// line tool environments, see [Using JSON for Parameters]. If your command-line tool or SDK requires
	// quotation marks around the parameter, you should use single quotes to avoid
	// confusion with the double quotes required in the JSON text.
	//
	// For tag quotas and naming restrictions, see [Service quotas for Tagging] in the Amazon Web Services General
	// Reference guide.
	//
	// [Limit access to identities with tags that match secrets' tags]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access_examples.html#auth-and-access_tags2
	// [Using JSON for Parameters]: https://docs.aws.amazon.com/cli/latest/userguide/cli-using-param.html#cli-using-param-json
	// [Service quotas for Tagging]: https://docs.aws.amazon.com/general/latest/gr/arg.html#taged-reference-quotas
	// [Control access to secrets using tags]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access_examples.html#tag-secrets-abac
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateSecretOutput struct {

	// The ARN of the new secret. The ARN includes the name of the secret followed by
	// six random characters. This ensures that if you create a new secret with the
	// same name as a deleted secret, then users with access to the old secret don't
	// get access to the new secret because the ARNs are different.
	ARN *string

	// The name of the new secret.
	Name *string

	// A list of the replicas of this secret and their status:
	//
	//   - Failed , which indicates that the replica was not created.
	//
	//   - InProgress , which indicates that Secrets Manager is in the process of
	//   creating the replica.
	//
	//   - InSync , which indicates that the replica was created.
	ReplicationStatus []types.ReplicationStatusType

	// The unique identifier associated with the version of the new secret.
	VersionId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateSecretMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateSecret{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateSecret{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateSecret"); err != nil {
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
	if err = addIdempotencyToken_opCreateSecretMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateSecretValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateSecret(options.Region), middleware.Before); err != nil {
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
	return nil
}

type idempotencyToken_initializeOpCreateSecret struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateSecret) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateSecret) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateSecretInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateSecretInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateSecretMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateSecret{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateSecret(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateSecret",
	}
}

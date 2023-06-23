// Code generated by smithy-go-codegen DO NOT EDIT.

package verifiedpermissions

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieve the details for the specified policy template in the specified policy
// store.
func (c *Client) GetPolicyTemplate(ctx context.Context, params *GetPolicyTemplateInput, optFns ...func(*Options)) (*GetPolicyTemplateOutput, error) {
	if params == nil {
		params = &GetPolicyTemplateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetPolicyTemplate", params, optFns, c.addOperationGetPolicyTemplateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetPolicyTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetPolicyTemplateInput struct {

	// Specifies the ID of the policy store that contains the policy template that you
	// want information about.
	//
	// This member is required.
	PolicyStoreId *string

	// Specifies the ID of the policy template that you want information about.
	//
	// This member is required.
	PolicyTemplateId *string

	noSmithyDocumentSerde
}

type GetPolicyTemplateOutput struct {

	// The date and time that the policy template was originally created.
	//
	// This member is required.
	CreatedDate *time.Time

	// The date and time that the policy template was most recently updated.
	//
	// This member is required.
	LastUpdatedDate *time.Time

	// The ID of the policy store that contains the policy template.
	//
	// This member is required.
	PolicyStoreId *string

	// The ID of the policy template.
	//
	// This member is required.
	PolicyTemplateId *string

	// The content of the body of the policy template written in the Cedar policy
	// language.
	//
	// This member is required.
	Statement *string

	// The description of the policy template.
	Description *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetPolicyTemplateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpGetPolicyTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpGetPolicyTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpGetPolicyTemplateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetPolicyTemplate(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
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
	return nil
}

func newServiceMetadataMiddleware_opGetPolicyTemplate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "verifiedpermissions",
		OperationName: "GetPolicyTemplate",
	}
}
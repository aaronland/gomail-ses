// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the individual task executions (one per target) for a particular task
// run as part of a maintenance window execution.
func (c *Client) DescribeMaintenanceWindowExecutionTaskInvocations(ctx context.Context, params *DescribeMaintenanceWindowExecutionTaskInvocationsInput, optFns ...func(*Options)) (*DescribeMaintenanceWindowExecutionTaskInvocationsOutput, error) {
	if params == nil {
		params = &DescribeMaintenanceWindowExecutionTaskInvocationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeMaintenanceWindowExecutionTaskInvocations", params, optFns, c.addOperationDescribeMaintenanceWindowExecutionTaskInvocationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeMaintenanceWindowExecutionTaskInvocationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeMaintenanceWindowExecutionTaskInvocationsInput struct {

	// The ID of the specific task in the maintenance window task that should be
	// retrieved.
	//
	// This member is required.
	TaskId *string

	// The ID of the maintenance window execution the task is part of.
	//
	// This member is required.
	WindowExecutionId *string

	// Optional filters used to scope down the returned task invocations. The
	// supported filter key is STATUS with the corresponding values PENDING ,
	// IN_PROGRESS , SUCCESS , FAILED , TIMED_OUT , CANCELLING , and CANCELLED .
	Filters []types.MaintenanceWindowFilter

	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	MaxResults *int32

	// The token for the next set of items to return. (You received this token from a
	// previous call.)
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeMaintenanceWindowExecutionTaskInvocationsOutput struct {

	// The token to use when requesting the next set of items. If there are no
	// additional items to return, the string is empty.
	NextToken *string

	// Information about the task invocation results per invocation.
	WindowExecutionTaskInvocationIdentities []types.MaintenanceWindowExecutionTaskInvocationIdentity

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeMaintenanceWindowExecutionTaskInvocationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeMaintenanceWindowExecutionTaskInvocations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeMaintenanceWindowExecutionTaskInvocations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeMaintenanceWindowExecutionTaskInvocations"); err != nil {
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
	if err = addOpDescribeMaintenanceWindowExecutionTaskInvocationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeMaintenanceWindowExecutionTaskInvocations(options.Region), middleware.Before); err != nil {
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

// DescribeMaintenanceWindowExecutionTaskInvocationsPaginatorOptions is the
// paginator options for DescribeMaintenanceWindowExecutionTaskInvocations
type DescribeMaintenanceWindowExecutionTaskInvocationsPaginatorOptions struct {
	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeMaintenanceWindowExecutionTaskInvocationsPaginator is a paginator for
// DescribeMaintenanceWindowExecutionTaskInvocations
type DescribeMaintenanceWindowExecutionTaskInvocationsPaginator struct {
	options   DescribeMaintenanceWindowExecutionTaskInvocationsPaginatorOptions
	client    DescribeMaintenanceWindowExecutionTaskInvocationsAPIClient
	params    *DescribeMaintenanceWindowExecutionTaskInvocationsInput
	nextToken *string
	firstPage bool
}

// NewDescribeMaintenanceWindowExecutionTaskInvocationsPaginator returns a new
// DescribeMaintenanceWindowExecutionTaskInvocationsPaginator
func NewDescribeMaintenanceWindowExecutionTaskInvocationsPaginator(client DescribeMaintenanceWindowExecutionTaskInvocationsAPIClient, params *DescribeMaintenanceWindowExecutionTaskInvocationsInput, optFns ...func(*DescribeMaintenanceWindowExecutionTaskInvocationsPaginatorOptions)) *DescribeMaintenanceWindowExecutionTaskInvocationsPaginator {
	if params == nil {
		params = &DescribeMaintenanceWindowExecutionTaskInvocationsInput{}
	}

	options := DescribeMaintenanceWindowExecutionTaskInvocationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeMaintenanceWindowExecutionTaskInvocationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeMaintenanceWindowExecutionTaskInvocationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeMaintenanceWindowExecutionTaskInvocations
// page.
func (p *DescribeMaintenanceWindowExecutionTaskInvocationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeMaintenanceWindowExecutionTaskInvocationsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.DescribeMaintenanceWindowExecutionTaskInvocations(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

// DescribeMaintenanceWindowExecutionTaskInvocationsAPIClient is a client that
// implements the DescribeMaintenanceWindowExecutionTaskInvocations operation.
type DescribeMaintenanceWindowExecutionTaskInvocationsAPIClient interface {
	DescribeMaintenanceWindowExecutionTaskInvocations(context.Context, *DescribeMaintenanceWindowExecutionTaskInvocationsInput, ...func(*Options)) (*DescribeMaintenanceWindowExecutionTaskInvocationsOutput, error)
}

var _ DescribeMaintenanceWindowExecutionTaskInvocationsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opDescribeMaintenanceWindowExecutionTaskInvocations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeMaintenanceWindowExecutionTaskInvocations",
	}
}

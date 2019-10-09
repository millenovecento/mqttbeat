// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/ec2query"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DeleteTagsRequest
type DeleteTagsInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// The IDs of the resources, separated by spaces.
	//
	// Constraints: Up to 1000 resource IDs. We recommend breaking up this request
	// into smaller batches.
	//
	// Resources is a required field
	Resources []string `locationName:"resourceId" type:"list" required:"true"`

	// The tags to delete. Specify a tag key and an optional tag value to delete
	// specific tags. If you specify a tag key without a tag value, we delete any
	// tag with this key regardless of its value. If you specify a tag key with
	// an empty string as the tag value, we delete the tag only if its value is
	// an empty string.
	//
	// If you omit this parameter, we delete all user-defined tags for the specified
	// resources. We do not delete AWS-generated tags (tags that have the aws: prefix).
	Tags []Tag `locationName:"tag" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s DeleteTagsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteTagsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteTagsInput"}

	if s.Resources == nil {
		invalidParams.Add(aws.NewErrParamRequired("Resources"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DeleteTagsOutput
type DeleteTagsOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteTagsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteTags = "DeleteTags"

// DeleteTagsRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Deletes the specified set of tags from the specified set of resources.
//
// To list the current tags, use DescribeTags. For more information about tags,
// see Tagging Your Resources (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Using_Tags.html)
// in the Amazon Elastic Compute Cloud User Guide.
//
//    // Example sending a request using DeleteTagsRequest.
//    req := client.DeleteTagsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DeleteTags
func (c *Client) DeleteTagsRequest(input *DeleteTagsInput) DeleteTagsRequest {
	op := &aws.Operation{
		Name:       opDeleteTags,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteTagsInput{}
	}

	req := c.newRequest(op, input, &DeleteTagsOutput{})
	req.Handlers.Unmarshal.Remove(ec2query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteTagsRequest{Request: req, Input: input, Copy: c.DeleteTagsRequest}
}

// DeleteTagsRequest is the request type for the
// DeleteTags API operation.
type DeleteTagsRequest struct {
	*aws.Request
	Input *DeleteTagsInput
	Copy  func(*DeleteTagsInput) DeleteTagsRequest
}

// Send marshals and sends the DeleteTags API request.
func (r DeleteTagsRequest) Send(ctx context.Context) (*DeleteTagsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteTagsResponse{
		DeleteTagsOutput: r.Request.Data.(*DeleteTagsOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteTagsResponse is the response type for the
// DeleteTags API operation.
type DeleteTagsResponse struct {
	*DeleteTagsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteTags request.
func (r *DeleteTagsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

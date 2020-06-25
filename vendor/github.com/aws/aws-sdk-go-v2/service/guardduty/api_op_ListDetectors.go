// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package guardduty

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/guardduty-2017-11-28/ListDetectorsRequest
type ListDetectorsInput struct {
	_ struct{} `type:"structure"`

	// You can use this parameter to indicate the maximum number of items you want
	// in the response. The default value is 50. The maximum value is 50.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	// You can use this parameter when paginating results. Set the value of this
	// parameter to null on your first call to the list action. For subsequent calls
	// to the action fill nextToken in the request with the value of NextToken from
	// the previous response to continue listing data.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListDetectorsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListDetectorsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListDetectorsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListDetectorsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "maxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/guardduty-2017-11-28/ListDetectorsResponse
type ListDetectorsOutput struct {
	_ struct{} `type:"structure"`

	// A list of detector Ids.
	//
	// DetectorIds is a required field
	DetectorIds []string `locationName:"detectorIds" type:"list" required:"true"`

	// Pagination parameter to be used on the next list operation to retrieve more
	// items.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListDetectorsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListDetectorsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.DetectorIds != nil {
		v := s.DetectorIds

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "detectorIds", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opListDetectors = "ListDetectors"

// ListDetectorsRequest returns a request value for making API operation for
// Amazon GuardDuty.
//
// Lists detectorIds of all the existing Amazon GuardDuty detector resources.
//
//    // Example sending a request using ListDetectorsRequest.
//    req := client.ListDetectorsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/guardduty-2017-11-28/ListDetectors
func (c *Client) ListDetectorsRequest(input *ListDetectorsInput) ListDetectorsRequest {
	op := &aws.Operation{
		Name:       opListDetectors,
		HTTPMethod: "GET",
		HTTPPath:   "/detector",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListDetectorsInput{}
	}

	req := c.newRequest(op, input, &ListDetectorsOutput{})
	return ListDetectorsRequest{Request: req, Input: input, Copy: c.ListDetectorsRequest}
}

// ListDetectorsRequest is the request type for the
// ListDetectors API operation.
type ListDetectorsRequest struct {
	*aws.Request
	Input *ListDetectorsInput
	Copy  func(*ListDetectorsInput) ListDetectorsRequest
}

// Send marshals and sends the ListDetectors API request.
func (r ListDetectorsRequest) Send(ctx context.Context) (*ListDetectorsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListDetectorsResponse{
		ListDetectorsOutput: r.Request.Data.(*ListDetectorsOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListDetectorsRequestPaginator returns a paginator for ListDetectors.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListDetectorsRequest(input)
//   p := guardduty.NewListDetectorsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListDetectorsPaginator(req ListDetectorsRequest) ListDetectorsPaginator {
	return ListDetectorsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListDetectorsInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// ListDetectorsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListDetectorsPaginator struct {
	aws.Pager
}

func (p *ListDetectorsPaginator) CurrentPage() *ListDetectorsOutput {
	return p.Pager.CurrentPage().(*ListDetectorsOutput)
}

// ListDetectorsResponse is the response type for the
// ListDetectors API operation.
type ListDetectorsResponse struct {
	*ListDetectorsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListDetectors request.
func (r *ListDetectorsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
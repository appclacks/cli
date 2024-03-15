package client

import (
	"context"

	apitypes "github.com/appclacks/go-types"
)

// ############################################################
// ############################################################
// 			Client Results public methods


func (c *Client) ListHealthchecksResultsWithContext(ctx context.Context, input apitypes.ListHealthchecksResultsInput) (apitypes.ListHealthchecksResultsOutput, error) {
	var result apitypes.ListHealthchecksResultsOutput

	if !c.Credentials.HasAuth() {
		return apitypes.ListHealthchecksResultsOutput{}, formatError("ListHealthchecksResultsWithContext", "Credentials missing")
	}

	if ctx == nil {
		return apitypes.ListHealthchecksResultsOutput{}, formatError("ListHealthchecksResultsWithContext", " 'nil' value for context.Context")
	}

	request, err := c.RequestsHelper.BuildGetResultsRequest(input)
	if err != nil {
		return apitypes.ListHealthchecksResultsOutput{}, err
	}

	request = request.WithContext(ctx)

	return sendRequestGetStruct(c.RequestsHelper, request, result)
}

func (c *Client) ListHealthchecksResults(input apitypes.ListHealthchecksResultsInput) (apitypes.ListHealthchecksResultsOutput, error) {
	var result apitypes.ListHealthchecksResultsOutput
	
	if !c.Credentials.HasAuth() {
		return apitypes.ListHealthchecksResultsOutput{}, formatError("ListHealthchecksResults", "Credentials missing")
	}

	request, err := c.RequestsHelper.BuildGetResultsRequest(input)
	if err != nil {
		return apitypes.ListHealthchecksResultsOutput{}, err
	}

	return sendRequestGetStruct(c.RequestsHelper, request, result)

}


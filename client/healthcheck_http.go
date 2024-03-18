package client

import (
	"context"

	apitypes "github.com/appclacks/go-types"
)

// ############################################################
// ############################################################
// 			Client public HTTPHealthcheck-useCases 

func (c *Client) CreateHTTPHealthcheckWithContext(ctx context.Context, input apitypes.CreateHTTPHealthcheckInput) (apitypes.Healthcheck, error) {
	var result apitypes.Healthcheck
	
	if !c.Credentials.HasAuth() {
		return apitypes.Healthcheck{}, formatError("CreateHTTPHealthcheckWithContext", "Credentials missing")
	}

	if ctx == nil {
		return apitypes.Healthcheck{}, formatError("CreateHTTPHealthcheckWithContext", " 'nil' value for context.Context")
	}

	request, err := c.RequestsHelper.BuildCreateHTTPHealthCheckRequest(input)
	if err != nil {
		return apitypes.Healthcheck{}, err 
	}

	request = request.WithContext(ctx)

	return sendRequestGetStruct(c.RequestsHelper, request, result)

}

func (c *Client) CreateHTTPHealthcheck(input apitypes.CreateHTTPHealthcheckInput) (apitypes.Healthcheck, error) {
	var result apitypes.Healthcheck
	
	if !c.Credentials.HasAuth() {
		return apitypes.Healthcheck{}, formatError("CreateHTTPHealthcheck", "Credentials missing")
	}

	request, err := c.RequestsHelper.BuildCreateHTTPHealthCheckRequest(input)
	if err != nil {
		return apitypes.Healthcheck{}, err 
	}

	return sendRequestGetStruct(c.RequestsHelper, request, result)

}

// -----------------------------------------------------------------------------------------------------------------------


func (c *Client) UpdateHTTPHealthcheckWithContext(ctx context.Context, input apitypes.UpdateHTTPHealthcheckInput) (apitypes.Healthcheck, error) {
	var result apitypes.Healthcheck
	
	if !c.Credentials.HasAuth() {
		return apitypes.Healthcheck{}, formatError("UpdateHTTPHealthcheckWithContext", "Credentials missing")
	}

	if ctx == nil {
		return apitypes.Healthcheck{}, formatError("UpdateHTTPHealthcheckWithContext", " 'nil' value for context.Context")
	}

	request, err := c.RequestsHelper.BuildUpdateHTTPHealthCheckRequest(input)
	if err != nil {
		return apitypes.Healthcheck{}, err 
	}

	request = request.WithContext(ctx)

	return sendRequestGetStruct(c.RequestsHelper, request, result)
}

func (c *Client) UpdateHTTPHealthcheck(input apitypes.UpdateHTTPHealthcheckInput) (apitypes.Healthcheck, error) {
	var result apitypes.Healthcheck
	
	if !c.Credentials.HasAuth() {
		return apitypes.Healthcheck{}, formatError("UpdateHTTPHealthcheck", "Credentials missing")
	}

	request, err := c.RequestsHelper.BuildUpdateHTTPHealthCheckRequest(input)
	if err != nil {
		return apitypes.Healthcheck{}, err 
	}

	return sendRequestGetStruct(c.RequestsHelper, request, result)
}
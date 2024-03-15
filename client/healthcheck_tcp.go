package client

import (
	"context"

	apitypes "github.com/appclacks/go-types"
)

// ############################################################
// ############################################################
// 			Client public TCPHealthcheck-useCases 

func (c *Client) CreateTCPHealthcheckWithContext(ctx context.Context, input apitypes.CreateTCPHealthcheckInput) (apitypes.Healthcheck, error) {
	var result apitypes.Healthcheck
	
	if !c.Credentials.HasAuth() {
		return apitypes.Healthcheck{}, formatError("CreateTCPHealthcheckWithContext", "Credentials missing")
	}

	if ctx == nil {
		return apitypes.Healthcheck{}, formatError("CreateTCPHealthcheckWithContext", " 'nil' value for context.Context")
	}

	request, err := c.RequestsHelper.BuildCreateTCPHealthCheckRequest(input)
	if err != nil {
		return apitypes.Healthcheck{}, err 
	}

	request = request.WithContext(ctx)

	return sendRequestGetStruct(c.RequestsHelper, request, result)

}

func (c *Client) CreateTCPHealthcheck(input apitypes.CreateTCPHealthcheckInput) (apitypes.Healthcheck, error) {
	var result apitypes.Healthcheck
	
	if !c.Credentials.HasAuth() {
		return apitypes.Healthcheck{}, formatError("CreateTCPHealthcheck", "Credentials missing")
	}

	request, err := c.RequestsHelper.BuildCreateTCPHealthCheckRequest(input)
	if err != nil {
		return apitypes.Healthcheck{}, err 
	}

	return sendRequestGetStruct(c.RequestsHelper, request, result)

}

// ----------------------------------------------------------------------------------------------------------------------------

func (c *Client) UpdateTCPHealthcheckWithContext(ctx context.Context, input apitypes.UpdateTCPHealthcheckInput) (apitypes.Healthcheck, error) {
	var result apitypes.Healthcheck
	
	if !c.Credentials.HasAuth() {
		return apitypes.Healthcheck{}, formatError("UpdateTCPHealthcheckWithContext", "Credentials missing")
	}

	if ctx == nil {
		return apitypes.Healthcheck{}, formatError("UpdateTCPHealthcheckWithContext", " 'nil' value for context.Context")
	}

	request, err := c.RequestsHelper.BuildUpdateTCPHealthCheckRequest(input)
	if err != nil {
		return apitypes.Healthcheck{}, err 
	}

	request = request.WithContext(ctx)

	return sendRequestGetStruct(c.RequestsHelper, request, result)

}

func (c *Client) UpdateTCPHealthcheck(input apitypes.UpdateTCPHealthcheckInput) (apitypes.Healthcheck, error) {
	var result apitypes.Healthcheck
	
	if !c.Credentials.HasAuth() {
		return apitypes.Healthcheck{}, formatError("UpdateTCPHealthcheck", "Credentials missing")
	}

	request, err := c.RequestsHelper.BuildUpdateTCPHealthCheckRequest(input)
	if err != nil {
		return apitypes.Healthcheck{}, err 
	}

	return sendRequestGetStruct(c.RequestsHelper, request, result)

}

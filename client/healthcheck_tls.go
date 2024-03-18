package client

import (
	"context"


	apitypes "github.com/appclacks/go-types"
)

// ############################################################
// ############################################################
// 			Client public TLSHealthcheck-useCases 


func (c *Client) CreateTLSHealthcheckWithContext(ctx context.Context, input apitypes.CreateTLSHealthcheckInput) (apitypes.Healthcheck, error) {
	var result apitypes.Healthcheck
	
	if !c.Credentials.HasAuth() {
		return apitypes.Healthcheck{}, formatError("CreateTLSHealthcheckWithContext", "Credentials missing")
	}

	if ctx == nil {
		return apitypes.Healthcheck{}, formatError("CreateTLSHealthcheckWithContext", " 'nil' value for context.Context")
	}

	request, err := c.RequestsHelper.BuildCreateTLSHealthCheckRequest(input)
	if err != nil {
		return apitypes.Healthcheck{}, err 
	}

	request = request.WithContext(ctx)

	return sendRequestGetStruct(c.RequestsHelper, request, result)

}

func (c *Client) CreateTLSHealthcheck(input apitypes.CreateTLSHealthcheckInput) (apitypes.Healthcheck, error) {
	var result apitypes.Healthcheck
	
	if !c.Credentials.HasAuth() {
		return apitypes.Healthcheck{}, formatError("CreateTLSHealthcheck", "Credentials missing")
	}

	request, err := c.RequestsHelper.BuildCreateTLSHealthCheckRequest(input)
	if err != nil {
		return apitypes.Healthcheck{}, err 
	}

	return sendRequestGetStruct(c.RequestsHelper, request, result)

}

// ---------------------------------------------------------------------------------------------------------------------------------

func (c *Client) UpdateTLSHealthcheckWithContext(ctx context.Context, input apitypes.UpdateTLSHealthcheckInput) (apitypes.Healthcheck, error) {
	var result apitypes.Healthcheck
	
	if !c.Credentials.HasAuth() {
		return apitypes.Healthcheck{}, formatError("UpdateTLSHealthcheckWithContext", "Credentials missing")
	}

	if ctx == nil {
		return apitypes.Healthcheck{}, formatError("UpdateTLSHealthcheckWithContext", " 'nil' value for context.Context")
	}

	request, err := c.RequestsHelper.BuildUpdateTLSHealthCheckRequest(input)
	if err != nil {
		return apitypes.Healthcheck{}, err 
	}

	request = request.WithContext(ctx)

	return sendRequestGetStruct(c.RequestsHelper, request, result)
}

func (c *Client) UpdateTLSHealthcheck(input apitypes.UpdateTLSHealthcheckInput) (apitypes.Healthcheck, error) {
	var result apitypes.Healthcheck
	
	if !c.Credentials.HasAuth() {
		return apitypes.Healthcheck{}, formatError("UpdateTLSHealthcheck", "Credentials missing")
	}

	request, err := c.RequestsHelper.BuildUpdateTLSHealthCheckRequest(input)
	if err != nil {
		return apitypes.Healthcheck{}, err 
	}

	return sendRequestGetStruct(c.RequestsHelper, request, result)
}
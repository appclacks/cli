package client

import (
	"context"

	apitypes "github.com/appclacks/go-types"
)

// ############################################################
// ############################################################
// 			Client public Account-useCases 

func (c *Client) CreateCommandHealthcheckWithContext(ctx context.Context, input apitypes.CreateCommandHealthcheckInput) (apitypes.Healthcheck, error){
	var result apitypes.Healthcheck
	
	if !c.Credentials.HasAuth() {
		return apitypes.Healthcheck{}, formatError("CreateCommandHealthcheckWithContext", "Credentials missing")
	}

	if ctx == nil {
		return apitypes.Healthcheck{}, formatError("CreateCommandHealthcheckWithContext", " 'nil' value for context.Context")
	}

	request, err := c.RequestsHelper.BuildCreateCommandHealthCheckRequest(input)
	if err != nil {
		return apitypes.Healthcheck{}, err 
	}

	request = request.WithContext(ctx)

	return sendRequestGetStruct(c.RequestsHelper, request, result)

}

func (c *Client) CreateCommandHealthcheck(input apitypes.CreateCommandHealthcheckInput) (apitypes.Healthcheck, error) {
	var result apitypes.Healthcheck

	if !c.Credentials.HasAuth() {
		return apitypes.Healthcheck{}, formatError("CreateCommandHealthcheck", "Credentials missing")
	}

	request, err := c.RequestsHelper.BuildCreateCommandHealthCheckRequest(input)
	if err != nil {
		return apitypes.Healthcheck{}, err 
	}

	return sendRequestGetStruct(c.RequestsHelper, request, result)
}

// -------------------------------------------------------------------------------------------------------------------------------


func (c *Client) UpdateCommandHealthcheckWithContext(ctx context.Context, input apitypes.UpdateCommandHealthcheckInput) (apitypes.Healthcheck, error) {
	var result apitypes.Healthcheck
	
	if !c.Credentials.HasAuth() {
		return apitypes.Healthcheck{}, formatError("UpdateCommandHealthcheckWithContext", "Credentials missing")
	}

	if ctx == nil {
		return apitypes.Healthcheck{}, formatError("UpdateCommandHealthcheckWithContext", " 'nil' value for context.Context")
	}
	
	request, err := c.RequestsHelper.BuildUpdateCommandHealthCheckRequest(input)
	if err != nil {
		return apitypes.Healthcheck{}, err 
	}

	request = request.WithContext(ctx)

	return sendRequestGetStruct(c.RequestsHelper, request, result)
}

func (c *Client) UpdateCommandHealthcheck(input apitypes.UpdateCommandHealthcheckInput) (apitypes.Healthcheck, error) {
	var result apitypes.Healthcheck

	if !c.Credentials.HasAuth() {
		return apitypes.Healthcheck{}, formatError("UpdateCommandHealthcheck", "Credentials missing")
	}

	request, err := c.RequestsHelper.BuildUpdateCommandHealthCheckRequest(input)
	if err != nil {
		return apitypes.Healthcheck{}, err 
	}


	return sendRequestGetStruct(c.RequestsHelper, request, result)
}

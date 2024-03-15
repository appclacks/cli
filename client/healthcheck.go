package client

import (
	"context"

	apitypes "github.com/appclacks/go-types"
)

// ############################################################
// ############################################################
// 			Client public Healthcheck-useCases 

func (c *Client) DeleteHealthcheckWithContext(ctx context.Context, input apitypes.DeleteHealthcheckInput) (apitypes.Response, error) {
	var result apitypes.Response
	
	if !c.Credentials.HasAuth() {
		return apitypes.Response{}, formatError("DeleteHealthcheckWithContext", "Credentials missing")
	}

	if ctx == nil {
		return apitypes.Response{}, formatError("DeleteHealthcheckWithContext", " 'nil' value for context.Context")
	}

	request, err := c.RequestsHelper.BuildDeleteHealthCheckRequest(input.ID)
	if err != nil {
		return apitypes.Response{}, err
	}

	request = request.WithContext(ctx)

	return sendRequestGetStruct(c.RequestsHelper, request, result)
}

func (c *Client) DeleteHealthcheck(input apitypes.DeleteHealthcheckInput) (apitypes.Response, error) {
	var result apitypes.Response
	
	if !c.Credentials.HasAuth() {
		return apitypes.Response{}, formatError("DeleteHealthcheckWithContext", "Credentials missing")
	}

	request, err := c.RequestsHelper.BuildDeleteHealthCheckRequest(input.ID)
	if err != nil {
		return apitypes.Response{}, err
	}

	return sendRequestGetStruct(c.RequestsHelper, request, result)
}

// -----------------------------------------------------------------------------------------------------

func (c *Client) GetHealthcheckWithContext(ctx context.Context, input apitypes.GetHealthcheckInput) (apitypes.Healthcheck, error) {
	var result apitypes.Healthcheck
	
	if !c.Credentials.HasAuth() {
		return apitypes.Healthcheck{}, formatError("GetHealthcheckWithContext", "Credentials missing")
	}

	if ctx == nil {
		return apitypes.Healthcheck{}, formatError("GetHealthcheckWithContext", " 'nil' value for context.Context")
	}

	request, err := c.RequestsHelper.BuildGetHealthCheckRequest(input.Identifier)
	if err != nil {
		return apitypes.Healthcheck{}, err
	}

	request = request.WithContext(ctx)

	return sendRequestGetStruct(c.RequestsHelper, request, result)
}

func (c *Client) GetHealthcheck(input apitypes.GetHealthcheckInput) (apitypes.Healthcheck, error) {
	var result apitypes.Healthcheck
	
	
	if !c.Credentials.HasAuth() {
		return apitypes.Healthcheck{}, formatError("GetHealthcheck", "Credentials missing")
	}

	request, err := c.RequestsHelper.BuildGetHealthCheckRequest(input.Identifier)
	if err != nil {
		return apitypes.Healthcheck{}, err
	}

	return sendRequestGetStruct(c.RequestsHelper, request, result)
}

// ----------------------------------------------------------------------------------------------------------


func (c *Client) ListHealthchecksWithContext(ctx context.Context) (apitypes.ListHealthchecksOutput, error) {
	var result apitypes.ListHealthchecksOutput
	
	if !c.Credentials.HasAuth() {
		return apitypes.ListHealthchecksOutput{}, formatError("ListHealthchecksWithContext", "Credentials missing")
	}

	if ctx == nil {
		return apitypes.ListHealthchecksOutput{}, formatError("ListHealthchecksWithContext", " 'nil' value for context.Context")
	}

	request, err := c.RequestsHelper.BuildGetListHealthCheckRequest()
	if err != nil {
		return apitypes.ListHealthchecksOutput{}, err
	}

	request = request.WithContext(ctx)

	return sendRequestGetStruct(c.RequestsHelper, request, result)
}

func (c *Client) ListHealthchecks() (apitypes.ListHealthchecksOutput, error) {
	var result apitypes.ListHealthchecksOutput
	
	if !c.Credentials.HasAuth() {
		return apitypes.ListHealthchecksOutput{}, formatError("ListHealthchecks", "Credentials missing")
	}

	request, err := c.RequestsHelper.BuildGetListHealthCheckRequest()
	if err != nil {
		return apitypes.ListHealthchecksOutput{}, err
	}

	return sendRequestGetStruct(c.RequestsHelper, request, result)
}

// --------------------------------------------------------------------------------------------------------------------------

func (c *Client) CabourotteDiscoveryWithContext(ctx context.Context, input apitypes.CabourotteDiscoveryInput) (apitypes.CabourotteDiscoveryOutput, error) {
	var result apitypes.CabourotteDiscoveryOutput
	
	if !c.Credentials.HasAuth() {
		return apitypes.CabourotteDiscoveryOutput{}, formatError("CabourotteDiscoveryWithContext", "Credentials missing")
	}

	if ctx == nil {
		return apitypes.CabourotteDiscoveryOutput{}, formatError("CabourotteDiscoveryWithContext", " 'nil' value for context.Context")
	}

	request, err := c.RequestsHelper.BuildCabourotteDiscoveryRequest(input.Labels)
	if err != nil {
		return apitypes.CabourotteDiscoveryOutput{}, err
	}

	request = request.WithContext(ctx)

	return sendRequestGetStruct(c.RequestsHelper, request, result)
}

func (c *Client) CabourotteDiscovery(input apitypes.CabourotteDiscoveryInput) (apitypes.CabourotteDiscoveryOutput, error) {
	var result apitypes.CabourotteDiscoveryOutput
	
	if !c.Credentials.HasAuth() {
		return apitypes.CabourotteDiscoveryOutput{}, formatError("CabourotteDiscovery", "Credentials missing")
	}

	request, err := c.RequestsHelper.BuildCabourotteDiscoveryRequest(input.Labels)
	if err != nil {
		return apitypes.CabourotteDiscoveryOutput{}, err
	}

	return sendRequestGetStruct(c.RequestsHelper, request, result)
}
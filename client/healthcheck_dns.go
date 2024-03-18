package client

import (
	"context"

	apitypes "github.com/appclacks/go-types"
)

// ############################################################
// ############################################################
// 			Client public HealthcheckDNS-useCases 


func (c *Client) CreateDNSHealthcheckWithContext(ctx context.Context, input apitypes.CreateDNSHealthcheckInput) (apitypes.Healthcheck, error) {
	var result apitypes.Healthcheck
	
	if !c.Credentials.HasAuth() {
		return apitypes.Healthcheck{}, formatError("CreateDNSHealthcheckWithContext", "Credentials missing")
	}

	if ctx == nil {
		return apitypes.Healthcheck{}, formatError("CreateDNSHealthcheckWithContext", " 'nil' value for context.Context")
	}

	request, err := c.RequestsHelper.BuildCreateDNSHealthCheckRequest(input)
	if err != nil {
		return apitypes.Healthcheck{}, err 
	}

	request = request.WithContext(ctx)

	return sendRequestGetStruct(c.RequestsHelper, request, result)

}


func (c *Client) CreateDNSHealthcheck(input apitypes.CreateDNSHealthcheckInput) (apitypes.Healthcheck, error) {
	var result apitypes.Healthcheck
	
	if !c.Credentials.HasAuth() {
		return apitypes.Healthcheck{}, formatError("CreateDNSHealthcheck", "Credentials missing")
	}

	request, err := c.RequestsHelper.BuildCreateDNSHealthCheckRequest(input)
	if err != nil {
		return apitypes.Healthcheck{}, err 
	}

	return sendRequestGetStruct(c.RequestsHelper, request, result)

}

// --------------------------------------------------------------------------------------------------------------------

type internalUpdateHealthcheckInput struct {
	Name        string            `json:"name"`
	Description string            `json:"description,omitempty"`
	Labels      map[string]string `json:"labels,omitempty"`
	Interval    string            `json:"interval"`
	Enabled     bool              `json:"enabled"`
	Timeout     string            `json:"timeout"`
}

func (c *Client) UpdateDNSHealthcheckWithContext(ctx context.Context, input apitypes.UpdateDNSHealthcheckInput) (apitypes.Healthcheck, error) {
	var result apitypes.Healthcheck
	
	if !c.Credentials.HasAuth() {
		return apitypes.Healthcheck{}, formatError("UpdateDNSHealthcheckWithContext", "Credentials missing")
	}

	if ctx == nil {
		return apitypes.Healthcheck{}, formatError("UpdateDNSHealthcheckWithContext", " 'nil' value for context.Context")
	}

	request, err := c.RequestsHelper.BuildUpdateDNSHealthCheckRequest(input)
	if err != nil {
		return apitypes.Healthcheck{}, err 
	}

	request = request.WithContext(ctx)

	return sendRequestGetStruct(c.RequestsHelper, request, result)
}

func (c *Client) UpdateDNSHealthcheck(input apitypes.UpdateDNSHealthcheckInput) (apitypes.Healthcheck, error) {
	var result apitypes.Healthcheck
	
	if !c.Credentials.HasAuth() {
		return apitypes.Healthcheck{}, formatError("UpdateDNSHealthcheck", "Credentials missing")
	}

	request, err := c.RequestsHelper.BuildUpdateDNSHealthCheckRequest(input)
	if err != nil {
		return apitypes.Healthcheck{}, err 
	}

	return sendRequestGetStruct(c.RequestsHelper, request, result)
}

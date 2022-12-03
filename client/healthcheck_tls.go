package client

import (
	"fmt"
	"net/http"

	apitypes "github.com/appclacks/go-types"
)

func (c *Client) CreateTLSHealthcheck(input apitypes.CreateTLSHealthcheckInput) (apitypes.Healthcheck, error) {
	var result apitypes.Healthcheck
	_, err := c.sendRequest("/api/v1/healthcheck/tls", http.MethodPost, input, &result, nil, TokenAuth)
	if err != nil {
		return apitypes.Healthcheck{}, err
	}
	return result, nil
}

func (c *Client) UpdateTLSHealthcheck(input apitypes.UpdateTLSHealthcheckInput) (apitypes.Response, error) {
	var result apitypes.Response
	internalInput := internalUpdateHealthcheckInput{
		Name:        input.Name,
		Description: input.Description,
		Labels:      input.Labels,
		Interval:    input.Interval,
		Enabled:     input.Enabled,
	}
	payload, err := jsonMerge(internalInput, input.HealthcheckTLSDefinition)
	if err != nil {
		return result, err
	}
	_, err = c.sendRequest(fmt.Sprintf("/api/v1/healthcheck/tls/%s", input.ID), http.MethodPut, payload, &result, nil, TokenAuth)
	if err != nil {
		return apitypes.Response{}, err
	}
	return result, nil
}

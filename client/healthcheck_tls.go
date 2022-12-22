package client

import (
	"context"
	"fmt"
	"net/http"

	apitypes "github.com/appclacks/go-types"
)

func (c *Client) CreateTLSHealthcheck(ctx context.Context, input apitypes.CreateTLSHealthcheckInput) (apitypes.Healthcheck, error) {
	var result apitypes.Healthcheck
	_, err := c.sendRequest(ctx, "/api/v1/healthcheck/tls", http.MethodPost, input, &result, nil, TokenAuth)
	if err != nil {
		return apitypes.Healthcheck{}, err
	}
	return result, nil
}

func (c *Client) UpdateTLSHealthcheck(ctx context.Context, input apitypes.UpdateTLSHealthcheckInput) (apitypes.Healthcheck, error) {
	var result apitypes.Healthcheck
	internalInput := internalUpdateHealthcheckInput{
		Name:        input.Name,
		Description: input.Description,
		Labels:      input.Labels,
		Interval:    input.Interval,
		Enabled:     input.Enabled,
		Timeout:     input.Timeout,
	}
	payload, err := jsonMerge(internalInput, input.HealthcheckTLSDefinition)
	if err != nil {
		return result, err
	}
	_, err = c.sendRequest(ctx, fmt.Sprintf("/api/v1/healthcheck/tls/%s", input.ID), http.MethodPut, payload, &result, nil, TokenAuth)
	if err != nil {
		return apitypes.Healthcheck{}, err
	}
	return result, nil
}

package client

import (
	"context"
	"fmt"
	"net/http"

	apitypes "github.com/appclacks/go-types"
)

func (c *Client) CreateTCPHealthcheck(ctx context.Context, input apitypes.CreateTCPHealthcheckInput) (apitypes.Healthcheck, error) {
	var result apitypes.Healthcheck
	_, err := c.sendRequest(ctx, "/api/v1/healthcheck/tcp", http.MethodPost, input, &result, nil, TokenAuth)
	if err != nil {
		return apitypes.Healthcheck{}, err
	}
	return result, nil
}

func (c *Client) UpdateTCPHealthcheck(ctx context.Context, input apitypes.UpdateTCPHealthcheckInput) (apitypes.Healthcheck, error) {
	var result apitypes.Healthcheck
	internalInput := internalUpdateHealthcheckInput{
		Name:        input.Name,
		Description: input.Description,
		Labels:      input.Labels,
		Interval:    input.Interval,
		Enabled:     input.Enabled,
		Timeout:     input.Timeout,
	}
	payload, err := jsonMerge(internalInput, input.HealthcheckTCPDefinition)
	if err != nil {
		return result, err
	}
	_, err = c.sendRequest(ctx, fmt.Sprintf("/api/v1/healthcheck/tcp/%s", input.ID), http.MethodPut, payload, &result, nil, TokenAuth)
	if err != nil {
		return apitypes.Healthcheck{}, err
	}
	return result, nil
}

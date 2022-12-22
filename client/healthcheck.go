package client

import (
	"context"
	"fmt"
	"net/http"

	apitypes "github.com/appclacks/go-types"
)

func (c *Client) DeleteHealthcheck(ctx context.Context, input apitypes.DeleteHealthcheckInput) (apitypes.Response, error) {
	var result apitypes.Response
	_, err := c.sendRequest(ctx, fmt.Sprintf("/api/v1/healthcheck/%s", input.ID), http.MethodDelete, nil, &result, nil, TokenAuth)
	if err != nil {
		return apitypes.Response{}, err
	}
	return result, nil
}

func (c *Client) GetHealthcheck(ctx context.Context, input apitypes.GetHealthcheckInput) (apitypes.Healthcheck, error) {
	var result apitypes.Healthcheck
	_, err := c.sendRequest(ctx, fmt.Sprintf("/api/v1/healthcheck/%s", input.ID), http.MethodGet, nil, &result, nil, TokenAuth)
	if err != nil {
		return apitypes.Healthcheck{}, err
	}
	return result, nil
}

func (c *Client) ListHealthchecks(ctx context.Context) (apitypes.ListHealthchecksOutput, error) {
	var result apitypes.ListHealthchecksOutput
	_, err := c.sendRequest(ctx, "/api/v1/healthcheck", http.MethodGet, nil, &result, nil, TokenAuth)
	if err != nil {
		return apitypes.ListHealthchecksOutput{}, err
	}
	return result, nil
}

func (c *Client) CabourotteDiscovery(ctx context.Context, input apitypes.CabourotteDiscoveryInput) (apitypes.CabourotteDiscoveryOutput, error) {
	var result apitypes.CabourotteDiscoveryOutput
	queryParams := make(map[string]string)
	if input.Labels != "" {
		queryParams["labels"] = input.Labels
	}
	_, err := c.sendRequest(ctx, "/cabourotte/discovery", http.MethodGet, nil, &result, queryParams, TokenAuth)
	if err != nil {
		return apitypes.CabourotteDiscoveryOutput{}, err
	}
	return result, nil
}

package client

import (
	"fmt"
	"net/http"

	apitypes "github.com/appclacks/go-types"
)

func (c *Client) DeleteHealthcheck(input apitypes.DeleteHealthcheckInput) (apitypes.Response, error) {
	var result apitypes.Response
	_, err := c.sendRequest(fmt.Sprintf("/api/v1/healthcheck/%s", input.ID), http.MethodDelete, nil, &result, nil, TokenAuth)
	if err != nil {
		return apitypes.Response{}, err
	}
	return result, nil
}

func (c *Client) GetHealthcheck(input apitypes.GetHealthcheckInput) (apitypes.Healthcheck, error) {
	var result apitypes.Healthcheck
	_, err := c.sendRequest(fmt.Sprintf("/api/v1/healthcheck/%s", input.ID), http.MethodGet, nil, &result, nil, TokenAuth)
	if err != nil {
		return apitypes.Healthcheck{}, err
	}
	return result, nil
}

func (c *Client) ListHealthchecks() (apitypes.ListHealthchecksOutput, error) {
	var result apitypes.ListHealthchecksOutput
	_, err := c.sendRequest("/api/v1/healthcheck", http.MethodGet, nil, &result, nil, TokenAuth)
	if err != nil {
		return apitypes.ListHealthchecksOutput{}, err
	}
	return result, nil
}

func (c *Client) CabourotteDiscovery(input apitypes.CabourotteDiscoveryInput) (apitypes.CabourotteDiscoveryOutput, error) {
	var result apitypes.CabourotteDiscoveryOutput
	queryParams := make(map[string]string)
	if input.Labels != "" {
		queryParams["labels"] = input.Labels
	}
	_, err := c.sendRequest("/cabourotte/discovery", http.MethodGet, nil, &result, queryParams, TokenAuth)
	if err != nil {
		return apitypes.CabourotteDiscoveryOutput{}, err
	}
	return result, nil
}

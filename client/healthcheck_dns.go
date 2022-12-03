package client

import (
	"fmt"
	"net/http"

	apitypes "github.com/appclacks/go-types"
)

func (c *Client) CreateDNSHealthcheck(input apitypes.CreateDNSHealthcheckInput) (apitypes.Healthcheck, error) {
	var result apitypes.Healthcheck
	_, err := c.sendRequest("/api/v1/healthcheck/dns", http.MethodPost, input, &result, nil, TokenAuth)
	if err != nil {
		return apitypes.Healthcheck{}, err
	}
	return result, nil
}

type internalUpdateHealthcheckInput struct {
	Name        string            `json:"name"`
	Description string            `json:"description,omitempty"`
	Labels      map[string]string `json:"labels,omitempty"`
	Interval    string            `json:"interval"`
	Enabled     bool              `json:"enabled"`
}

func (c *Client) UpdateDNSHealthcheck(input apitypes.UpdateDNSHealthcheckInput) (apitypes.Response, error) {
	var result apitypes.Response
	internalInput := internalUpdateHealthcheckInput{
		Name:        input.Name,
		Description: input.Description,
		Labels:      input.Labels,
		Interval:    input.Interval,
		Enabled:     input.Enabled,
	}
	payload, err := jsonMerge(internalInput, input.HealthcheckDNSDefinition)
	if err != nil {
		return result, err
	}
	_, err = c.sendRequest(fmt.Sprintf("/api/v1/healthcheck/dns/%s", input.ID), http.MethodPut, payload, &result, nil, TokenAuth)
	if err != nil {
		return apitypes.Response{}, err
	}
	return result, nil
}

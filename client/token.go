package client

import (
	"fmt"
	"net/http"

	apitypes "github.com/appclacks/go-types"
)

func (c *Client) CreateAPIToken(payload apitypes.CreateAPITokenInput) (apitypes.APIToken, error) {
	var result apitypes.APIToken
	_, err := c.sendRequest("/app/v1/token", http.MethodPost, payload, &result, nil, PasswordAuth)
	if err != nil {
		return apitypes.APIToken{}, err
	}
	return result, nil
}
func (c *Client) ListAPITokens() (apitypes.ListAPITokensOutput, error) {
	var result apitypes.ListAPITokensOutput
	_, err := c.sendRequest("/api/v1/token", http.MethodGet, nil, &result, nil, TokenAuth)
	if err != nil {
		return apitypes.ListAPITokensOutput{}, err
	}
	return result, nil
}

func (c *Client) GetAPIToken(input apitypes.GetAPITokenInput) (apitypes.APIToken, error) {
	var result apitypes.APIToken
	_, err := c.sendRequest(fmt.Sprintf("/api/v1/token/%s", input.ID), http.MethodGet, nil, &result, nil, TokenAuth)
	if err != nil {
		return apitypes.APIToken{}, err
	}
	return result, nil
}

func (c *Client) DeleteAPIToken(input apitypes.DeleteAPITokenInput) (apitypes.Response, error) {
	var result apitypes.Response
	_, err := c.sendRequest(fmt.Sprintf("/api/v1/token/%s", input.ID), http.MethodDelete, nil, &result, nil, TokenAuth)
	if err != nil {
		return apitypes.Response{}, err
	}
	return result, nil
}

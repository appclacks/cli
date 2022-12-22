package client

import (
	"context"
	"fmt"
	"net/http"

	apitypes "github.com/appclacks/go-types"
)

func (c *Client) CreateAPIToken(ctx context.Context, payload apitypes.CreateAPITokenInput) (apitypes.APIToken, error) {
	var result apitypes.APIToken
	_, err := c.sendRequest(ctx, "/app/v1/token", http.MethodPost, payload, &result, nil, PasswordAuth)
	if err != nil {
		return apitypes.APIToken{}, err
	}
	return result, nil
}
func (c *Client) ListAPITokens(ctx context.Context) (apitypes.ListAPITokensOutput, error) {
	var result apitypes.ListAPITokensOutput
	_, err := c.sendRequest(ctx, "/api/v1/token", http.MethodGet, nil, &result, nil, TokenAuth)
	if err != nil {
		return apitypes.ListAPITokensOutput{}, err
	}
	return result, nil
}

func (c *Client) GetAPIToken(ctx context.Context, input apitypes.GetAPITokenInput) (apitypes.APIToken, error) {
	var result apitypes.APIToken
	_, err := c.sendRequest(ctx, fmt.Sprintf("/api/v1/token/%s", input.ID), http.MethodGet, nil, &result, nil, TokenAuth)
	if err != nil {
		return apitypes.APIToken{}, err
	}
	return result, nil
}

func (c *Client) DeleteAPIToken(ctx context.Context, input apitypes.DeleteAPITokenInput) (apitypes.Response, error) {
	var result apitypes.Response
	_, err := c.sendRequest(ctx, fmt.Sprintf("/api/v1/token/%s", input.ID), http.MethodDelete, nil, &result, nil, TokenAuth)
	if err != nil {
		return apitypes.Response{}, err
	}
	return result, nil
}

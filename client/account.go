package client

import (
	"context"
	"net/http"

	apitypes "github.com/appclacks/go-types"
)

func (c *Client) ChangeAccountPassword(ctx context.Context, payload apitypes.ChangeAccountPasswordInput) (apitypes.Response, error) {
	var result apitypes.Response
	_, err := c.sendRequest(ctx, "/app/v1/password/change", http.MethodPost, payload, &result, nil, PasswordAuth)
	if err != nil {
		return apitypes.Response{}, err
	}
	return result, nil
}

func (c *Client) AskResetPasswordLink(ctx context.Context, payload apitypes.ResetAccountPasswordLinkInput) (apitypes.Response, error) {
	var result apitypes.Response
	_, err := c.sendRequest(ctx, "/account/password/ask-reset", http.MethodPost, payload, &result, nil, NoAuth)
	if err != nil {
		return apitypes.Response{}, err
	}
	return result, nil
}

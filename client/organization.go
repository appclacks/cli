package client

import (
	"context"
	"net/http"

	apitypes "github.com/appclacks/go-types"
)

func (c *Client) CreateOrganization(ctx context.Context, payload apitypes.CreateOrganizationInput) (apitypes.CreateOrganizationOutput, error) {
	var result apitypes.CreateOrganizationOutput
	_, err := c.sendRequest(ctx, "/register", http.MethodPost, payload, &result, nil, NoAuth)
	if err != nil {
		return apitypes.CreateOrganizationOutput{}, err
	}
	return result, nil
}

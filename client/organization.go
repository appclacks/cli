package client

import (
	"net/http"

	apitypes "github.com/appclacks/go-types"
)

func (c *Client) CreateOrganization(payload apitypes.CreateOrganizationInput) (apitypes.CreateOrganizationOutput, error) {
	var result apitypes.CreateOrganizationOutput
	_, err := c.sendRequest("/register", http.MethodPost, payload, &result, nil, NoAuth)
	if err != nil {
		return apitypes.CreateOrganizationOutput{}, err
	}
	return result, nil
}

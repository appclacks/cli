package client

import (
	"context"

	apitypes "github.com/appclacks/go-types"
)

// ############################################################
// ############################################################
// 			Client public Organization-useCases



func (c *Client) CreateOrganizationWithContext(ctx context.Context, payload apitypes.CreateOrganizationInput) (apitypes.CreateOrganizationOutput, error) {
	var result apitypes.CreateOrganizationOutput

	if !c.Credentials.HasAuth() {
		return apitypes.CreateOrganizationOutput{}, formatError("CreateOrganizationWithContext", "Credentials missing")
	}

	if ctx == nil {
		return apitypes.CreateOrganizationOutput{}, formatError("CreateOrganizationWithContext", " 'nil' value for context.Context")
	}

	request, err := c.RequestsHelper.BuildCreateOrganizationRequest(payload)
	if err != nil {
		return apitypes.CreateOrganizationOutput{}, err 
	}

	request = request.WithContext(ctx)

	return sendRequestGetStruct(c.RequestsHelper, request, result)
	
}

func (c *Client) CreateOrganization(ctx context.Context, payload apitypes.CreateOrganizationInput) (apitypes.CreateOrganizationOutput, error) {
	var result apitypes.CreateOrganizationOutput

	if !c.Credentials.HasAuth() {
		return apitypes.CreateOrganizationOutput{}, formatError("CreateOrganization", "Credentials missing")
	}

	request, err := c.RequestsHelper.BuildCreateOrganizationRequest(payload)
	if err != nil {
		return apitypes.CreateOrganizationOutput{}, err 
	}

	return sendRequestGetStruct(c.RequestsHelper, request, result)
	
}

// ------------------------------------------------------------------------------------------------------------------------

func (c *Client) GetOrganizationForAccountWithContext(ctx context.Context) (apitypes.Organization, error) {
	
	var result apitypes.Organization
	
	if !c.Credentials.HasAuth() {
		return apitypes.Organization{}, formatError("GetOrganizationForAccountWithContext", "Credentials missing")
	}

	if ctx == nil {
		return apitypes.Organization{}, formatError("GetOrganizationForAccountWithContext", " 'nil' value for context.Context")
	}

	request, err := c.RequestsHelper.BuildGetOrganizationRequest()
	
	if err != nil {
		return apitypes.Organization{}, err 
	}
	
	request = request.WithContext(ctx)

	return sendRequestGetStruct(c.RequestsHelper, request, result)
}

func (c *Client) GetOrganizationForAccount() (apitypes.Organization, error) {
	var result apitypes.Organization
	
	if !c.Credentials.HasAuth() {
		return apitypes.Organization{}, formatError("GetOrganizationForAccountWithContext", "Credentials missing")
	}

	request, err := c.RequestsHelper.BuildGetOrganizationRequest()
	if err != nil {
		return apitypes.Organization{}, err 
	}

	return sendRequestGetStruct(c.RequestsHelper, request, result)
}
package client

import (
	"context"

	apitypes "github.com/appclacks/go-types"
)

// ############################################################
// ############################################################
// 			Client public Token-useCases 


func (c *Client) CreateAPITokenWithContext(ctx context.Context, payload apitypes.CreateAPITokenInput) (apitypes.APIToken, error) {
	var result apitypes.APIToken

	if !c.Credentials.HasAuth() {
		return apitypes.APIToken{}, formatError("CreateAPITokenWithContext", "Credentials missing")
	}

	if ctx == nil {
		return apitypes.APIToken{}, formatError("CreateAPITokenWithContext", " 'nil' value for context.Context")
	}

	request, err := c.RequestsHelper.BuildCreateTokenRequest(payload)
	if err != nil {
		return apitypes.APIToken{}, err
	}

	request = request.WithContext(ctx)

	return sendRequestGetStruct(c.RequestsHelper, request, result)
}

func (c *Client) CreateAPIToken(payload apitypes.CreateAPITokenInput) (apitypes.APIToken, error) {
	var result apitypes.APIToken

	if !c.Credentials.HasAuth() {
		return apitypes.APIToken{}, formatError("CreateAPIToken", "Credentials missing")
	}

	request, err := c.RequestsHelper.BuildCreateTokenRequest(payload)
	if err != nil {
		return apitypes.APIToken{}, err
	}


	return sendRequestGetStruct(c.RequestsHelper, request, result)
}

// -----------------------------------------------------------------------------

func (c *Client) ListAPITokensWithContext(ctx context.Context) (apitypes.ListAPITokensOutput, error) {
	var result apitypes.ListAPITokensOutput

	if !c.Credentials.HasAuth() {
		return apitypes.ListAPITokensOutput{}, formatError("ListAPITokensWithContext", "Credentials missing")
	}

	if ctx == nil {
		return apitypes.ListAPITokensOutput{}, formatError("ListAPITokensWithContext", " 'nil' value for context.Context")
	}
	
	request, err := c.RequestsHelper.BuildGetListTokenRequest()
	if err != nil {
		return apitypes.ListAPITokensOutput{}, err 
	}

	return sendRequestGetStruct(c.RequestsHelper, request, result)
}


func (c *Client) ListAPITokens() (apitypes.ListAPITokensOutput, error) {
	var result apitypes.ListAPITokensOutput

	if !c.Credentials.HasAuth() {
		return apitypes.ListAPITokensOutput{}, formatError("ListAPITokens", "Credentials missing")
	}

	request, err := c.RequestsHelper.BuildGetListTokenRequest()
	if err != nil {
		return apitypes.ListAPITokensOutput{}, err 
	}

	return sendRequestGetStruct(c.RequestsHelper, request, result)
}

// -----------------------------------------------------------------------------

func (c *Client) GetAPITokenWithContext(ctx context.Context, input apitypes.GetAPITokenInput) (apitypes.APIToken, error) {
	var result apitypes.APIToken
	
	if !c.Credentials.HasAuth() {
		return apitypes.APIToken{}, formatError("GetAPITokenWithContext", "Credentials missing")
	}

	if ctx == nil {
		return apitypes.APIToken{}, formatError("GetAPITokenWithContext", " 'nil' value for context.Context")
	}

	request, err := c.RequestsHelper.BuildGetTokenRequest(input.ID)
	if err != nil {
		return apitypes.APIToken{}, err
	}

	request = request.WithContext(ctx)

	return sendRequestGetStruct(c.RequestsHelper, request, result)
}

func (c *Client) GetAPIToken(input apitypes.GetAPITokenInput) (apitypes.APIToken, error) {
	var result apitypes.APIToken
	
	if !c.Credentials.HasAuth() {
		return apitypes.APIToken{}, formatError("GetAPITokenWithContext", "Credentials missing")
	}

	request, err := c.RequestsHelper.BuildGetTokenRequest(input.ID)
	if err != nil {
		return apitypes.APIToken{}, err
	}

	return sendRequestGetStruct(c.RequestsHelper, request, result)
}

// -----------------------------------------------------------------------------


func (c *Client) DeleteAPITokenWithContext(ctx context.Context, input apitypes.DeleteAPITokenInput) (apitypes.Response, error) {
	var result apitypes.Response
	
	if !c.Credentials.HasAuth() {
		return apitypes.Response{}, formatError("DeleteAPITokenWithContext", "Credentials missing")
	}

	if ctx == nil {
		return apitypes.Response{}, formatError("DeleteAPITokenWithContext", " 'nil' value for context.Context")
	}

	request, err := c.RequestsHelper.BuildDeleteTokenRequest(input.ID)
	if err != nil {
		return apitypes.Response{}, err
	}

	request = request.WithContext(ctx)

	return sendRequestGetStruct(c.RequestsHelper, request, result)
}

func (c *Client) DeleteAPIToken(input apitypes.DeleteAPITokenInput) (apitypes.Response, error) {
	var result apitypes.Response

	if !c.Credentials.HasAuth() {
		return apitypes.Response{}, formatError("DeleteAPIToken", "Credentials missing")
	}

	request, err := c.RequestsHelper.BuildDeleteTokenRequest(input.ID)
	if err != nil {
		return apitypes.Response{}, err
	}

	return sendRequestGetStruct(c.RequestsHelper, request, result)
}
package client

import (
	"context"

	apitypes "github.com/appclacks/go-types"
)


// ############################################################
// ############################################################
// 			Client public Account-useCases 

func (c *Client) ChangeAccountPasswordWithContext(ctx context.Context, payload apitypes.ChangeAccountPasswordInput) (apitypes.Response, error) {
	var result apitypes.Response

	if !c.Credentials.HasAuth() {
		return apitypes.Response{}, formatError("ChangeAccountPasswordWithContext", "Credentials missing")
	}

	if ctx == nil {
		return apitypes.Response{}, formatError("ChangeAccountPasswordWithContext", " 'nil' value for context.Context")
	}

	request, err := c.RequestsHelper.BuildChangePasswordRequest(payload)
	if err != nil {
		return apitypes.Response{}, err 
	}

	request = request.WithContext(ctx)

	return sendRequestGetStruct(c.RequestsHelper, request, result)

}

func (c *Client) ChangeAccountPassword(payload apitypes.ChangeAccountPasswordInput) (apitypes.Response, error) {
	var result apitypes.Response

	if !c.Credentials.HasAuth() {
		return apitypes.Response{}, formatError("ChangeAccountPassword", "Credentials missing")
	}

	request, err := c.RequestsHelper.BuildChangePasswordRequest(payload)
	if err != nil {
		return apitypes.Response{}, err 
	}

	return sendRequestGetStruct(c.RequestsHelper, request, result)
}



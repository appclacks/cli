package client

import (
	"context"

)

// ############################################################
// ############################################################
// 			Client public Metrics-useCases 

func (c *Client) GetHealthchecksMetricsWithContext(ctx context.Context) (string, error) {
	
	if !c.Credentials.HasAuth() {
		return "", formatError("GetHealthchecksMetricsWithContext", "Credentials missing")
	}

	if ctx == nil {
		return "", formatError("GetHealthchecksMetricsWithContext", " 'nil' value for context.Context")
	}

	request, err := c.RequestsHelper.BuildGetMetricsRequest()
	if err != nil {
		return "",err 
	}
	
	request = request.WithContext(ctx)

	return sendRequestGetString(c.RequestsHelper, request)
}

func (c *Client) GetHealthchecksMetrics() (string, error) {
	if !c.Credentials.HasAuth() {
		return "", formatError("GetHealthchecksMetrics", "Credentials missing")
	}

	request, err := c.RequestsHelper.BuildGetMetricsRequest()
	if err != nil {
		return "",err 
	}

	return sendRequestGetString(c.RequestsHelper, request)
}

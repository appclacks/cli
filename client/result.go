package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	apitypes "github.com/appclacks/go-types"
)

func (c *Client) ListHealthchecksResults(input apitypes.ListHealthchecksResultsInput) (apitypes.ListHealthchecksResultsOutput, error) {
	var result apitypes.ListHealthchecksResultsOutput
	queryParams := make(map[string]string)
	startBytes, err := json.Marshal(input.StartDate)
	if err != nil {
		return apitypes.ListHealthchecksResultsOutput{}, err
	}
	queryParams["start-date"], err = strconv.Unquote(string(startBytes))
	if err != nil {
		return apitypes.ListHealthchecksResultsOutput{}, err
	}
	endBytes, err := json.Marshal(input.EndDate)
	if err != nil {
		return apitypes.ListHealthchecksResultsOutput{}, err
	}
	queryParams["end-date"], err = strconv.Unquote(string(endBytes))
	if err != nil {
		return apitypes.ListHealthchecksResultsOutput{}, err
	}
	if input.HealthcheckID != "" {
		queryParams["healthcheck-id"] = input.HealthcheckID
	}
	if input.Page != 0 {
		queryParams["page"] = fmt.Sprintf("%d", input.Page)
	}
	if input.Success != nil {
		if *input.Success {
			queryParams["success"] = "true"
		} else {
			queryParams["success"] = "false"
		}
	}
	_, err = c.sendRequest("/api/v1/result/healthchecks", http.MethodGet, nil, &result, queryParams, TokenAuth)
	if err != nil {
		return apitypes.ListHealthchecksResultsOutput{}, err
	}
	return result, nil
}

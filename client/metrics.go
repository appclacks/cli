package client

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetHealthchecksMetrics() (string, error) {
	request, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("%s/api/v1/metrics/healthchecks", c.endpoint),
		nil)
	if err != nil {
		return "", err
	}
	// TODO: mutualize code
	authString := fmt.Sprintf("%s:%s", c.orgID, c.token)
	creds := base64.StdEncoding.EncodeToString([]byte(authString))
	request.Header.Add("Authorization", fmt.Sprintf("Basic %s", creds))
	response, err := c.http.Do(request)
	if err != nil {
		return "", err
	}
	b, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	if response.StatusCode >= 400 {
		return "", fmt.Errorf("The API returned an error: status %d\n%s", response.StatusCode, string(b))
	}
	return string(b), nil
}

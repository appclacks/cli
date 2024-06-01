package client

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Client struct {
	http     *http.Client
	username string
	password string
	endpoint string
	profile  string
}

var (
	ErrNotFound = errors.New("Not found")
)

func loadEnv(client *Client) {
	if os.Getenv("APPCLACKS_USER_NAME") != "" {
		client.username = os.Getenv("APPCLACKS_USER_NAME")
	}

	if os.Getenv("APPCLACKS_USER_PASSWORD") != "" {
		client.password = os.Getenv("APPCLACKS_USER_PASSWORD")
	}

	if os.Getenv("APPCLACKS_PROFILE") != "" {
		client.profile = os.Getenv("APPCLACKS_PROFILE")
	}

	if os.Getenv("APPCLACKS_API_ENDPOINT") != "" {
		client.endpoint = os.Getenv("APPCLACKS_API_ENDPOINT")
	}
}

func New() (*Client, error) {
	client := &Client{
		http: &http.Client{},
	}

	loadEnv(client)
	if (client.username == "" && client.password != "") || (client.username != "" && client.password == "") {
		return nil, errors.New("invalid basic auth credentials")
	}
	if client.endpoint == "" {
		client.endpoint = "http://127.0.0.1:9000"
	}
	return client, nil
}

func (c *Client) sendRequest(ctx context.Context, url string, method string, body any, result any, queryParams map[string]string) (*http.Response, error) {
	var reqBody io.Reader
	if body != nil {
		json, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		reqBody = bytes.NewBuffer(json)
	}
	request, err := http.NewRequestWithContext(
		ctx,
		method,
		fmt.Sprintf("%s%s", c.endpoint, url),
		reqBody)
	if err != nil {
		return nil, err
	}
	if len(queryParams) != 0 {
		q := request.URL.Query()
		for k, v := range queryParams {
			q.Add(k, v)
		}
		request.URL.RawQuery = q.Encode()
	}
	request.Header.Add("content-type", "application/json")
	if c.username != "" {
		authString := fmt.Sprintf("%s:%s", c.username, c.password)
		creds := base64.StdEncoding.EncodeToString([]byte(authString))
		request.Header.Add("Authorization", fmt.Sprintf("Basic %s", creds))
	}
	response, err := c.http.Do(request)
	if err != nil {
		return nil, err
	}
	b, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode >= 400 {
		if response.StatusCode == 404 {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("the API returned an error: status %d\n%s", response.StatusCode, string(b))
	}
	if result != nil {
		err = json.Unmarshal(b, result)
		if err != nil {
			return nil, err
		}
	}
	return response, nil
}

func jsonMerge(s1 any, s2 any) (map[string]any, error) {
	result := make(map[string]any)
	str1, err := json.Marshal(s1)
	if err != nil {
		return nil, err
	}
	str2, err := json.Marshal(s2)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(str1, &result)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(str2, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

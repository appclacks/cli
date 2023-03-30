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
	http            *http.Client
	endpoint        string
	orgID           string
	accountEmail    string
	token           string
	accountPassword string
	profile         string
}

type AuthMode string

const TokenAuth AuthMode = "token"
const PasswordAuth AuthMode = "password"
const NoAuth AuthMode = "none"

var (
	ErrNotFound = errors.New("Not found")
)

type CustomAuth func(*Client)
type WithTokenOpts func(*Client)
type WithUserPasswordOpts func(*Client)

func WithToken(opts ...WithTokenOpts) CustomAuth {
	return func(config *Client) {
		for _, opt := range opts {
			opt(config)
		}
	}
}

func WithProfile(profile string) CustomAuth {
	return func(config *Client) {
		config.profile = profile
	}
}

func OrganizationID(orgID string) WithTokenOpts {
	return func(c *Client) {
		c.orgID = orgID
	}
}

func Token(token string) WithTokenOpts {
	return func(c *Client) {
		c.token = token
	}
}

func WithUserPassword(opts ...WithUserPasswordOpts) CustomAuth {
	return func(config *Client) {
		for _, opt := range opts {
			opt(config)
		}
	}
}

func AccountEmail(email string) WithUserPasswordOpts {
	return func(c *Client) {
		c.accountEmail = email
	}
}

func AccountPassword(password string) WithUserPasswordOpts {
	return func(c *Client) {
		c.accountPassword = password
	}
}

func loadEnv(client *Client) {
	if os.Getenv("APPCLACKS_ORGANIZATION_ID") != "" {
		client.orgID = os.Getenv("APPCLACKS_ORGANIZATION_ID")
	}

	if os.Getenv("APPCLACKS_TOKEN") != "" {
		client.token = os.Getenv("APPCLACKS_TOKEN")
	}

	if os.Getenv("APPCLACKS_ACCOUNT_EMAIL") != "" {
		client.accountEmail = os.Getenv("APPCLACKS_ACCOUNT_EMAIL")
	}

	if os.Getenv("APPCLACKS_ACCOUNT_PASSWORD") != "" {
		client.accountPassword = os.Getenv("APPCLACKS_ACCOUNT_PASSWORD")
	}

	if os.Getenv("APPCLACKS_PROFILE") != "" {
		client.profile = os.Getenv("APPCLACKS_PROFILE")
	}
}

func loadConfigFile(client *Client) error {
	configPath, err := GetConfigFilePath()
	if err != nil {
		return err
	}
	config, err := ReadConfig(configPath)
	if err != nil {
		return err
	}
	if len(config.Profiles) == 0 {
		return fmt.Errorf("The configuration file %s is empty", configPath)
	}
	profile := client.profile
	if profile == "" {
		profile = config.DefaultProfile
	}
	if profile == "" {
		return fmt.Errorf("No profile selected and no default profile specified in the configuration file %s", configPath)
	}
	configProfile, ok := config.Profiles[profile]
	if !ok {
		return fmt.Errorf("Profile %s not found in the configuration file %s", profile, configPath)
	}

	client.orgID = configProfile.OrganizationID
	client.token = configProfile.APIToken
	return nil
}

func New(endpoint string, customAuth ...CustomAuth) (*Client, error) {

	client := &Client{
		http:     &http.Client{},
		endpoint: endpoint,
	}

	for _, auth := range customAuth {
		auth(client)
	}

	loadEnv(client)

	if client.accountEmail == "" && client.accountPassword == "" && client.token == "" && client.orgID == "" {
		err := loadConfigFile(client)
		if err != nil {
			return nil, fmt.Errorf("No authentication variables defined and error while loading the Appclacks configuration file: %w", err)
		}
	}

	return client, nil

}

func (c *Client) sendRequest(ctx context.Context, url string, method string, body any, result any, queryParams map[string]string, auth AuthMode) (*http.Response, error) {
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
	if len(queryParams) != 0 {
		q := request.URL.Query()
		for k, v := range queryParams {
			q.Add(k, v)
		}
		request.URL.RawQuery = q.Encode()
	}
	request.Header.Add("content-type", "application/json")
	if auth != NoAuth {
		var authString string
		if auth == TokenAuth {
			authString = fmt.Sprintf("%s:%s", c.orgID, c.token)
		}
		if auth == PasswordAuth {
			authString = fmt.Sprintf("%s:%s", c.accountEmail, c.accountPassword)
		}
		creds := base64.StdEncoding.EncodeToString([]byte(authString))
		request.Header.Add("Authorization", fmt.Sprintf("Basic %s", creds))
	}
	if err != nil {
		return nil, err
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
		return nil, fmt.Errorf("The API returned an error: status %d\n%s", response.StatusCode, string(b))
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

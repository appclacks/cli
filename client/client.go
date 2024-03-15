package client

import (
	"encoding/json"
	"fmt"
	"github.com/appclacks/cli/client/api"
	"net/http"
)

func deserializeResult(data []byte, v any) error{
	return json.Unmarshal(data, v)
}

func formatError(funcName string, message string) error {
	return fmt.Errorf(fmt.Sprintf("AppClacks Client Error :: %s(...) :: %s", funcName, message))	
}

func sendRequestGetStruct[T any](requestsHelper *api.RequestsHelper, request *http.Request, result T) (T, error){

	data, err := requestsHelper.ConsumeRequest(request)
	if err != nil {
		return result, err
	}
	
	err = deserializeResult(data, &result)
	
	return result, err
}

func sendRequestGetString(requestsHelper *api.RequestsHelper, request *http.Request) (string, error){
	var result string 

	data, err := requestsHelper.ConsumeRequest(request)
	if err != nil {
		return result, err
	}

	result = string(data) 
	
	return result, err
}


// ############################################################
// ############################################################
//              	Options Struct

type ClientOptions struct {
	baseUrl 		string
	credentials 		*api.Credentials
}

func (opts *ClientOptions) UsePasswordAuth(email string, password string) (*ClientOptions) {
	opts.credentials.Type = api.PasswordAuth
	opts.credentials.Identifier = email 
	opts.credentials.Secret = password
	return opts
}

func (opts *ClientOptions) UseTokenAuth(orgID string, token string) (*ClientOptions) {
	opts.credentials.Type = api.TokenAuth
	opts.credentials.Identifier = orgID
	opts.credentials.Secret = token
	return opts 
}

func (opts *ClientOptions) SetBaseUrl(baseUrl string) (*ClientOptions) {
	opts.baseUrl = baseUrl
	return opts
}

func GetDefaultClientOptions() (*ClientOptions){
	return &ClientOptions{
		baseUrl: api.API_DEFAULT_BASE_URL,
		credentials: api.CreateCredentials(api.NoAuth, "", ""),
	}
}

//
// ############################################################
// ############################################################


// ############################################################
// ############################################################
//              	Client Struct

type Client struct {
	Credentials *api.Credentials 
	RequestsHelper *api.RequestsHelper
}

func New(clientOptions *ClientOptions) (*Client) {

	return &Client{
		Credentials:  clientOptions.credentials,
		RequestsHelper: api.CreateRequestsHelper(clientOptions.credentials, clientOptions.baseUrl),
	}

}



// TODO : SUPPRIMER
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

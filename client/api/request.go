package api

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"net/http"
	apitypes "github.com/appclacks/go-types"
)

// ############################################################
//

type AuthMode string

const TokenAuth AuthMode = "token"
const PasswordAuth AuthMode = "password"
const NoAuth AuthMode = "none"

var (
	ErrNotFound = errors.New("Not found")
)

// ############################################################
// ############################################################
// 					Credentials Struct et Functions

type Credentials struct {
	Type AuthMode 
	Identifier string
	Secret string 
}

func (c *Credentials) HasAuth() bool {
	if c.Type == NoAuth { return false} 
	return true
}

func (c *Credentials) Base64AuthStr() string {
	authString := fmt.Sprintf("%s:%s", c.Identifier, c.Secret)
	return base64.StdEncoding.EncodeToString([]byte(authString))

}

func CreateCredentials(mType AuthMode, identifier string, secret string) (*Credentials){
	return &Credentials{
		Type : mType,
		Identifier: identifier,
		Secret: secret,
	}
}
//
// ###############################################################
// ###############################################################

// ############################################################
// ############################################################
// 			Requester Struct and Request maker functions 

func CreateRequestsHelper(credentials *Credentials, baseURL string) (*RequestsHelper){
	return &RequestsHelper{
		Credentials: credentials,
		Profiler: CreateApiProfiler(baseURL),
		HttpClient: &http.Client{},
	}
}

type RequestsHelper struct {
	Credentials *Credentials
	Profiler *profiler
	HttpClient *http.Client
}

func (r *RequestsHelper) ConsumeRequest(request *http.Request) ([]byte, error) {
	
	response, err := r.HttpClient.Do(request)

	if err != nil {
		return nil,err
	}
	b, err := io.ReadAll(response.Body)
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		if response.StatusCode == 404 {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("The API returned an error: status %d\n%s", response.StatusCode, string(b))
	}

	return b, nil
	
}

func (r *RequestsHelper) BuildCreateOrganizationRequest(newOrganization apitypes.CreateOrganizationInput) (*http.Request, error){
	
	OCEndPoint 	:= r.Profiler.OrganizationCreate
	fullURL 	:= getFullURL(r.Profiler.BaseURL, OCEndPoint.getFullURI(NoAuth))
	method 		:= string(OCEndPoint.Method)
	
	body, err := marshalBody(newOrganization)
	if err != nil {
		return nil, err
	}
	
	request,err := createRequest(fullURL, method, body, nil)
	if err != nil {
		return nil, err 
	}
	
	setRequestHeader(r.Credentials, request)
	
	return request, nil

}

func (r *RequestsHelper) BuildGetOrganizationRequest()(*http.Request, error){

	OGOEndPoint	:= r.Profiler.OrganizationGetOne
	fullURL 	:= getFullURL(r.Profiler.BaseURL, OGOEndPoint.getFullURI(r.Credentials.Type))
	method		:= string(OGOEndPoint.Method)
	
	request,err := createRequest(fullURL, method, nil, nil)
	if err != nil {
		return nil, err 
	}
	
	setRequestHeader(r.Credentials, request)
	
	return request, nil 
}

func (r *RequestsHelper) BuildChangePasswordRequest(newPassword apitypes.ChangeAccountPasswordInput)(*http.Request, error){

	CPEndPoint	:= r.Profiler.PasswordChange
	fullURL		:= getFullURL(r.Profiler.BaseURL, CPEndPoint.getFullURI((r.Credentials.Type)))
	method		:= string(CPEndPoint.Method)

	body, err := marshalBody(newPassword)
	if err != nil {
		return nil, err
	}

	request,err := createRequest(fullURL, method, body, nil)
	if err != nil {
		return nil, err 
	}
	
	setRequestHeader(r.Credentials, request)
	
	return request, nil

}

func (r *RequestsHelper) BuildCreateTokenRequest(newToken apitypes.CreateAPITokenInput) (*http.Request, error){
	CTEndPoint	:= r.Profiler.TokenCreate
	fullURL		:= getFullURL(r.Profiler.BaseURL, CTEndPoint.getFullURI((r.Credentials.Type)))
	method		:= string(CTEndPoint.Method)

	body, err := marshalBody(newToken)
	if err != nil {
		return nil, err
	}

	request,err := createRequest(fullURL, method, body, nil)
	if err != nil {
		return nil, err 
	}
	
	setRequestHeader(r.Credentials, request)
	
	return request, nil

}

func (r *RequestsHelper) BuildGetListTokenRequest() (*http.Request, error){
	GLTEndPoint	:= r.Profiler.TokenGetList
	fullURL		:= getFullURL(r.Profiler.BaseURL, GLTEndPoint.getFullURI(r.Credentials.Type))
	method 		:= string(GLTEndPoint.Method)

	request,err := createRequest(fullURL, method, nil, nil)
	if err != nil {
		return nil, err 
	}
	
	setRequestHeader(r.Credentials, request)
	
	return request, nil 
}

func (r *RequestsHelper) BuildGetTokenRequest(tokenID string) (*http.Request, error){
	GTEndPoint	:= r.Profiler.TokenGetList
	fullURL 	:= getFullURL(r.Profiler.BaseURL, GTEndPoint.getFullURI(r.Credentials.Type, tokenID))
	method 		:= string(GTEndPoint.Method)

	request,err := createRequest(fullURL, method, nil, nil)
	if err != nil {
		return nil, err 
	}
	
	setRequestHeader(r.Credentials, request)
	
	return request, nil 

}

func (r *RequestsHelper) BuildDeleteTokenRequest(tokenID string) (*http.Request, error){
	DTEndPoint	:= r.Profiler.TokenDelete
	fullURL 	:= getFullURL(r.Profiler.BaseURL, DTEndPoint.getFullURI(r.Credentials.Type, tokenID))
	method 		:= string(DTEndPoint.Method)

	request,err := createRequest(fullURL, method, nil, nil)
	if err != nil {
		return nil, err 
	}
	
	setRequestHeader(r.Credentials, request)
	
	return request, nil 
}

func (r *RequestsHelper) BuildGetMetricsRequest() (*http.Request, error){
	GMEndPoint	:= r.Profiler.MetricGetList
	fullURL 	:= getFullURL(r.Profiler.BaseURL, GMEndPoint.getFullURI(r.Credentials.Type))
	method 		:= string(GMEndPoint.Method)

	request,err := createRequest(fullURL, method, nil, nil)
	if err != nil {
		return nil, err 
	}
	
	setRequestHeader(r.Credentials, request)
	
	return request, nil 
}

func (r *RequestsHelper) BuildGetResultsRequest(params apitypes.ListHealthchecksResultsInput) (*http.Request, error){
	GREndPoint 	:= r.Profiler.ResultGetList
	fullURL 	:= getFullURL(r.Profiler.BaseURL, GREndPoint.getFullURI(r.Credentials.Type))
	method 		:= string(GREndPoint.Method)

	queryParams := make(map[string]string)
	queryParams["start-date"] = params.StartDate.String()
	queryParams["end-date"] = params.EndDate.String()
	if params.HealthcheckID != "" {
		queryParams["healthcheck-id"] = params.HealthcheckID
	}
	if params.Page != 0 {
		queryParams["page"] = fmt.Sprintf("%d", params.Page)
	}
	if params.Success != nil {
		queryParams["success"] = "false"
		if *params.Success {
			queryParams["success"] = "true"
		}
	}

	request,err := createRequest(fullURL, method, nil, queryParams)
	if err != nil {
		return nil, err 
	}
	
	setRequestHeader(r.Credentials, request)
	
	return request, nil 

}

func (r *RequestsHelper) BuildGetListHealthCheckRequest() (*http.Request, error){
	GLHCEndPoint 	:= r.Profiler.HealthCheckGetList
	fullURL 		:= getFullURL(r.Profiler.BaseURL, GLHCEndPoint.getFullURI(r.Credentials.Type))
	method 			:= string(GLHCEndPoint.Method)

	request,err := createRequest(fullURL, method, nil, nil)
	if err != nil {
		return nil, err 
	}
	
	setRequestHeader(r.Credentials, request)
	
	return request, nil 

}

func (r *RequestsHelper) BuildGetHealthCheckRequest(healthCheckID string) (*http.Request, error){
	GHCEndPoint 	:= r.Profiler.HealthCheckGetOne
	fullURL 		:= getFullURL(r.Profiler.BaseURL, GHCEndPoint.getFullURI(r.Credentials.Type, healthCheckID))
	method 			:= string(GHCEndPoint.Method)

	request,err := createRequest(fullURL, method, nil, nil)
	if err != nil {
		return nil, err 
	}
	
	setRequestHeader(r.Credentials, request)
	
	return request, nil 

}

func (r *RequestsHelper) BuildDeleteHealthCheckRequest(healthCheckID string) (*http.Request, error){
	DHCEndPoint 	:= r.Profiler.HealthCheckDeleteOne
	fullURL 		:= getFullURL(r.Profiler.BaseURL, DHCEndPoint.getFullURI(r.Credentials.Type, healthCheckID))
	method 			:= string(DHCEndPoint.Method)

	request,err := createRequest(fullURL, method, nil, nil)
	if err != nil {
		return nil, err 
	}
	
	setRequestHeader(r.Credentials, request)
	
	return request, nil 
}

func (r *RequestsHelper) BuildCabourotteDiscoveryRequest(labels string) (*http.Request, error){
	CDEndPoint 	:= r.Profiler.CabourotteDiscovery
	fullURL 	:= getFullURL(r.Profiler.BaseURL, CDEndPoint.getFullURI(r.Credentials.Type))
	method 		:= string(CDEndPoint.Method)

	queryParams := make(map[string]string)
	if labels != ""{
		queryParams["labels"] = labels 
	}

	request,err := createRequest(fullURL, method, nil, queryParams)
	if err != nil {
		return nil, err 
	}
	
	setRequestHeader(r.Credentials, request)
	
	return request, nil 

}

func (r *RequestsHelper) BuildCreateCommandHealthCheckRequest(newHealthCheckCommand apitypes.CreateCommandHealthcheckInput) (*http.Request, error){
	CCHCEndPoint 	:= r.Profiler.HealthCheckCmdCreate
	fullURL 		:= getFullURL(r.Profiler.BaseURL, CCHCEndPoint.getFullURI(r.Credentials.Type))
	method 			:= string(CCHCEndPoint.Method)

	body, err := marshalBody(newHealthCheckCommand)
	if err != nil {
		return nil, err
	}

	request,err := createRequest(fullURL, method, body, nil)
	if err != nil {
		return nil, err 
	}
	
	setRequestHeader(r.Credentials, request)
	
	return request, nil 

}

func (r *RequestsHelper) BuildUpdateCommandHealthCheckRequest(healthCheckCommand apitypes.UpdateCommandHealthcheckInput) (*http.Request, error){
	UCHCEndPoint 	:= r.Profiler.HealthCheckCmdUpdate
	fullURL 		:= getFullURL(r.Profiler.BaseURL, UCHCEndPoint.getFullURI(r.Credentials.Type, healthCheckCommand.ID))
	method 			:= string(UCHCEndPoint.Method)

	body, err := marshalBody(healthCheckCommand)
	if err != nil {
		return nil, err
	}

	request,err := createRequest(fullURL, method, body, nil)
	if err != nil {
		return nil, err 
	}
	
	setRequestHeader(r.Credentials, request)
	
	return request, nil 
}

func (r *RequestsHelper) BuildCreateDNSHealthCheckRequest(newHealthCheckDNS apitypes.CreateDNSHealthcheckInput) (*http.Request, error){
	CDHCEndPoint 	:= r.Profiler.HealthCheckDNSCreate
	fullURL 		:= getFullURL(r.Profiler.BaseURL, CDHCEndPoint.getFullURI(r.Credentials.Type))
	method 			:= string(CDHCEndPoint.Method)

	body, err := marshalBody(newHealthCheckDNS)
	if err != nil {
		return nil, err
	}

	request,err := createRequest(fullURL, method, body, nil)
	if err != nil {
		return nil, err 
	}
	
	setRequestHeader(r.Credentials, request)
	
	return request, nil 

}

func (r *RequestsHelper) BuildUpdateDNSHealthCheckRequest(healthCheckDNS apitypes.UpdateDNSHealthcheckInput) (*http.Request, error){
	UDHCEndPoint 	:= r.Profiler.HealthCheckDNSUpdate
	fullURL 		:= getFullURL(r.Profiler.BaseURL, UDHCEndPoint.getFullURI(r.Credentials.Type, healthCheckDNS.ID))
	method 			:= string(UDHCEndPoint.Method)

	body, err := marshalBody(healthCheckDNS)
	if err != nil {
		return nil, err
	}

	request,err := createRequest(fullURL, method, body, nil)
	if err != nil {
		return nil, err 
	}
	
	setRequestHeader(r.Credentials, request)
	
	return request, nil 
}

func (r *RequestsHelper) BuildCreateHTTPHealthCheckRequest(newHealthCheckHTTP apitypes.CreateHTTPHealthcheckInput) (*http.Request, error){
	CHHCEndPoint 	:= r.Profiler.HealthCheckHTTPCreate
	fullURL 		:= getFullURL(r.Profiler.BaseURL, CHHCEndPoint.getFullURI(r.Credentials.Type))
	method 			:= string(CHHCEndPoint.Method)

	body, err := marshalBody(newHealthCheckHTTP)
	if err != nil {
		return nil, err
	}

	request,err := createRequest(fullURL, method, body, nil)
	if err != nil {
		return nil, err 
	}
	
	setRequestHeader(r.Credentials, request)
	
	return request, nil 

}

func (r *RequestsHelper) BuildUpdateHTTPHealthCheckRequest(healthCheckHTTP apitypes.UpdateHTTPHealthcheckInput) (*http.Request, error){
	UHHCEndPoint 	:= r.Profiler.HealthCheckHTTPUpdate
	fullURL 		:= getFullURL(r.Profiler.BaseURL, UHHCEndPoint.getFullURI(r.Credentials.Type, healthCheckHTTP.ID))
	method 			:= string(UHHCEndPoint.Method)

	body, err := marshalBody(healthCheckHTTP)
	if err != nil {
		return nil, err
	}

	request,err := createRequest(fullURL, method, body, nil)
	if err != nil {
		return nil, err 
	}
	
	setRequestHeader(r.Credentials, request)
	
	return request, nil 
}

func (r *RequestsHelper) BuildCreateTCPHealthCheckRequest(newHealthCheckTCP apitypes.CreateTCPHealthcheckInput) (*http.Request, error){
	CTHCEndPoint 	:= r.Profiler.HealthCheckTCPCreate
	fullURL 		:= getFullURL(r.Profiler.BaseURL, CTHCEndPoint.getFullURI(r.Credentials.Type))
	method 			:= string(CTHCEndPoint.Method)

	body, err := marshalBody(newHealthCheckTCP)
	if err != nil {
		return nil, err
	}

	request,err := createRequest(fullURL, method, body, nil)
	if err != nil {
		return nil, err 
	}
	
	setRequestHeader(r.Credentials, request)
	
	return request, nil 

}

func (r *RequestsHelper) BuildUpdateTCPHealthCheckRequest(healthCheckTCP apitypes.UpdateTCPHealthcheckInput) (*http.Request, error){
	UTHCEndPoint 	:= r.Profiler.HealthCheckTCPUpdate
	fullURL 		:= getFullURL(r.Profiler.BaseURL, UTHCEndPoint.getFullURI(r.Credentials.Type, healthCheckTCP.ID))
	method 			:= string(UTHCEndPoint.Method)

	body, err := marshalBody(healthCheckTCP)
	if err != nil {
		return nil, err
	}

	request,err := createRequest(fullURL, method, body, nil)
	if err != nil {
		return nil, err 
	}
	
	setRequestHeader(r.Credentials, request)
	
	return request, nil 
}

func (r *RequestsHelper) BuildCreateTLSHealthCheckRequest(newHealthCheckTLS apitypes.CreateTLSHealthcheckInput) (*http.Request, error){
	CTHCEndPoint 	:= r.Profiler.HealthCheckTLSCreate
	fullURL 		:= getFullURL(r.Profiler.BaseURL, CTHCEndPoint.getFullURI(r.Credentials.Type))
	method 			:= string(CTHCEndPoint.Method)

	body, err := marshalBody(newHealthCheckTLS)
	if err != nil {
		return nil, err
	}

	request,err := createRequest(fullURL, method, body, nil)
	if err != nil {
		return nil, err 
	}
	
	setRequestHeader(r.Credentials, request)
	
	return request, nil 

}

func (r *RequestsHelper) BuildUpdateTLSHealthCheckRequest(healthCheckTLS apitypes.UpdateTLSHealthcheckInput) (*http.Request, error){
	UTHCEndPoint 	:= r.Profiler.HealthCheckTLSUpdate
	fullURL 		:= getFullURL(r.Profiler.BaseURL, UTHCEndPoint.getFullURI(r.Credentials.Type, healthCheckTLS.ID))
	method 			:= string(UTHCEndPoint.Method)

	body, err := marshalBody(healthCheckTLS)
	if err != nil {
		return nil, err
	}

	request,err := createRequest(fullURL, method, body, nil)
	if err != nil {
		return nil, err 
	}
	
	setRequestHeader(r.Credentials, request)
	
	return request, nil 
}
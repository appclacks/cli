package api


import (
	"fmt"
	"net/http"
)

type RequestURI string 
type HttpMethod string
type PrefixURI 	string

const API_DEFAULT_BASE_URL 		string = "https://api.appclacks.com"

const (
	TOKEN_PREFIX_URI PrefixURI = "/api"
	PASSWORD_PREFIX_URI PrefixURI = "/app"
	NONE_PREFIX_URI PrefixURI = ""
)

const (
	V1_HC_URI				RequestURI = "/v1/healthcheck"
	V1_HC_COMMAND_URI 		RequestURI = "/v1/healthcheck/command"
	V1_HC_DNS_URI 			RequestURI = "/v1/healthcheck/dns"
	V1_HC_HTTP_URI			RequestURI = "/v1/healthcheck/http"
	V1_HC_TCP_URI			RequestURI = "/v1/healthcheck/tcp"
	V1_HC_TLS_URI			RequestURI = "/v1/healthcheck/tls"
	
	V1_METRICS_URI			RequestURI = "/v1/metrics/healthchecks"
	V1_RESULTS_URI			RequestURI = "/v1/result/healthchecks"
	
	V1_ORGANIZATION_URI		RequestURI = "/v1/organization"
	V1_TOKEN_URI			RequestURI = "/v1/token"
	V1_PASSWORD_URI			RequestURI = "/v1/password/change"

	REGISTER_URI			RequestURI = "/register"
	CABOUROTTE_DISCOVERY_URI	RequestURI = "/cabourotte/discovery"
	
)

const (
	METH_GET 	HttpMethod = http.MethodGet
	METH_HEAD 	HttpMethod = http.MethodHead
	METH_POST 	HttpMethod = http.MethodPost
	METH_PUT 	HttpMethod = http.MethodPut
	METH_DELETE HttpMethod = http.MethodDelete
)

type endPoint struct {
	Name string
	Uri RequestURI
	Method HttpMethod
	Suffix string
}

func (ep *endPoint) getFullURI(authMode AuthMode, params ...any) string{
	var authPrefix string = ""
	if authMode == TokenAuth {
		authPrefix = string(TOKEN_PREFIX_URI)
	}
	if authMode == PasswordAuth {
		authPrefix = string(PASSWORD_PREFIX_URI)
	}

	return fmt.Sprintf(authPrefix+string(ep.Uri)+ep.Suffix, params...)
}

var healthCheckGetList *endPoint = &endPoint{
	Name : "HealthCheckGetList",
	Uri : V1_HC_URI,
	Method: METH_GET,
	Suffix: "",
}
var healthCheckGetOne *endPoint = &endPoint{
	Name : "HealthCheckGetOne",
	Uri : V1_HC_URI,
	Method: METH_GET,
	Suffix: "/%s",
}
var healthCheckDeleteOne *endPoint = &endPoint{
	Name : "HealthCheckDeleteOne",
	Uri : V1_HC_URI,
	Method: METH_DELETE,
	Suffix: "/%s",
}
var healthCheckCmdCreate *endPoint = &endPoint{
	Name : "HealthCheckCmdCreate",
	Uri : V1_HC_COMMAND_URI,
	Method: METH_POST,
	Suffix: "",
}
var healthCheckCmdUpdate *endPoint = &endPoint{
	Name : "HealthCheckCmdUpdate",
	Uri : V1_HC_COMMAND_URI,
	Method : METH_PUT,
	Suffix: "/%s",
}
var healthCheckDNSCreate *endPoint = &endPoint{
	Name : "HealthCheckDNSCreate",
	Uri : V1_HC_DNS_URI,
	Method: METH_POST,
	Suffix: "",
}
var healthCheckDNSUpdate *endPoint = &endPoint{
	Name : "HealthCheckDNSUpdate",
	Uri : V1_HC_DNS_URI,
	Method: METH_PUT,
	Suffix: "/%s",
}
var healthCheckHTTPCreate *endPoint = &endPoint{
	Name : "HealthCheckHTTPCreate",
	Uri : V1_HC_HTTP_URI,
	Method: METH_POST,
	Suffix: "",
}
var healthCheckHTTPUpdate *endPoint = &endPoint{
	Name : "HealthCheckHTTPUpdate",
	Uri : V1_HC_HTTP_URI,
	Method: METH_PUT,
	Suffix: "/%s",
}
var healthCheckTCPCreate *endPoint = &endPoint{
	Name : "HealthCheckTCPCreate",
	Uri : V1_HC_TCP_URI,
	Method: METH_POST,
	Suffix: "",
}
var healthCheckTCPUpdate *endPoint = &endPoint{
	Name : "HealthCheckTCPUpdate",
	Uri : V1_HC_TCP_URI,
	Method: METH_PUT,
	Suffix: "/%s",
}
var healthCheckTLSCreate *endPoint = &endPoint{
	Name : "HealthCheckTLSCreate",
	Uri : V1_HC_TLS_URI,
	Method: METH_POST,
	Suffix: "",
}
var healthCheckTLSUpdate *endPoint = &endPoint{
	Name : "HealthCheckTLSUpdate",
	Uri : V1_HC_TLS_URI,
	Method: METH_PUT,
	Suffix: "/%s",
}
var metricGetList *endPoint = &endPoint{
	Name : "MetricGetList",
	Uri : V1_METRICS_URI,
	Method: METH_GET,
	Suffix: "",
}
var resultGetList *endPoint = &endPoint{
	Name : "ResultGetList",
	Uri : V1_RESULTS_URI,
	Method: METH_GET,
	Suffix: "",
}
var organizationCreate *endPoint = &endPoint{
	Name : "OrganizationCreate",
	Uri : REGISTER_URI,
	Method: METH_POST,
	Suffix: "",
}
var organizationGetOne *endPoint = &endPoint{
	Name : "OrganizationGetOne",
	Uri : V1_ORGANIZATION_URI,
	Method: METH_GET,
	Suffix: "",
}
var tokenCreate *endPoint = &endPoint{
	Name : "TokenCreate",
	Uri : V1_TOKEN_URI,
	Method : METH_POST,
	Suffix : "",
}
var tokenDelete *endPoint = &endPoint{
	Name : "TokenDelete",
	Uri : V1_TOKEN_URI,
	Method : METH_DELETE,
	Suffix : "/%s",
}
var tokenGetOne *endPoint = &endPoint{
	Name : "TokenGetOne",
	Uri : V1_TOKEN_URI,
	Method : METH_GET,
	Suffix : "/%s",
}
var tokenGetList *endPoint = &endPoint{
	Name : "TokenGetList",
	Uri : V1_TOKEN_URI,
	Method : METH_GET,
	Suffix : "",
}
var passwordChange *endPoint = &endPoint{
	Name : "PasswordChange",
	Uri : V1_PASSWORD_URI,
	Method : METH_POST,
	Suffix : "",
}
var cabourotteDiscovery *endPoint = &endPoint{
	Name : "CabourotteDiscovery",
	Uri : CABOUROTTE_DISCOVERY_URI,
	Method: METH_GET,
	Suffix : "",
}


type profiler struct {
	BaseURL				string

	HealthCheckGetList 		*endPoint
	HealthCheckGetOne 		*endPoint
	HealthCheckDeleteOne	*endPoint 
	
	HealthCheckCmdCreate	*endPoint
	HealthCheckCmdUpdate 	*endPoint
	
	HealthCheckDNSCreate	*endPoint
	HealthCheckDNSUpdate	*endPoint

	HealthCheckHTTPCreate 	*endPoint
	HealthCheckHTTPUpdate	*endPoint 

	HealthCheckTCPCreate 	*endPoint
	HealthCheckTCPUpdate 	*endPoint

	HealthCheckTLSCreate	*endPoint
	HealthCheckTLSUpdate	*endPoint

	MetricGetList			*endPoint 
	ResultGetList			*endPoint

	OrganizationCreate		*endPoint
	OrganizationGetOne		*endPoint

	TokenCreate				*endPoint
	TokenDelete				*endPoint
	TokenGetOne				*endPoint
	TokenGetList			*endPoint

	CabourotteDiscovery		*endPoint
	PasswordChange 			*endPoint

	
}

func CreateApiProfiler(baseURL string) (*profiler){
	var ApiProfiler profiler = profiler{
		BaseURL: baseURL,
		HealthCheckGetList: healthCheckGetList,
		HealthCheckGetOne: healthCheckGetOne,
		HealthCheckDeleteOne: healthCheckDeleteOne,
		HealthCheckCmdCreate: healthCheckCmdCreate,
		HealthCheckCmdUpdate: healthCheckCmdUpdate,
		HealthCheckDNSCreate: healthCheckDNSCreate,
		HealthCheckDNSUpdate: healthCheckDNSUpdate,
		HealthCheckHTTPCreate: healthCheckHTTPCreate,
		HealthCheckHTTPUpdate: healthCheckHTTPUpdate,
		HealthCheckTCPCreate: healthCheckTCPCreate,
		HealthCheckTCPUpdate: healthCheckTCPUpdate,
		HealthCheckTLSCreate: healthCheckTLSCreate,
		HealthCheckTLSUpdate: healthCheckTLSUpdate,
		MetricGetList: metricGetList,
		ResultGetList: resultGetList,
		OrganizationCreate: organizationCreate,
		OrganizationGetOne: organizationGetOne,
		TokenCreate: tokenCreate,
		TokenDelete: tokenDelete,
		TokenGetOne: tokenGetOne,
		TokenGetList: tokenGetList,
		CabourotteDiscovery: cabourotteDiscovery,
		PasswordChange: passwordChange,
	}

	return &ApiProfiler
}

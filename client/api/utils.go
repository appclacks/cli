package api
import(
	"bytes"
	"encoding/json"
	"net/url"
	"net/http"
	"fmt"
	"io"
)

func rawQueryParams(queryParams map[string]string) string {
	q := url.Values{}
	for k, v := range queryParams {
		q.Add(k,v)
	}
	return q.Encode()
}

func marshalBody(body any) (*bytes.Buffer, error){
	json, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	return bytes.NewBuffer(json), nil
}

func setRequestHeader(credentials *Credentials, request *http.Request){
	request.Header.Add("content-type", "application/json")
	if credentials.Type != NoAuth {
		request.Header.Add("Authorization", fmt.Sprintf("Basic %s", credentials.Base64AuthStr()))
	}
}

func createRequest(url string, method string,body *bytes.Buffer, queryParams map[string]string) (*http.Request, error){
	
	var ioBody io.Reader = nil
	if body != nil {
		ioBody = body 
	} 
	
	request, err := http.NewRequest(
		method,
		url,
		ioBody)
	if err != nil {
		return nil, err 
	}
	if queryParams != nil {
		request.URL.RawQuery = rawQueryParams(queryParams)
	}

	return request, nil

}

func getFullURL(baseURL string, fullURI string) string {
	return fmt.Sprintf("%s%s", baseURL, fullURI)
}

package isuphttp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// HTTPRequest A request for a http call
// If InsecureRequest is true the ssl certificate is not validated
// TimeOut is the call timeout in milisecounds, default is 2000 ms, max is 60000 ms
type HTTPRequest struct {
	url             string
	method          string
	headers         map[string]interface{}
	body            map[string]interface{}
	queryParams     map[string]interface{}
	insecureRequest bool
	timeOut         int
}

const (
	timeOut    = 2000
	minTimeOut = 50
	maxTimeout = 60000
)

// GetHTTPRequest Instantiate a HTTP request object
func GetHTTPRequest(method string, url string) HTTPRequest {
	h := HTTPRequest{url: url, method: strings.ToUpper(method), timeOut: timeOut}

	return h
}

// SetHeaders Set headers values
func (h HTTPRequest) SetHeaders(headers map[string]interface{}) HTTPRequest {
	if h.headers == nil {
		h.headers = make(map[string]interface{})
	}

	for index, value := range headers {
		h.headers[index] = value
	}
	return h
}

// SetHeaderValue Set a header value
func (h HTTPRequest) SetHeaderValue(name string, value interface{}) HTTPRequest {
	if h.headers == nil {
		h.headers = make(map[string]interface{})
	}

	h.headers[name] = value

	return h
}

// SetBody Set body values
func (h HTTPRequest) SetBody(body map[string]interface{}) HTTPRequest {
	if h.body == nil {
		h.body = make(map[string]interface{})
	}

	for index, value := range body {
		h.body[index] = value
	}

	return h
}

// SetQueryParams Set forms values
func (h HTTPRequest) SetQueryParams(queryParams map[string]interface{}) HTTPRequest {
	if h.queryParams == nil {
		h.queryParams = make(map[string]interface{})
	}

	for index, value := range queryParams {
		h.queryParams[index] = value
	}

	return h
}

// SetInsecureRequest Set if request is insecure (no cert validation)
func (h HTTPRequest) SetInsecureRequest(InsecureRequest bool) HTTPRequest {
	h.insecureRequest = InsecureRequest
	return h
}

// SetTimeOut Set request timeout
func (h HTTPRequest) SetTimeOut(timeOut int) HTTPRequest {
	h.timeOut = timeOut

	if h.timeOut < minTimeOut {
		h.timeOut = minTimeOut
	}

	if h.timeOut > maxTimeout {
		h.timeOut = maxTimeout
	}

	return h
}

// GetTimeOut Get request timeout
func (h HTTPRequest) GetTimeOut() int {
	return h.timeOut
}

// GetInsecureRequest Get if request is insecure
func (h HTTPRequest) GetInsecureRequest() bool {
	return h.insecureRequest
}

// ToGoHTTPRequest Create a go http.Request from a HTTPRequest
func (h HTTPRequest) ToGoHTTPRequest() (*http.Request, error) {

	body, err := json.Marshal(h.body)

	if err != nil {
		return nil, err
	}

	request, errReq := http.NewRequest(h.method, h.getURLWithQueryParans(), bytes.NewBuffer(body))

	if errReq != nil {
		return nil, errReq
	}

	for index, value := range h.headers {
		request.Header.Set(index, fmt.Sprintf("%v", value))
	}

	return request, nil
}

// Return the url with query parameters
func (h HTTPRequest) getURLWithQueryParans() (url string) {
	url = h.url

	if len(h.queryParams) == 0 {
		return
	}

	if !strings.Contains(url, "?") {
		url += "?"
	}

	for index, val := range h.queryParams {
		if v, ok := val.(bool); ok {
			if v {
				url += fmt.Sprintf("&%s=true", index)
			} else {
				url += fmt.Sprintf("&%s=false", index)
			}
		} else {
			url += fmt.Sprintf("&%s=%v", index, val)
		}
	}

	return
}

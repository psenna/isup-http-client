package isuphttp

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"
)

// HTTPClient HTTPClient
type HTTPClient struct {
	mockEnable   bool
	mockResponse map[string]HTTPResponse
}

// HTTPCall Make a http call
func (c HTTPClient) HTTPCall(request HTTPRequest) HTTPResponse {
	if c.mockEnable {
		return c.GetMockResponse(request.method, request.url)
	}

	return c.httpRequest(request)
}

func (c HTTPClient) httpRequest(request HTTPRequest) HTTPResponse {

	// request configuration
	goRequest, err := request.ToGoHTTPRequest()

	if err != nil {
		fmt.Println(err)
		return HTTPResponse{StatusCode: 0}
	}

	// Make request
	start := time.Now()

	response, err := c.getHTTPClient(request).Do(goRequest)

	elapsed := time.Since(start)

	if err != nil {
		return c.handleRequestError(err)
	}

	defer response.Body.Close()

	// Process response

	returnresponse := GetHTTPResponse(response)

	returnresponse.ResponseTime = float64(elapsed.Nanoseconds() / 1000000.0)

	return returnresponse
}

func (c HTTPClient) handleRequestError(err error) HTTPResponse {
	if strings.Contains(err.Error(), "(Client.Timeout exceeded while awaiting headers)") {
		return HTTPResponse{Error: err.Error(), StatusCode: 599}
	}

	return HTTPResponse{Error: err.Error()}
}

func (c HTTPClient) getHTTPClient(request HTTPRequest) *http.Client {
	timeout := time.Duration(request.GetTimeOut()) * time.Millisecond

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: request.GetInsecureRequest()},
		Dial: (&net.Dialer{
			Timeout:   timeout,
			KeepAlive: timeout,
		}).Dial,
		TLSHandshakeTimeout:   timeout,
		ResponseHeaderTimeout: timeout,
		ExpectContinueTimeout: timeout,
	}

	return &http.Client{
		Transport: tr,
		Timeout:   timeout,
	}
}

// AddMockResponse Add a mock response for a api call
func (c *HTTPClient) AddMockResponse(expectedResponse HTTPResponse, apiMethod string, apiURL string) {
	if c.mockResponse == nil {
		c.mockResponse = make(map[string]HTTPResponse)
	}

	c.mockResponse[apiMethod+"-"+apiURL] = expectedResponse
}

// GetMockResponse Get a mock response for a api call
func (c *HTTPClient) GetMockResponse(apiMethod string, apiURL string) HTTPResponse {
	if c.mockResponse == nil {
		return HTTPResponse{Method: apiMethod, URL: apiURL, StatusCode: 404}
	}

	if _, ok := c.mockResponse[apiMethod+"-"+apiURL]; !ok {
		return HTTPResponse{Method: apiMethod, URL: apiURL, StatusCode: 404}
	}

	return c.mockResponse[apiMethod+"-"+apiURL]
}

// SetMockEnable Enable or disable mock api call
func (c *HTTPClient) SetMockEnable(enable bool) {
	c.mockEnable = enable
}

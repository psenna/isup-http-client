package isuphttp_test

import (
	"net/http"
	"testing"

	"github.com/psenna/isup-http-client/isuphttp"
	"github.com/stretchr/testify/assert"
)

func TestGetMockResponse(t *testing.T) {
	var tests = []struct {
		setResponse      isuphttp.HTTPResponse
		expectedResponse isuphttp.HTTPResponse
		apiMethod        string
		apiURL           string
	}{
		{isuphttp.HTTPResponse{}, isuphttp.HTTPResponse{Method: isuphttp.GET, URL: "localhost:8080/api", StatusCode: 404}, isuphttp.GET, "localhost:8080/api"},
		{isuphttp.HTTPResponse{Method: isuphttp.POST, URL: "localhost:8080/api", StatusCode: 200}, isuphttp.HTTPResponse{Method: isuphttp.GET, URL: "localhost:8080/api", StatusCode: 404}, "GET", "localhost:8080/api"},
		{isuphttp.HTTPResponse{Method: isuphttp.GET, URL: "localhost:8080/v2/api", StatusCode: 200}, isuphttp.HTTPResponse{Method: isuphttp.GET, URL: "localhost:8080/api", StatusCode: 404}, "GET", "localhost:8080/api"},
		{isuphttp.HTTPResponse{Method: isuphttp.GET, URL: "localhost:8080/api", StatusCode: 200}, isuphttp.HTTPResponse{Method: isuphttp.GET, URL: "localhost:8080/api", StatusCode: 200}, "GET", "localhost:8080/api"},
	}

	for _, test := range tests {
		HTTPClient := isuphttp.HTTPClient{}

		if test.setResponse.URL != "" {
			HTTPClient.AddMockResponse(test.setResponse, test.setResponse.Method, test.setResponse.URL)
		}

		resultResponse := HTTPClient.GetMockResponse(test.apiMethod, test.apiURL)

		assert.True(t, assert.ObjectsAreEqualValues(test.expectedResponse, resultResponse), "The response object are different from the expected response")
	}
}

// Get a response with mock enable
func TestGetResponseMockEnable(t *testing.T) {
	var tests = []struct {
		setResponse      isuphttp.HTTPResponse
		expectedResponse isuphttp.HTTPResponse
		apiMethod        string
		apiURL           string
	}{
		{isuphttp.HTTPResponse{}, isuphttp.HTTPResponse{Method: isuphttp.GET, URL: "localhost:8080/api", StatusCode: 404}, isuphttp.GET, "localhost:8080/api"},
		{isuphttp.HTTPResponse{Method: isuphttp.POST, URL: "localhost:8080/api", StatusCode: 200}, isuphttp.HTTPResponse{Method: isuphttp.GET, URL: "localhost:8080/api", StatusCode: 404}, "GET", "localhost:8080/api"},
		{isuphttp.HTTPResponse{Method: isuphttp.GET, URL: "localhost:8080/v2/api", StatusCode: 200}, isuphttp.HTTPResponse{Method: isuphttp.GET, URL: "localhost:8080/api", StatusCode: 404}, "GET", "localhost:8080/api"},
		{isuphttp.HTTPResponse{Method: isuphttp.GET, URL: "localhost:8080/api", StatusCode: 200}, isuphttp.HTTPResponse{Method: isuphttp.GET, URL: "localhost:8080/api", StatusCode: 200}, "GET", "localhost:8080/api"},
	}

	for _, test := range tests {
		HTTPClient := isuphttp.HTTPClient{}
		HTTPClient.SetMockEnable(true)

		if test.setResponse.URL != "" {
			HTTPClient.AddMockResponse(test.setResponse, test.setResponse.Method, test.setResponse.URL)
		}

		resultResponse := HTTPClient.HTTPCall(isuphttp.GetHTTPRequest(test.apiMethod, test.apiURL))

		assert.True(t, assert.ObjectsAreEqualValues(test.expectedResponse, resultResponse), "The response object are different from the expected response")
	}
}

// Handle request timeout
// todo: make with httptest
func TestGetResponseFromApiWithTimeout(t *testing.T) {
	HTTPClient := isuphttp.HTTPClient{}

	httpRequest := isuphttp.GetHTTPRequest(isuphttp.GET, "https://www.ufms.br/")

	httpRequest = httpRequest.SetTimeOut(1)

	response := HTTPClient.HTTPCall(httpRequest)

	assert.Equal(t, isuphttp.StatusTimeout, response.StatusCode)

	assert.Equal(t, isuphttp.StatusText(isuphttp.StatusTimeout), response.Error)
}

// Get response from a api
// Not a perfect test, but work as long as github is up and you have internet :)
func TestGetResponseFromApi(t *testing.T) {
	HTTPClient := isuphttp.HTTPClient{}

	httpRequest := isuphttp.GetHTTPRequest(isuphttp.GET, "https://github.com/")

	response := HTTPClient.HTTPCall(httpRequest)

	assert.Equal(t, http.StatusOK, response.StatusCode)
}

// Todo: Invalid cert test
// https://github.com/golang/go/blob/968e18eebd736870a1e3bf06d941dc06e7b20688/src/net/http/client_test.go#L845

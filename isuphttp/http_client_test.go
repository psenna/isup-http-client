package isuphttp_test

import (
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
		{isuphttp.HTTPResponse{}, isuphttp.HTTPResponse{Method: "GET", URL: "localhost:8080/api", StatusCode: 404}, "GET", "localhost:8080/api"},
		{isuphttp.HTTPResponse{Method: "POST", URL: "localhost:8080/api", StatusCode: 200}, isuphttp.HTTPResponse{Method: "GET", URL: "localhost:8080/api", StatusCode: 404}, "GET", "localhost:8080/api"},
		{isuphttp.HTTPResponse{Method: "GET", URL: "localhost:8080/v2/api", StatusCode: 200}, isuphttp.HTTPResponse{Method: "GET", URL: "localhost:8080/api", StatusCode: 404}, "GET", "localhost:8080/api"},
		{isuphttp.HTTPResponse{Method: "GET", URL: "localhost:8080/api", StatusCode: 200}, isuphttp.HTTPResponse{Method: "GET", URL: "localhost:8080/api", StatusCode: 200}, "GET", "localhost:8080/api"},
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
		{isuphttp.HTTPResponse{}, isuphttp.HTTPResponse{Method: "GET", URL: "localhost:8080/api", StatusCode: 404}, "GET", "localhost:8080/api"},
		{isuphttp.HTTPResponse{Method: "POST", URL: "localhost:8080/api", StatusCode: 200}, isuphttp.HTTPResponse{Method: "GET", URL: "localhost:8080/api", StatusCode: 404}, "GET", "localhost:8080/api"},
		{isuphttp.HTTPResponse{Method: "GET", URL: "localhost:8080/v2/api", StatusCode: 200}, isuphttp.HTTPResponse{Method: "GET", URL: "localhost:8080/api", StatusCode: 404}, "GET", "localhost:8080/api"},
		{isuphttp.HTTPResponse{Method: "GET", URL: "localhost:8080/api", StatusCode: 200}, isuphttp.HTTPResponse{Method: "GET", URL: "localhost:8080/api", StatusCode: 200}, "GET", "localhost:8080/api"},
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

// Get response from a api
// Not a perfect test, but work as long as github is up and you have internet :)
func TestGetResponseFromApi(t *testing.T) {
	HTTPClient := isuphttp.HTTPClient{}

	httpRequest := isuphttp.GetHTTPRequest("GET", "https://github.com/")

	response := HTTPClient.HTTPCall(httpRequest)

	assert.Equal(t, 200, response.StatusCode)
}

// Handle request timeout
func TestGetResponseFromApiWithTimeout(t *testing.T) {
	HTTPClient := isuphttp.HTTPClient{}

	httpRequest := isuphttp.GetHTTPRequest("GET", "https://github.com/")

	httpRequest = httpRequest.SetTimeOut(1)

	response := HTTPClient.HTTPCall(httpRequest)

	assert.Equal(t, 599, response.StatusCode)
}

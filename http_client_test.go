package isup_http_test

import (
	"testing"

	"github.com/psenna/isup-http-client/isup_http"
	"github.com/stretchr/testify/assert"
)

func TestGetMockResponse(t *testing.T) {
	var tests = []struct {
		setResponse      isup_http.HTTPResponse
		expectedResponse isup_http.HTTPResponse
		apiMethod        string
		apiURL           string
	}{
		{isup_http.HTTPResponse{}, isup_http.HTTPResponse{Method: "GET", URL: "localhost:8080/api", StatusCode: 404}, "GET", "localhost:8080/api"},
		{isup_http.HTTPResponse{Method: "POST", URL: "localhost:8080/api", StatusCode: 200}, isup_http.HTTPResponse{Method: "GET", URL: "localhost:8080/api", StatusCode: 404}, "GET", "localhost:8080/api"},
		{isup_http.HTTPResponse{Method: "GET", URL: "localhost:8080/v2/api", StatusCode: 200}, isup_http.HTTPResponse{Method: "GET", URL: "localhost:8080/api", StatusCode: 404}, "GET", "localhost:8080/api"},
		{isup_http.HTTPResponse{Method: "GET", URL: "localhost:8080/api", StatusCode: 200}, isup_http.HTTPResponse{Method: "GET", URL: "localhost:8080/api", StatusCode: 200}, "GET", "localhost:8080/api"},
	}

	for _, test := range tests {
		HTTPClient := isup_http.HTTPClient{}

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
		setResponse      isup_http.HTTPResponse
		expectedResponse isup_http.HTTPResponse
		apiMethod        string
		apiURL           string
	}{
		{isup_http.HTTPResponse{}, isup_http.HTTPResponse{Method: "GET", URL: "localhost:8080/api", StatusCode: 404}, "GET", "localhost:8080/api"},
		{isup_http.HTTPResponse{Method: "POST", URL: "localhost:8080/api", StatusCode: 200}, isup_http.HTTPResponse{Method: "GET", URL: "localhost:8080/api", StatusCode: 404}, "GET", "localhost:8080/api"},
		{isup_http.HTTPResponse{Method: "GET", URL: "localhost:8080/v2/api", StatusCode: 200}, isup_http.HTTPResponse{Method: "GET", URL: "localhost:8080/api", StatusCode: 404}, "GET", "localhost:8080/api"},
		{isup_http.HTTPResponse{Method: "GET", URL: "localhost:8080/api", StatusCode: 200}, isup_http.HTTPResponse{Method: "GET", URL: "localhost:8080/api", StatusCode: 200}, "GET", "localhost:8080/api"},
	}

	for _, test := range tests {
		HTTPClient := isup_http.HTTPClient{}
		HTTPClient.SetMockEnable(true)

		if test.setResponse.URL != "" {
			HTTPClient.AddMockResponse(test.setResponse, test.setResponse.Method, test.setResponse.URL)
		}

		resultResponse := HTTPClient.HTTPCall(isup_http.GetHTTPRequest(test.apiMethod, test.apiURL))

		assert.True(t, assert.ObjectsAreEqualValues(test.expectedResponse, resultResponse), "The response object are different from the expected response")
	}
}

// Get response from a api
// Not a perfect test, but work
func TestGetResponseFromApi(t *testing.T) {
	HTTPClient := isup_http.HTTPClient{}

	httpRequest := isup_http.GetHTTPRequest("GET", "https://github.com/")

	response := HTTPClient.HTTPCall(httpRequest)

	assert.Equal(t, 200, response.StatusCode)
}

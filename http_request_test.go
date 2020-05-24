package isuphttp_test

import (
	"testing"

	"github.com/psenna/isup-http-client/isuphttp"
	"github.com/stretchr/testify/assert"
)

const (
	timeOut    = 2000
	minTimeOut = 50
	maxTimeout = 60000
)

func TestGetHTTPRequestTimeOut(t *testing.T) {
	var tests = []struct {
		apiMethod       string
		apiURL          string
		timeOut         int
		expectedTimeout int
	}{
		{"GET", "localhost:8080/api", 1, minTimeOut},
		{"GET", "localhost:8080/api", 50, 50},
		{"GET", "localhost:8080/api", 0, timeOut},
		{"GET", "localhost:8080/api", 60000, 60000},
		{"GET", "localhost:8080/api", 60001, maxTimeout},
	}

	for _, test := range tests {
		request := isuphttp.GetHTTPRequest(test.apiMethod, test.apiURL)

		if test.timeOut != 0 {
			request = request.SetTimeOut(test.timeOut)
		}

		assert.Equal(t, test.expectedTimeout, request.GetTimeOut())
	}
}

func TestGetHTTPRequestHeaders(t *testing.T) {
	var tests = []struct {
		apiMethod       string
		apiURL          string
		headers         map[string]interface{}
		expectedHeaders map[string]string
	}{
		{"GET", "localhost:8080/api", map[string]interface{}{"value1": 12345}, map[string]string{"value1": "12345"}},
		{"GET", "localhost:8080/api", map[string]interface{}{"value1": "qqCoisa", "1234": true}, map[string]string{"value1": "qqCoisa", "1234": "true"}},
		{"GET", "localhost:8080/api", map[string]interface{}{"value1": "qqCoisa", "1234": true, "asdasd": 11.6}, map[string]string{"value1": "qqCoisa", "1234": "true", "asdasd": "11.6"}},
	}

	for _, test := range tests {
		request := isuphttp.GetHTTPRequest(test.apiMethod, test.apiURL)

		request = request.SetHeaders(test.headers)

		requestGo, _ := request.ToGoHTTPRequest()

		for index, value := range test.expectedHeaders {
			assert.Equal(t, value, requestGo.Header.Get(index))
		}

	}
}

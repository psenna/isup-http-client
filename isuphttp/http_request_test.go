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

func TestSetHTTPRequestTimeOut(t *testing.T) {
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

func TestSetHTTPRequestHeaders(t *testing.T) {
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

func TestHTTPRequestSetHeaderValue(t *testing.T) {
	var tests = []struct {
		apiMethod       string
		apiURL          string
		baseHeaders     map[string]interface{}
		headersToSet    map[string]interface{}
		expectedHeaders map[string]string
	}{
		{"GET", "localhost:8080/api", map[string]interface{}{}, map[string]interface{}{"value1": 12345}, map[string]string{"value1": "12345"}},
		{"GET", "localhost:8080/api", map[string]interface{}{"value1": 12345}, map[string]interface{}{}, map[string]string{"value1": "12345"}},
		{"GET", "localhost:8080/api", map[string]interface{}{"value1": "qqCoisa", "1234": true}, map[string]interface{}{"value2": "asdasd"}, map[string]string{"value1": "qqCoisa", "value2": "asdasd", "1234": "true"}},
		{"GET", "localhost:8080/api", map[string]interface{}{"value1": "qqCoisa", "1234": true, "asdasd": 11.6}, map[string]interface{}{"asdasd": 1237}, map[string]string{"value1": "qqCoisa", "1234": "true", "asdasd": "1237"}},
	}

	for _, test := range tests {
		request := isuphttp.GetHTTPRequest(test.apiMethod, test.apiURL)

		if len(test.baseHeaders) > 0 {
			request = request.SetHeaders(test.baseHeaders)
		}

		for name, value := range test.headersToSet {
			request = request.SetHeaderValue(name, value)
		}

		requestGo, _ := request.ToGoHTTPRequest()

		for index, value := range test.expectedHeaders {
			assert.Equal(t, value, requestGo.Header.Get(index))
		}

	}
}

func TestHTTPRequestSetHeaderContentType(t *testing.T) {
	var tests = []struct {
		apiMethod       string
		apiURL          string
		baseHeaders     map[string]interface{}
		content         string
		expectedHeaders map[string]string
	}{
		{"GET", "localhost:8080/api", map[string]interface{}{}, isuphttp.ApplicationJSON, map[string]string{"Content-type": isuphttp.ApplicationJSON}},
		{"GET", "localhost:8080/api", map[string]interface{}{"App": "123"}, isuphttp.ApplicationJSON, map[string]string{"Content-type": isuphttp.ApplicationJSON, "App": "123"}},
		{"GET", "localhost:8080/api", map[string]interface{}{"Content-type": "PDF"}, isuphttp.ApplicationXML, map[string]string{"Content-type": isuphttp.ApplicationXML}},
		{"GET", "localhost:8080/api", map[string]interface{}{"Content-type": "PDF", "App": "123"}, isuphttp.ApplicationXML, map[string]string{"Content-type": isuphttp.ApplicationXML, "App": "123"}},
	}

	for _, test := range tests {
		request := isuphttp.GetHTTPRequest(test.apiMethod, test.apiURL)

		if len(test.baseHeaders) > 0 {
			request = request.SetHeaders(test.baseHeaders)
		}

		request = request.SetContentType(test.content)

		requestGo, _ := request.ToGoHTTPRequest()

		for index, value := range test.expectedHeaders {
			assert.Equal(t, value, requestGo.Header.Get(index))
		}

	}
}

func TestHTTPRequestSetHeaderAccept(t *testing.T) {
	var tests = []struct {
		apiMethod       string
		apiURL          string
		baseHeaders     map[string]interface{}
		accept          string
		expectedHeaders map[string]string
	}{
		{"GET", "localhost:8080/api", map[string]interface{}{}, isuphttp.ApplicationJSON, map[string]string{"Accept": isuphttp.ApplicationJSON}},
		{"GET", "localhost:8080/api", map[string]interface{}{"App": "123"}, isuphttp.ApplicationJSON, map[string]string{"Accept": isuphttp.ApplicationJSON, "App": "123"}},
		{"GET", "localhost:8080/api", map[string]interface{}{"Accept": "PDF"}, isuphttp.ApplicationXML, map[string]string{"Accept": isuphttp.ApplicationXML}},
		{"GET", "localhost:8080/api", map[string]interface{}{"Accept": "PDF", "App": "123"}, isuphttp.ApplicationXML, map[string]string{"Accept": isuphttp.ApplicationXML, "App": "123"}},
	}

	for _, test := range tests {
		request := isuphttp.GetHTTPRequest(test.apiMethod, test.apiURL)

		if len(test.baseHeaders) > 0 {
			request = request.SetHeaders(test.baseHeaders)
		}

		request = request.SetAccept(test.accept)

		requestGo, _ := request.ToGoHTTPRequest()

		for index, value := range test.expectedHeaders {
			assert.Equal(t, value, requestGo.Header.Get(index))
		}

	}
}

func TestHTTPRequestSetHeaderAuthorization(t *testing.T) {
	var tests = []struct {
		apiMethod       string
		apiURL          string
		baseHeaders     map[string]interface{}
		authorization   string
		expectedHeaders map[string]string
	}{
		{"GET", "localhost:8080/api", map[string]interface{}{}, "Bearer AASASaSdasdasdadsasdasdSDasD", map[string]string{"Authorization": "Bearer AASASaSdasdasdadsasdasdSDasD"}},
		{"GET", "localhost:8080/api", map[string]interface{}{"App": "123"}, "Basic KJLHlkJklJkljoiasdjdoaisj", map[string]string{"Authorization": "Basic KJLHlkJklJkljoiasdjdoaisj", "App": "123"}},
		{"GET", "localhost:8080/api", map[string]interface{}{"Authorization": "PDF"}, "Bearer AASASaSdasdasdadsasdasdSDasD", map[string]string{"Authorization": "Bearer AASASaSdasdasdadsasdasdSDasD"}},
		{"GET", "localhost:8080/api", map[string]interface{}{"Authorization": "PDF", "App": "123"}, "Basic KJLHlkJklJkljoiasdjdoaisj", map[string]string{"Authorization": "Basic KJLHlkJklJkljoiasdjdoaisj", "App": "123"}},
	}

	for _, test := range tests {
		request := isuphttp.GetHTTPRequest(test.apiMethod, test.apiURL)

		if len(test.baseHeaders) > 0 {
			request = request.SetHeaders(test.baseHeaders)
		}

		request = request.SetAuthorization(test.authorization)

		requestGo, _ := request.ToGoHTTPRequest()

		for index, value := range test.expectedHeaders {
			assert.Equal(t, value, requestGo.Header.Get(index))
		}

	}
}

package isuphttp

import (
	"io/ioutil"
	"net/http"
)

// HTTPResponse A response from a http call
type HTTPResponse struct {
	URL           string
	Method        string
	StatusCode    int
	Body          string
	ResponseTime  float64
	ContentLength int64
	ContentType   string
	Error         string
	Headers       map[string]interface{}
}

// GetHTTPResponse Instantiate a HTTP request object
func GetHTTPResponse(response *http.Response) HTTPResponse {

	bodyBytes, err := ioutil.ReadAll(response.Body)
	bodyString := ""

	if err == nil {
		bodyString = string(bodyBytes)
	}

	h := HTTPResponse{
		URL:           response.Request.URL.Hostname(),
		Method:        response.Request.Method,
		StatusCode:    response.StatusCode,
		Body:          bodyString,
		ContentLength: response.ContentLength,
	}

	return h
}

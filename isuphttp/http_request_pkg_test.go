package isuphttp

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetHTTPRequestQueryParameters(t *testing.T) {
	var tests = []struct {
		apiMethod          string
		apiURL             string
		queryParameters    map[string]interface{}
		expectedParameters map[string]string
	}{
		{"GET", "localhost:8080/api", map[string]interface{}{"value1": 12345}, map[string]string{"value1": "12345"}},
		{"GET", "localhost:8080/api", map[string]interface{}{"value1": true}, map[string]string{"value1": "true"}},
		{"GET", "localhost:8080/api", map[string]interface{}{"value1": "qqCoisa", "1234": true}, map[string]string{"value1": "qqCoisa", "1234": "true"}},
		{"GET", "localhost:8080/api", map[string]interface{}{"value1": "qqCoisa", "1234": true, "asdasd": 11.6}, map[string]string{"value1": "qqCoisa", "1234": "true", "asdasd": "11.6"}},
	}

	for _, test := range tests {
		request := GetHTTPRequest(test.apiMethod, test.apiURL)

		request = request.SetQueryParams(test.queryParameters)

		url := request.getURLWithQueryParans()

		assert.True(
			t,
			strings.Contains(url, test.apiURL),
			fmt.Sprintf("Url (%s) \nshould contain api url %s", url, test.apiURL),
		)

		for index, value := range test.expectedParameters {
			assert.True(
				t,
				strings.Contains(url, index+"="+value),
				fmt.Sprintf("Url (%s) \nshould contain query parameter %s", url, index+"="+value),
			)
		}

	}
}

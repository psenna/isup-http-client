package isuphttp

// HTTP status code for request errors
const (
	StatusTimeout     = 1 // Request Timeout
	StatusInvalidCert = 2 // Invalid SSL Certificate
)

var statusText = map[int]string{
	StatusTimeout:     "Request Timeout",
	StatusInvalidCert: "Invalid SSL Certificate",
}

// StatusText returns a text for the HTTP errors status code. It returns the empty
// string if the code is unknown.
func StatusText(code int) string {
	return statusText[code]
}

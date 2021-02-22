package paystack

import (
	"encoding/json"
	"net/http"
	"net/url"
)

// APIError includes the response from the Paystack API and some HTTP request info
type APIError struct {
	Message        string        `json:"message,omitempty"`
	HTTPStatusCode int           `json:"code,omitempty"`
	Details        ErrorResponse `json:"details,omitempty"`
	URL            *url.URL      `json:"url,omitempty"`
	Header         http.Header   `json:"header,omitempty"`
}

// APIError supports the error interface
func (e *APIError) Error() string {
	ret, _ := json.Marshal(e)
	return string(ret)
}

// ErrorResponse represents an error response from the Paystack API server
type ErrorResponse struct {
	Status  bool                   `json:"status,omitempty"`
	Message string                 `json:"message,omitempty"`
	Errors  map[string]interface{} `json:"errors,omitempty"`
}

func newAPIError(httpResponse *http.Response, responseBuffer []byte) *APIError {
	var paystackErrorResp ErrorResponse
	// Todo: don't unmarshal again since, it's already been done in the caller
	_ = json.Unmarshal(responseBuffer, &paystackErrorResp)
	return &APIError{
		Message:        paystackErrorResp.Message,
		HTTPStatusCode: httpResponse.StatusCode,
		Header:         httpResponse.Header,
		Details:        paystackErrorResp,
		URL:            httpResponse.Request.URL,
	}
}

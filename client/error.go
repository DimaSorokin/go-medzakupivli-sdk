package client

import "fmt"

// ErrorResponse represents a JSON error response from a MOZ API
// endpoint.
type ErrorResponse struct {
	StatusCode int    `json:"statusCode"`
	ErrorCode  string `json:"errorCode"`
	Message    string `json:"message"`
}

func (err *ErrorResponse) Error() string {
	if err.ErrorCode == "" {
		return fmt.Sprintf("%d API error: %s", err.StatusCode, err.Message)
	}
	return fmt.Sprintf("%d (%s) API error: %s", err.StatusCode, err.ErrorCode, err.Message)
}

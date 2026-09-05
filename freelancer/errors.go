package freelancer

import (
	"errors"
	"fmt"
)

type APIError struct {
	StatusCode int    `json:"-"`
	Status     string `json:"status"`
	Message    string `json:"message"`
	RequestID  string `json:"request_id"`

	InnerError struct {
		Code     string `json:"code"`
		Detail   string `json:"detail"`
		HTTPCode int    `json:"http_code"`
		Source   string `json:"source"`
	} `json:"error"`

	LegacyErrorCode string `json:"error_code"`

	RawPayload []byte `json:"_"`

	Meta *ResponseMeta `json:"-"`
}

func (e *APIError) Error() string {
	// if the inner error is set, use it
	msg := e.Message
	if e.InnerError.Detail != "" {
		msg = e.InnerError.Detail
	}

	code := e.InnerError.Code
	if code == "" {
		code = e.LegacyErrorCode
	}

	return fmt.Sprintf("freelancer api (%d): [%s] %s (request_id: %s)", e.StatusCode, code, msg, e.RequestID)

}

// TODO: add to changelog
// IsAPIError returns the *APIError and a boolean if the error is an API-specific failure.
func IsAPIError(err error) (*APIError, bool) {
	var apiErr *APIError
	if errors.As(err, &apiErr) {
		return apiErr, true
	}
	return nil, false
}

package freelancer

import "fmt"

type APIError struct {
	StatusCode int    `json:"-"` // will set this manually from the HTTP
	Status     string `json:"status"`
	Message    string `json:"message"`
	RequestID  string `json:"request_id"`

	InnerError struct {
		Code     string `json:"code"`
		Detail   string `json:"detail"`
		HTTPCode int    `json:"http_code"`
		Source   string `json:"source"`
	} `json:"error"`

	RawPayload []byte `json:"-"`

	LegacyErrorCode string `json:"error_code"`
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

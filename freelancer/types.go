package freelancer

import (
	"fmt"
)

type APIError struct {
	Code     int64  `json:"code"`
	Message  string `json:"message"`
	Response []byte `json:"-"`
}

type APIError2 struct {
	Status    string `json:"status"`
	Message   string `json:"message"`
	ErrorCode string `json:"error_code"`
	ErrorObj  error2 `json:"error"`
	RequestID string `json:"request_id"`
	Response  []byte `json:"-"`
}

type error2 struct {
	Code     string `json:"code"`
	Detail   string `json:"detail"`
	HTTPCode int    `json:"http_code"`
	Source   string `json:"source"`
}

func (e APIError) Error() string {
	if e.IsValid() {
		return fmt.Sprintf("<APIError> code: %d, message: %s", e.Code, e.Message)
	}
	return fmt.Sprintf("<APIError> response: %s", string(e.Response))
}

func (e APIError2) Error() string {
	return fmt.Sprintf("<APIError2> status: %s, message: %s", e.Status, e.Message)
}

func (e APIError2) IsValid() bool {
	return e.Status != "" || e.Message != ""
}

func (e APIError) IsValid() bool {
	return e.Code != 0 || e.Message != ""
}

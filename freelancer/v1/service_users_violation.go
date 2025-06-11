package freelancer

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

// Create a user violation report
type ReportUserViolationService struct {
	client *Client
}

type ReportUserViolationRequestBody struct {
	ContextID        int                       `json:"context_id"`
	ContextType      ViolationContext          `json:"context_type"`
	ViolatorUserID   int64                     `json:"violator_user_id"`
	Reason           ViolationReason           `json:"reason"`
	AdditionalReason ViolationAdditionalReason `json:"additional_reason,omitempty"`
	Comments         string                    `json:"comments,omitempty"`
	Url              string                    `json:"url,omitempty"`
}

func (s *ReportUserViolationService) DO(ctx context.Context, b ReportUserViolationRequestBody) (*BaseResponse, error) {
	m, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}

	r := &request{
		method:   http.MethodPost,
		endpoint: string(USERS_VIOLATION),
		body:     bytes.NewBuffer(m),
	}

	data, err := s.client.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}

	resp := &BaseResponse{}
	if err := json.Unmarshal(data, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

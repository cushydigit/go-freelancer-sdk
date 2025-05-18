package freelancer

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

type ReportUserViolationService struct {
	client *Client
}

type ReportUserViolationRequestBody struct {
	ContextID        int                           `json:"context_id"`
	ContextType      ViolationContextType          `json:"context_type"`
	ViolatorUserID   int64                         `json:"violator_user_id"`
	Reason           ViolationReasonType           `json:"reason"`
	AdditionalReason ViolationAdditionalReasonType `json:"additional_reason"`
	Comments         string                        `json:"comments"`
	Url              string                        `json:"url"`
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

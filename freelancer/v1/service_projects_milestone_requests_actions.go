package freelancer

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type milestoneRequestsActionsService struct {
	client *Client
}

type ActionMilestoneRequestBody struct {
	Action MilestoneActionRequest `json:"action"`
}

func (s *milestoneRequestsActionsService) Do(ctx context.Context, milestoneRequestID int, b ActionMilestoneRequestBody) (*BaseResponse, error) {
	m, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodPut,
		endpoint: fmt.Sprintf("%s/%d", string(PROJECTS_MILESTONES), milestoneRequestID),
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

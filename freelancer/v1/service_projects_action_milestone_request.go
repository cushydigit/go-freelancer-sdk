package freelancer

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type actionMilestonRequestService struct {
	client *Client
}

type ActionMilestoneReqeustRequestBody struct {
	Action MilestoneActionRequest `json:"action"`
}

func (s *actionMilestonRequestService) Do(ctx context.Context, milestoneReqeustID int, b ActionMilestoneRequestBody) (*BaseResponse, error) {
	m, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodPut,
		endpoint: fmt.Sprintf("%s/%d", string(PROJECTS_MILESTONES), milestoneReqeustID),
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

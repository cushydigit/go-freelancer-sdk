package freelancer

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

type milestoneRequestsCreateService struct {
	client *Client
}

// ProjectID, BidID, Amount and Description are required
type CreateMilestoneRequestBody struct {
	ProjectID   int64  `json:"project_id"`
	BidID       int    `json:"bid_id"`
	Amount      int    `json:"amount"`
	Description string `json:"description"`
}

func (s *milestoneRequestsCreateService) Do(ctx context.Context, b CreateMilestoneBody) (*BaseResponse, error) {
	m, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodPost,
		endpoint: string(PROJECTS_MILESTONE_REQUESTS),
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

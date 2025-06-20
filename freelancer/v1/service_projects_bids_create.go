package freelancer

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

type bidsCreateService struct {
	client *Client
}

// ProjectID, BidderID, Amount, Period (days), MilestonePercentage (0-100) required
type CreateBidBody struct {
	ProjectID           int64  `json:"project_id"`
	BidderID            int64  `json:"bidder_id"`
	Amount              int    `json:"amount"`
	Period              int    `json:"period"`
	MilestonePercentage int    `json:"milestone_percentage"`
	Description         string `json:"description"`
	ProfileID           int    `json:"profile_id"`
}

func (s *bidsCreateService) Do(ctx context.Context, b CreateBidBody) (*BaseResponse, error) {
	m, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodPost,
		endpoint: string(PROJECTS_BIDS),
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

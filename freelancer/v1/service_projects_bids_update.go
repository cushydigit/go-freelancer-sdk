package freelancer

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type bidsUpdateService struct {
	client *Client
}

type UpdateBidBody struct {
	Amount              int    `json:"amount"`
	MilestonePercentage int    `json:"milestone_percentage"`
	Description         string `json:"description"`
}

func (s *bidsUpdateService) Do(ctx context.Context, bidID int, b UpdateBidBody) (*BaseResponse, error) {
	m, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodPut,
		endpoint: fmt.Sprintf("%s/%d", string(PROJECTS_BIDS), bidID),
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

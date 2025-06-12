package freelancer

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

type bidEditRequestsCreateService struct {
	client *Client
}

// BidID, NewAmount, and Period (days) are required
type CreateBidEditRequestsBody struct {
	BidID     int `json:"bid_id"`
	NewAmount int `json:"new_amount"`
	NewPeriod int `json:"new_period"`
	Comment   int `json:"comment"`
}

func (s *bidEditRequestsCreateService) Do(ctx context.Context, b CreateBidEditRequestsBody) (*BaseResponse, error) {
	m, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodPost,
		endpoint: string(PROJECTS_BIDS_EDIT_REQUESTS),
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

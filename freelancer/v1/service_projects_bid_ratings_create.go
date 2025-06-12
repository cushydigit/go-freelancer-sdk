package freelancer

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type bidRatingsCreateService struct {
	client *Client
}

// Rating required
type CreateBidRatingsBody struct {
	Rating  int    `json:"rating"`
	Comment string `json:"comment"`
}

func (s *bidRatingsCreateService) Do(ctx context.Context, bidID int, b CreateBidRatingsBody) (*BaseResponse, error) {
	m, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodPost,
		endpoint: fmt.Sprintf("%s/%d/bid_ratings", string(PROJECTS_BIDS), bidID),
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

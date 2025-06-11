package freelancer

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type createBidRatingService struct {
	client *Client
}

// Rating required
type CreateBidRatingRequestBody struct {
	Rating  int    `json:"rating"`
	Comment string `json:"comment"`
}

func (s *createBidRatingService) Do(ctx context.Context, bidID int, b CreateBidRatingRequestBody) (*BaseResponse, error) {
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

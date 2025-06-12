package freelancer

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type bidsRatingsUpdateService struct {
	client *Client
}

type UpdateBidRatingBody struct {
	Rating  int    `json:"rating"`
	Comment string `json:"comment"`
}

func (s *bidsRatingsUpdateService) Do(ctx context.Context, bidID int, bidRatingID int, b UpdateBidRatingBody) (*BaseResponse, error) {
	m, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodPut,
		endpoint: fmt.Sprintf("%s/%d/bid_ratings/%d", string(PROJECTS_BIDS), bidID, bidRatingID),
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

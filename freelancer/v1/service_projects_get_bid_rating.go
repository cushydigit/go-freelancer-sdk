package freelancer

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type getBidRatingService struct {
	client *Client
}

func (s *getBidRatingService) Do(ctx context.Context, bidID int) (*BaseResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("%s/%d/bid_ratings", string(PROJECTS_BIDS), bidID),
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

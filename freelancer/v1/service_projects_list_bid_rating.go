package freelancer

import (
	"context"
	"encoding/json"
	"net/http"
)

type listBidRatingService struct {
	client *Client
	bids   []int
}

func (s *listBidRatingService) Do(ctx context.Context, bidID int) (*BaseResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: string(PROJECTS_BID_RATINGS),
	}

	for _, val := range s.bids {
		r.addParam("bids[]", val)
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

func (s *listBidRatingService) SetBids(vals []int) *listBidRatingService {
	s.bids = vals
	return s
}

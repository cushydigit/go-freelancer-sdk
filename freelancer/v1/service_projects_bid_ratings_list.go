package freelancer

import (
	"context"
	"encoding/json"
	"net/http"
)

type bidsRatingsListService struct {
	client *Client
	bids   []int
}

func (s *bidsRatingsListService) Do(ctx context.Context, bidID int) (*BaseResponse, error) {
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

func (s *bidsRatingsListService) SetBids(vals []int) *bidsRatingsListService {
	s.bids = vals
	return s
}

package freelancer

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type bidEditRequestsGetService struct {
	client            *Client
	statuses          []BidStatus
	bidEditRequestIDs []int
}

func (s *bidEditRequestsGetService) Do(ctx context.Context, bidID int) (*BaseResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("%s/%d/bid_edit_requests", string(PROJECTS_BIDS), bidID),
	}

	for _, val := range s.statuses {
		r.addParam("statuses[]", val)
	}
	for _, val := range s.bidEditRequestIDs {
		r.addParam("bid_edit_request_ids[]", val)
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

func (s *bidEditRequestsGetService) SetStatuses(vals []BidStatus) *bidEditRequestsGetService {
	s.statuses = vals
	return s
}

func (s *bidEditRequestsGetService) setBidEditRequestIDs(vals []int) *bidEditRequestsGetService {
	s.bidEditRequestIDs = vals
	return s
}

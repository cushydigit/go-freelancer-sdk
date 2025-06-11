package freelancer

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type getBidEditRequestService struct {
	client           *Client
	statuses         []BidStatus
	bidEditRequetIDs []int
}

func (s *getBidEditRequestService) Do(ctx context.Context, bidID int) (*BaseResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("%s/%d/bid_edit_requests", string(PROJECTS_BIDS), bidID),
	}

	for _, val := range s.statuses {
		r.addParam("statuses[]", val)
	}
	for _, val := range s.bidEditRequetIDs {
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

func (s *getBidEditRequestService) SetStatuses(vals []BidStatus) *getBidEditRequestService {
	s.statuses = vals
	return s
}

func (s *getBidEditRequestService) setBidEditReqeuestIDs(vals []int) *getBidEditRequestService {
	s.bidEditRequetIDs = vals
	return s
}

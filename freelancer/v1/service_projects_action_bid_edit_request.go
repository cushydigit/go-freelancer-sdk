package freelancer

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type actionBidEditRequestService struct {
	client *Client
}

type ActionBidEditRequestRequestBody struct {
	Action ActionBidEditRequest `json:"action"`
}

func (s *actionBidEditRequestService) Do(ctx context.Context, bidID int, bidEditRequestID int, b ActionBidEditRequestRequestBody) (*BaseResponse, error) {
	m, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodPut,
		endpoint: fmt.Sprintf("%s/%d/bid_edit_requests/%d", string(PROJECTS_BIDS), bidID, bidEditRequestID),
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

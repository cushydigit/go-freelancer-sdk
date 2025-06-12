package freelancer

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type bidEditRequestsActionService struct {
	client *Client
}

type ActionBidEditRequestsBody struct {
	Action ActionBidEditRequest `json:"action"`
}

func (s *bidEditRequestsActionService) Do(ctx context.Context, bidID int, bidEditRequestID int, b ActionBidEditRequestsBody) (*BaseResponse, error) {
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

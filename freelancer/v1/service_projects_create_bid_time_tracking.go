package freelancer

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type createTimeTrackingBidService struct {
	client *Client
}

// StartTime and Seconds (Duration of session in seconds) required
type CreateTimeTrackingBidRequestBody struct {
	StartTime int64  `json:"time_start"`
	Seconds   int    `json:"seconds"`
	Note      string `json:"note"`
}

func (s *createTimeTrackingBidService) Do(ctx context.Context, bidID int, b CreateTimeTrackingBidRequestBody) (*BaseResponse, error) {
	m, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodPost,
		endpoint: fmt.Sprintf("%s/%d/time_tracking_sessions", string(PROJECTS_BIDS), bidID),
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

package freelancer

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type updateBidService struct {
	client *Client
}

type UpdateBidRequestBody struct {
	Amount            int    `json:"amount"`
	MiestonePecentage int    `json:"milestone_percentage"`
	Description       string `json:"description"`
}

func (s *updateBidService) Do(ctx context.Context, bidID int, b UpdateBidRequestBody) (*BaseResponse, error) {
	m, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodPut,
		endpoint: fmt.Sprintf("%s/%d", string(PROJECTS_BIDS), bidID),
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

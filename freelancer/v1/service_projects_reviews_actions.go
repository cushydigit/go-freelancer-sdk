package freelancer

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type reviewActionService struct {
	client *Client
}

type ActionReviewBody struct {
	Action     ActionBid  `json:"action"`
	ReviewType ReviewType `json:"review_type"`
}

func (s *reviewActionService) Do(ctx context.Context, Id int64, b ActionReviewBody) (*BaseResponse, error) {
	m, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodPut,
		endpoint: fmt.Sprintf("%s/%s", string(PROJECTS_REVIEWS), strconv.FormatInt(Id, 10)),
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

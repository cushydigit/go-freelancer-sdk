package freelancer

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

type JobsBody struct {
	Jobs []int32 `json:"jobs[]"`
}

type selfJobsAddService struct {
	client *Client
}

func (s *selfJobsAddService) DO(ctx context.Context, b JobsBody) (*BaseResponse, error) {
	marshaledJson, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}

	r := &request{
		method:   http.MethodPost,
		endpoint: string(USERS_SELF_JOBS),
		body:     bytes.NewBuffer(marshaledJson),
	}

	data, err := s.client.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := &BaseResponse{}
	if err := json.Unmarshal(data, res); err != nil {
		return nil, err
	}
	return res, nil
}

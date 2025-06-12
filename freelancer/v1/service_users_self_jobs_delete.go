package freelancer

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

type selfJobsDeleteService struct {
	client *Client
}

func (s *selfJobsDeleteService) DO(ctx context.Context, b JobsBody) (*BaseResponse, error) {
	marshaledJson, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodDelete,
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

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

type addJobs struct {
	client *Client
}

func (s *addJobs) DO(ctx context.Context, b JobsBody) (*BaseResponse, error) {
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

type updateJobs struct {
	client *Client
}

func (s *updateJobs) Do(ctx context.Context, b JobsBody) (*BaseResponse, error) {
	marshaledJson, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodPut,
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

type deleteJobs struct {
	client *Client
}

func (s *deleteJobs) DO(ctx context.Context, b JobsBody) (*BaseResponse, error) {
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

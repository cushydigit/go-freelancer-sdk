package freelancer

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

type JobsRequestBody struct {
	Jobs []int32 `json:"jobs[]"`
}

// Add a list of jobs to the job list of current User
type AddSelfJobsService struct {
	client *Client
}

func (s *AddSelfJobsService) DO(ctx context.Context, b JobsRequestBody) (*BaseResponse, error) {
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

type UpdateSelfJobsService struct {
	client *Client
}

func (s *UpdateSelfJobsService) DO(ctx context.Context, b JobsRequestBody) (*BaseResponse, error) {
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

type DeleteSelfJobsService struct {
	client *Client
}

func (s *DeleteSelfJobsService) DO(ctx context.Context, b JobsRequestBody) (*BaseResponse, error) {
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

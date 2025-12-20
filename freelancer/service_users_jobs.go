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

// TODO: refine with typed response
func (s *addJobs) Do(ctx context.Context, b JobsBody) (*RawResponse, error) {
	marshaledJson, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}

	r := &request{
		method:   http.MethodPost,
		endpoint: string(USERS_SELF_JOBS),
		body:     bytes.NewBuffer(marshaledJson),
	}
	return execute[*RawResponse](ctx, s.client, r)
}

type updateJobs struct {
	client *Client
}

func (s *updateJobs) Do(ctx context.Context, b JobsBody) (*RawResponse, error) {
	marshaledJson, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodPut,
		endpoint: string(USERS_SELF_JOBS),
		body:     bytes.NewBuffer(marshaledJson),
	}
	return execute[*RawResponse](ctx, s.client, r)
}

type deleteJobs struct {
	client *Client
}

func (s *deleteJobs) DO(ctx context.Context, b JobsBody) (*RawResponse, error) {
	marshaledJson, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodDelete,
		endpoint: string(USERS_SELF_JOBS),
		body:     bytes.NewBuffer(marshaledJson),
	}
	return execute[*RawResponse](ctx, s.client, r)
}

package freelancer

import (
	"context"
	"encoding/json"
	"net/http"
)

type ListProjectsServicesService struct {
	client   *Client
	services []int
	owners   []int64
	statuses []string
}

func (s *ListProjectsServicesService) Do(ctx context.Context, id int64) (*BaseResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: string(PROJECTS_SERVICES),
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

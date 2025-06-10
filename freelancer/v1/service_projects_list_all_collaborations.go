package freelancer

import (
	"context"
	"encoding/json"
	"net/http"
)

// Returns a list of milestones on a project
type ListAllCollaborationsService struct {
	client *Client
}

func (s *ListAllCollaborationsService) Do(ctx context.Context) (*BaseResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: string(PROJECTS_COLLABORATIONS),
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

package freelancer

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// Returns a list of milestones on a project
type ListCollaborationsService struct {
	client *Client
}

func (s *ListCollaborationsService) Do(ctx context.Context, id int64) (*BaseResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("%s/%s/collaborations/", string(PROJECTS_PROJECTS), strconv.FormatInt(id, 10)),
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

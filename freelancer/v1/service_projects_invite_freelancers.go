package freelancer

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// Invite specific freelancers to bid on a project
type InviteFreelancersService struct {
	client *Client
}

type InviteFreelancersRequestBody struct {
	FreelancerID int64 `json:"freelancer_id"`
}

func (s *InviteFreelancersService) Do(ctx context.Context, projectID int64, b InviteFreelancersRequestBody) (*BaseResponse, error) {
	m, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodPost,
		endpoint: fmt.Sprintf("%s/%s/invite", PROJECTS_PROJECTS, strconv.FormatInt(projectID, 10)),
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

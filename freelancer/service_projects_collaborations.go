package freelancer

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// TODO: refine with typed response
type listCollaborations struct {
	client    *Client
	projectID int64
}

func (s *listCollaborations) Do(ctx context.Context) (*RawResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("%s/%s/collaborations/", string(ProjectsProjects), strconv.FormatInt(s.projectID, 10)),
	}
	return execute[*RawResponse](ctx, s.client, r)
}

type createCollaboration struct {
	client    *Client
	projectID int64
}
type CreateProjectCollaborationsBody struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Title    string `json:"title"`
	// required
	Permissions struct {
		Chat     bool `json:"CHAT"`
		BidAward bool `json:"BID_AWARD"`
	}
}

func (s *createCollaboration) Do(ctx context.Context, b CreateProjectCollaborationsBody) (*RawResponse, error) {
	m, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodPost,
		endpoint: fmt.Sprintf("%s/%s/collaborations/", string(ProjectsProjects), strconv.FormatInt(s.projectID, 10)),
		body:     bytes.NewBuffer(m),
	}
	return execute[*RawResponse](ctx, s.client, r)
}

type actionCollaboration struct {
	client          *Client
	projectID       int64
	collaborationID int
}
type ActionProjectCollaborationsBody struct {
	// required
	Action ProjectCollaborationAction `json:"action"`
	// required
	Permissions struct {
		Chat     bool `json:"CHAT"`
		BidAward bool `json:"BID_AWARD"`
	}
}

func (s *actionCollaboration) Do(ctx context.Context, b ActionProjectCollaborationsBody) (*RawResponse, error) {
	m, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodPut,
		endpoint: fmt.Sprintf("%s/%s/collaborations/%d", string(ProjectsProjects), strconv.FormatInt(s.projectID, 10), s.collaborationID),
		body:     bytes.NewBuffer(m),
	}
	return execute[*RawResponse](ctx, s.client, r)
}

// TODO: refine with typed response
type listAllCollaborations struct {
	client *Client
}

func (s *listAllCollaborations) Do(ctx context.Context) (*RawResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: string(ProjectsCollaborations),
	}
	return execute[*RawResponse](ctx, s.client, r)
}

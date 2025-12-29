package freelancer

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/cushydigit/go-freelancer-sdk/freelancer/endpoints"
)

// TODO: refine with typed response

// Returns a list of project collaboration data for a project.
// it maps to the `GET` `/projects/0.1/projects/{project_id}/collaborations` endpoint
func (s *CollaborationsService) List(ctx context.Context, projectID int64) (*RawResponse, error) {
	path := fmt.Sprintf("%s/%s/collaborations", endpoints.Projects, strconv.FormatInt(projectID, 10))
	return execute[*RawResponse](ctx, s.client, http.MethodGet, path, nil, nil)
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

// Creates a new project collaboration.
// It maps to the `POST` `/projects/0.1/projects/{project_id}/collaborations` endpoint
func (s *CollaborationsService) Create(ctx context.Context, projectID int64, b CreateProjectCollaborationsBody) (*RawResponse, error) {
	path := fmt.Sprintf("%s/%s/collaborations", endpoints.Projects, strconv.FormatInt(projectID, 10))
	return execute[*RawResponse](ctx, s.client, http.MethodPost, path, nil, nil)
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

// Performs an action on a collaboration.
// it maps to the `POST` `/projects/0.1/projects/{project_id}/collaborations/{collaboration_id}/actions` endpoint
func (s *CollaborationsService) Action(ctx context.Context, projectID int64, collaborationID int, b ActionProjectCollaborationsBody) (*RawResponse, error) {
	path := fmt.Sprintf("%s/%s/collaborations/%d/actions", endpoints.Projects, strconv.FormatInt(projectID, 10), collaborationID)
	return execute[*RawResponse](ctx, s.client, http.MethodPost, path, nil, b)
}

// TODO: refine with typed response

// Returns a list of all collaboration data for a user.
// It maps toi the `GET` `/projects/0.1/collaborations` endpoint
func (s *CollaborationsService) ListAll(ctx context.Context) (*RawResponse, error) {
	path := endpoints.ProjectsCollaborations
	return execute[*RawResponse](ctx, s.client, http.MethodGet, path, nil, nil)
}

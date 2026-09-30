package freelancer

import (
	"context"
	"net/http"

	"github.com/cushydigit/go-freelancer-sdk/freelancer/internal/endpoints"
	rr "github.com/cushydigit/go-freelancer-sdk/freelancer/reqres"
)

// --------------------------------------
// Resource-Collaboration
// --------------------------------------

// Returns a list of project collaboration data for a project.
// it maps to the `GET` `/projects/0.1/projects/{project_id}/collaborations` endpoint
func (r *Collaborations) List(
	ctx context.Context,
	projectID int64,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		r.client,
		http.MethodGet,
		endpoints.ProjectCollaborations(projectID),
		nil,
		nil,
	)
}

// Creates a new project collaboration.
// It maps to the `POST` `/projects/0.1/projects/{project_id}/collaborations` endpoint
func (r *Collaborations) Create(
	ctx context.Context,
	projectID int64, b rr.CreateCollaborationBody,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		r.client,
		http.MethodPost,
		endpoints.ProjectCollaborations(projectID),
		nil,
		b,
	)
}

// Performs an action on a collaboration.
// it maps to the `PUT` `/projects/0.1/projects/{project_id}/collaborations/{collaboration_id}/actions` endpoint
func (r *Collaborations) Action(
	ctx context.Context,
	projectID int64,
	collaborationID int64,
	b rr.ActionCollaborationBody,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		r.client,
		http.MethodPut,
		endpoints.ProjectCollaborationsActions(projectID, collaborationID),
		nil,
		b,
	)
}

// Returns a list of all collaboration data for a user.
// It maps toi the `GET` `/projects/0.1/projects/collaborations` endpoint
func (r *Collaborations) ListAll(
	ctx context.Context,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		r.client,
		http.MethodGet,
		endpoints.ProjectsCollaborations,
		nil,
		nil,
	)
}

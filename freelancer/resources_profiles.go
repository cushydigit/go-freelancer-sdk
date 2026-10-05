package freelancer

import (
	"context"
	"net/http"

	"github.com/cushydigit/go-freelancer-sdk/freelancer/internal/endpoints"
	rr "github.com/cushydigit/go-freelancer-sdk/freelancer/reqres"
)

// --------------------------------------
// Resources - Profiles
// --------------------------------------

// Create a new profile for a user. Returns the created profile
// It maps to the `POST` `/users/0.1/profiles` endpoint.
func (r *Profiles) Create(
	ctx context.Context,
	b rr.CreateProfileBody,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		r.client,
		http.MethodPost,
		endpoints.Profiles,
		nil,
		b,
	)
}

// NOTE: the api does not have solid on this endpoint (the get should not have body)

// Get profile(s)
// It maps to the `GET` `/users/0.1/profiles` endpoint.
func (r *Profiles) Get(
	ctx context.Context,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		r.client,
		http.MethodGet,
		endpoints.Profiles,
		nil,
		nil,
	)
}

// Update a profile
// It maps to the `PUT` `/users/0.1/profiles` endpoint.
func (r *Profiles) Update(
	ctx context.Context,
	b rr.UpdateProfileBody,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		r.client,
		http.MethodPut,
		endpoints.Profiles,
		nil,
		b,
	)
}

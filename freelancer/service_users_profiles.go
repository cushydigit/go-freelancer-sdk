package freelancer

import (
	"context"
	"net/http"

	"github.com/cushydigit/go-freelancer-sdk/freelancer/internal/endpoints"
)

type CreateProfileBody struct {
	Tagline     string `json:"tagline"`
	HourlyRate  int    `json:"hourly_Rate"`
	Description string `json:"description"`
	ProfileName string `json:"profile_name,omitempty"`
	SkillIDs    []int  `json:"skill_ids,omitempty"`
}

// Create a new profile for a user. Returns the created profile
// It maps to the `POST` `/users/0.1/profiles` endpoint.
func (s *ProfilesService) Create(ctx context.Context, b CreateProfileBody) (*RawResponse, error) {
	p := endpoints.UsersProfiles
	return execute[*RawResponse](ctx, s.client, http.MethodPost, p, nil, b)
}

// TODO: the api does not have solid on this endpoint (the get should not have body)
// TODO: refined with typed response
// Get profile(s)
// It maps to the `GET` `/users/0.1/profiles` endpoint.
func (s *ProfilesService) Get(ctx context.Context) (*RawResponse, error) {
	p := endpoints.UsersProfiles
	return execute[*RawResponse](ctx, s.client, http.MethodGet, p, nil, nil)
}

type UpdateProfileBody struct {
	ProfileID   int    `json:"profile_id"`
	Tagline     string `json:"tagline,omitempty"`
	HourlyRate  int    `json:"hourly_Rate,omitempty"`
	Description string `json:"description,omitempty"`
	ProfileName string `json:"profile_name,omitempty"`
	SkillIDs    []int  `json:"skill_ids,omitempty"`
}

// TODO: refine with typed response

// Update a profile
// It maps to the `PUT` `/users/0.1/profiles` endpoint.
func (s *ProfilesService) Update(ctx context.Context, b UpdateProfileBody) (*RawResponse, error) {
	p := endpoints.UsersProfiles
	return execute[*RawResponse](ctx, s.client, http.MethodPut, p, nil, b)
}

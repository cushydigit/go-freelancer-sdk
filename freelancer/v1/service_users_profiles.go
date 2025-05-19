package freelancer

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

// Create a new profile for a user.
type CreateProfilesService struct {
	client *Client
}

type CreateProfilesRequestBody struct {
	Tagline     string `json:"tagline"`
	HourlyRate  int    `json:"hourly_Rate"`
	Description string `json:"description"`
	ProfileName string `json:"profile_name"`
	SkillIDs    []int  `json:"skill_ids"`
}

func (s *CreateProfilesService) DO(ctx context.Context, b CreateProfilesRequestBody) (*BaseResponse, error) {
	m, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodPost,
		endpoint: string(USERS_PROFILES),
		body:     bytes.NewBuffer(m),
	}

	data, err := s.client.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	resp := BaseResponse{}
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Get profiles
type GetProfilesService struct {
	client *Client
}

type GetProfilesRequestBody struct {
	IDs    []int  `json:"ids"`
	UserID int64  `json:"user_id"`
	SeoUrl string `json:"seo_url"`
}

func (s *GetProfilesService) DO(ctx context.Context, b GetProfilesRequestBody) (*BaseResponse, error) {
	m, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodGet,
		endpoint: string(USERS_PROFILES),
		body:     bytes.NewBuffer(m),
	}
	data, err := s.client.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	resp := BaseResponse{}
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Update a profile
type UpdateProfilesService struct {
	client *Client
}

type UpdateProfilesRequestBody struct {
	ProfileID   int    `json:"profile_id"`
	Tagline     string `json:"tagline"`
	HourlyRate  int    `json:"hourly_Rate"`
	Description string `json:"description"`
	ProfileName string `json:"profile_name"`
	SkillIDs    []int  `json:"skill_ids"`
}

func (s *UpdateProfilesService) DO(ctx context.Context, b UpdateProfilesRequestBody) (*BaseResponse, error) {
	m, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodPut,
		endpoint: string(USERS_PROFILES),
		body:     bytes.NewBuffer(m),
	}
	data, err := s.client.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	resp := BaseResponse{}
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

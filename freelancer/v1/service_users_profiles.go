package freelancer

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

// Create a new profile for a user.
type CreateUsersProfilesService struct {
	client *Client
}

type CreateUsersProfilesRequestBody struct {
	Tagline     string `json:"tagline"`
	HourlyRate  int    `json:"hourly_Rate"`
	Description string `json:"description"`
	ProfileName string `json:"profile_name"`
	SkillIDs    []int  `json:"skill_ids"`
}

func (s *CreateUsersProfilesService) DO(ctx context.Context, b CreateUsersProfilesRequestBody) (*BaseResponse, error) {
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
type GetUsersProfilesService struct {
	client *Client
}

type GetUsersProfilesRequestBody struct {
	IDs    []int  `json:"ids"`
	UserID int64  `json:"user_id"`
	SeoUrl string `json:"seo_url"`
}

func (s *GetUsersProfilesService) DO(ctx context.Context, b GetUsersProfilesRequestBody) (*BaseResponse, error) {
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
type UpdateUsersProfilesService struct {
	client *Client
}

type UpdateUsersProfilesRequestBody struct {
	ProfileID   int    `json:"profile_id"`
	Tagline     string `json:"tagline"`
	HourlyRate  int    `json:"hourly_Rate"`
	Description string `json:"description"`
	ProfileName string `json:"profile_name"`
	SkillIDs    []int  `json:"skill_ids"`
}

func (s *UpdateUsersProfilesService) DO(ctx context.Context, b UpdateUsersProfilesRequestBody) (*BaseResponse, error) {
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

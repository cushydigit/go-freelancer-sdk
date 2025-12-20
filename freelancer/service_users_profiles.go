package freelancer

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

type createProfile struct {
	client *Client
}

type CreateProfileBody struct {
	Tagline     string `json:"tagline"`
	HourlyRate  int    `json:"hourly_Rate"`
	Description string `json:"description"`
	ProfileName string `json:"profile_name,omitempty"`
	SkillIDs    []int  `json:"skill_ids,omitempty"`
}

func (s *createProfile) Do(ctx context.Context, b CreateProfileBody) (*BaseResponse, error) {
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

type getProfile struct {
	client *Client
}

// TODO: the api does not have solid on this endpoint (the get should not have body)
// it should tested
// type GetProfilesRequestBody struct {
// 	IDs    []int  `json:"ids"`
// 	UserID int64  `json:"user_id"`
// 	SeoUrl string `json:"seo_url"`
// }

func (s *getProfile) Do(ctx context.Context) (*BaseResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: string(USERS_PROFILES),
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

type updateProfile struct {
	client *Client
}

type UpdateProfileBody struct {
	ProfileID   int    `json:"profile_id"`
	Tagline     string `json:"tagline,omitempty"`
	HourlyRate  int    `json:"hourly_Rate,omitempty"`
	Description string `json:"description,omitempty"`
	ProfileName string `json:"profile_name,omitempty"`
	SkillIDs    []int  `json:"skill_ids,omitempty"`
}

func (s *updateProfile) Do(ctx context.Context, b UpdateProfileBody) (*BaseResponse, error) {
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

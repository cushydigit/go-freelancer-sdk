package freelancer

import (
	"context"
	"encoding/json"
	"net/http"
)

type profilesGetService struct {
	client *Client
}

// TODO: the api does not have solid on this endpoint (the get should not have body)
// it should tested
// type GetProfilesRequestBody struct {
// 	IDs    []int  `json:"ids"`
// 	UserID int64  `json:"user_id"`
// 	SeoUrl string `json:"seo_url"`
// }

func (s *profilesGetService) Do(ctx context.Context) (*BaseResponse, error) {
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

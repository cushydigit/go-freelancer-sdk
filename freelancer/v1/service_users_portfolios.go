package freelancer

import (
	"context"
	"encoding/json"
	"net/http"
)

type ListUsersPortfoliosService struct {
	client *Client
	users  []int64
	limit  *int
	offset *int
}

type ListUsersPortfoliosResponse struct {
	Status    string          `json:"status"`
	RequestID string          `json:"request_id"`
	Result    json.RawMessage `json:"result"`
}

func (s *ListUsersPortfoliosService) DO(ctx context.Context) (*ListUsersPortfoliosResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: string(USERS_PORTFOLIOS),
	}

	for id := range s.users {
		r.addParam("users[]", id)
	}

	if s.limit != nil {
		r.addParam("limit", s.limit)
	}

	if s.offset != nil {
		r.addParam("offset", s.offset)
	}

	data, err := s.client.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}

	resp := &ListUsersPortfoliosResponse{}
	if err := json.Unmarshal(data, resp); err != nil {
		return nil, err
	}

	return resp, nil

}

func (s *ListUsersPortfoliosService) SetUsers(users []int64) *ListUsersPortfoliosService {
	s.users = users
	return s
}

func (s *ListUsersPortfoliosService) SetLimit(limit int) *ListUsersPortfoliosService {
	s.limit = &limit
	return s
}

func (s *ListUsersPortfoliosService) SetOffset(offset int) *ListUsersPortfoliosService {
	s.offset = &offset
	return s
}

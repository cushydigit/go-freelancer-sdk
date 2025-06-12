package freelancer

import (
	"context"
	"encoding/json"
	"net/http"
)

type portfoliosListService struct {
	client *Client
	users  []int64
	limit  *int
	offset *int
}

func (s *portfoliosListService) DO(ctx context.Context) (*BaseResponse, error) {
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

	resp := &BaseResponse{}
	if err := json.Unmarshal(data, resp); err != nil {
		return nil, err
	}

	return resp, nil

}

func (s *portfoliosListService) SetUsers(users []int64) *portfoliosListService {
	s.users = users
	return s
}

func (s *portfoliosListService) SetLimit(limit int) *portfoliosListService {
	s.limit = &limit
	return s
}

func (s *portfoliosListService) SetOffset(offset int) *portfoliosListService {
	s.offset = &offset
	return s
}

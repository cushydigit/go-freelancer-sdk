package freelancer

import (
	"context"
	"encoding/json"
	"net/http"
)

type ListUsersPoolsService struct {
	client          *Client
	pools           []int
	names           []string
	seoUrls         []string
	ignoreTest      *bool
	isTalentNetwork *bool
	limit           *int
	offset          *int
}

func (s *ListUsersPoolsService) DO(ctx context.Context) (*BaseResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: string(USERS_ENTERPRISES),
	}

	for _, val := range s.pools {
		r.addParam("pools[]", val)
	}

	for _, val := range s.names {
		r.addParam("internal_names[]", val)
	}

	for _, val := range s.seoUrls {
		r.addParam("seo_urls[]", val)
	}

	if s.ignoreTest != nil {
		r.addParam("ignore_test", s.ignoreTest)
	}

	if s.isTalentNetwork != nil {
		r.addParam("is_talent_network", s.isTalentNetwork)
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

	return resp, err
}

func (s *ListUsersPoolsService) SetPools(val []int) *ListUsersPoolsService {
	s.pools = val
	return s
}

func (s *ListUsersPoolsService) SetNames(val []string) *ListUsersPoolsService {
	s.names = val
	return s
}

func (s *ListUsersPoolsService) SetSeoUrls(val []string) *ListUsersPoolsService {
	s.seoUrls = val
	return s
}

func (s *ListUsersPoolsService) SetIgnoreTest(val bool) *ListUsersPoolsService {
	s.ignoreTest = &val
	return s
}

func (s *ListUsersPoolsService) SetIsTalentNetwork(val bool) *ListUsersPoolsService {
	s.isTalentNetwork = &val
	return s
}

func (s *ListUsersPoolsService) SetLimit(val int) *ListUsersPoolsService {
	s.limit = &val
	return s
}

func (s *ListUsersPoolsService) SetOffset(val int) *ListUsersPoolsService {
	s.offset = &val
	return s
}

package freelancer

import (
	"context"
	"encoding/json"
	"net/http"
)

// Returns a list of pools belonging to the current user
type ListPoolsService struct {
	client          *Client
	pools           []int
	names           []string
	seoUrls         []string
	ignoreTest      *bool
	isTalentNetwork *bool
	limit           *int
	offset          *int
}

func (s *ListPoolsService) DO(ctx context.Context) (*BaseResponse, error) {
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
		r.addParam("ignore_test", *s.ignoreTest)
	}
	if s.isTalentNetwork != nil {
		r.addParam("is_talent_network", *s.isTalentNetwork)
	}
	if s.limit != nil {
		r.addParam("limit", *s.limit)
	}
	if s.offset != nil {
		r.addParam("offset", *s.offset)
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

func (s *ListPoolsService) SetPools(val []int) *ListPoolsService {
	s.pools = val
	return s
}

func (s *ListPoolsService) SetNames(val []string) *ListPoolsService {
	s.names = val
	return s
}

func (s *ListPoolsService) SetSeoUrls(val []string) *ListPoolsService {
	s.seoUrls = val
	return s
}

func (s *ListPoolsService) SetIgnoreTest(val bool) *ListPoolsService {
	s.ignoreTest = &val
	return s
}

func (s *ListPoolsService) SetIsTalentNetwork(val bool) *ListPoolsService {
	s.isTalentNetwork = &val
	return s
}

func (s *ListPoolsService) SetLimit(val int) *ListPoolsService {
	s.limit = &val
	return s
}

func (s *ListPoolsService) SetOffset(val int) *ListPoolsService {
	s.offset = &val
	return s
}

package freelancer

import (
	"context"
	"encoding/json"
	"net/http"
)

// Returns a list of pools belonging to the current user
type poolsListService struct {
	client          *Client
	pools           []int
	names           []string
	seoUrls         []string
	ignoreTest      *bool
	isTalentNetwork *bool
	limit           *int
	offset          *int
}

func (s *poolsListService) Do(ctx context.Context) (*BaseResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: string(USERS_POOLS),
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

// SetPools sets the list of pool IDs to filter results by.
func (s *poolsListService) SetPools(vals []int) *poolsListService {
	s.pools = vals
	return s
}

// SetNames sets the list of pool names to filter results by.
func (s *poolsListService) SetNames(vals []string) *poolsListService {
	s.names = vals
	return s
}

// SetSeoUrls sets the list of SEO URLs to filter results by.
func (s *poolsListService) SetSeoUrls(vals []string) *poolsListService {
	s.seoUrls = vals
	return s
}

// SetIgnoreTest sets whether to ignore test pools.
func (s *poolsListService) SetIgnoreTest(val bool) *poolsListService {
	s.ignoreTest = &val
	return s
}

// SetIsTalentNetwork sets whether to filter pools based on the talent network flag.
func (s *poolsListService) SetIsTalentNetwork(val bool) *poolsListService {
	s.isTalentNetwork = &val
	return s
}

// SetLimit sets the maximum number of results to return. Default is 100.
func (s *poolsListService) SetLimit(val int) *poolsListService {
	s.limit = &val
	return s
}

// SetOffset sets the number of results to skip for pagination. Default is 0.
func (s *poolsListService) SetOffset(val int) *poolsListService {
	s.offset = &val
	return s
}

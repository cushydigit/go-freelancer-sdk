package freelancer

import (
	"context"
	"encoding/json"
	"net/http"
)

type enterprisesListService struct {
	client        *Client
	enterprises   []int
	internalNames []string
	names         []string
	seoUrls       []string
	userID        *int64
	ignoreTest    *bool
	limit         *int
	offset        *int
}

func (s *enterprisesListService) Do(ctx context.Context) (*BaseResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: string(USERS_ENTERPRISES),
	}

	for _, val := range s.enterprises {
		r.addParam("enterprises[]", val)
	}

	for _, val := range s.internalNames {
		r.addParam("internal_names[]", val)
	}

	for _, val := range s.names {
		r.addParam("names[]", val)
	}

	for _, val := range s.seoUrls {
		r.addParam("seo_urls[]", val)
	}

	if s.userID != nil {
		r.addParam("user_id", *s.userID)
	}

	if s.ignoreTest != nil {
		r.addParam("ignore_test", *s.ignoreTest)
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

// SetEnterprises sets the filter to return enterprises with the specified enterprise IDs.
func (s *enterprisesListService) SetEnterprises(vals []int) *enterprisesListService {
	s.enterprises = vals
	return s
}

// SetInternalNames sets the filter to return enterprises with the specified internal names.
func (s *enterprisesListService) SetInternalNames(vals []string) *enterprisesListService {
	s.internalNames = vals
	return s
}

// SetNames sets the filter to return enterprises with the specified names.
func (s *enterprisesListService) SetNames(vals []string) *enterprisesListService {
	s.names = vals
	return s
}

// SetSeoUrls sets the filter to return enterprises with the specified SEO URLs.
func (s *enterprisesListService) SetSeoUrls(vals []string) *enterprisesListService {
	s.seoUrls = vals
	return s
}

// SetUserID sets the filter to return enterprises for the specified user ID.
func (s *enterprisesListService) SetUserID(val int64) *enterprisesListService {
	s.userID = &val
	return s
}

// SetIgnoreTest sets the filter to ignore test enterprises.
func (s *enterprisesListService) SetIgnoreTest(val bool) *enterprisesListService {
	s.ignoreTest = &val
	return s
}

// SetLimit sets the maximum number of results to return. Default is 100.
func (s *enterprisesListService) SetLimit(val int) *enterprisesListService {
	s.limit = &val
	return s
}

// SetOffset sets the number of results to skip, for pagination. Default is 0.
func (s *enterprisesListService) SetOffset(val int) *enterprisesListService {
	s.offset = &val
	return s
}

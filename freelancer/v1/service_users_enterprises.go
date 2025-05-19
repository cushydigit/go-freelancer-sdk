package freelancer

import (
	"context"
	"encoding/json"
	"net/http"
)

// Returns a list of enterprises
type ListEnterprisesService struct {
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

func (s *ListEnterprisesService) DO(ctx context.Context) (*BaseResponse, error) {
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

func (s *ListEnterprisesService) SetEnterprises(val []int) *ListEnterprisesService {
	s.enterprises = val
	return s
}

func (s *ListEnterprisesService) SetInternalNames(val []string) *ListEnterprisesService {
	s.internalNames = val
	return s
}

func (s *ListEnterprisesService) SetNames(val []string) *ListEnterprisesService {
	s.names = val
	return s
}

func (s *ListEnterprisesService) SetSeoUrls(val []string) *ListEnterprisesService {
	s.seoUrls = val
	return s
}

func (s *ListEnterprisesService) SetUserID(val int64) *ListEnterprisesService {
	s.userID = &val
	return s
}

func (s *ListEnterprisesService) SetIgnoreTest(val bool) *ListEnterprisesService {
	s.ignoreTest = &val
	return s
}

func (s *ListEnterprisesService) SetLimit(val int) *ListEnterprisesService {
	s.limit = &val
	return s
}

func (s *ListEnterprisesService) SetOffset(val int) *ListEnterprisesService {
	s.offset = &val
	return s
}

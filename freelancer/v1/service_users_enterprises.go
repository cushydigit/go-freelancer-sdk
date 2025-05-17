package freelancer

import (
	"context"
	"encoding/json"
	"net/http"
)

type ListUsersEnterprisesService struct {
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

func (s *ListUsersEnterprisesService) DO(ctx context.Context) (*BaseResponse, error) {
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
		r.addParam("user_id", s.userID)
	}

	if s.ignoreTest != nil {
		r.addParam("ignore_test", s.ignoreTest)
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

func (s *ListUsersEnterprisesService) SetEnterprises(val []int) *ListUsersEnterprisesService {
	s.enterprises = val
	return s
}

func (s *ListUsersEnterprisesService) SetInternalNames(val []string) *ListUsersEnterprisesService {
	s.internalNames = val
	return s
}

func (s *ListUsersEnterprisesService) SetNames(val []string) *ListUsersEnterprisesService {
	s.names = val
	return s
}

func (s *ListUsersEnterprisesService) SetSeoUrls(val []string) *ListUsersEnterprisesService {
	s.seoUrls = val
	return s
}

func (s *ListUsersEnterprisesService) SetUserID(val int64) *ListUsersEnterprisesService {
	s.userID = &val
	return s
}

func (s *ListUsersEnterprisesService) SetIgnoreTest(val bool) *ListUsersEnterprisesService {
	s.ignoreTest = &val
	return s
}

func (s *ListUsersEnterprisesService) SetLimit(val int) *ListUsersEnterprisesService {
	s.limit = &val
	return s
}

func (s *ListUsersEnterprisesService) SetOffset(val int) *ListUsersEnterprisesService {
	s.offset = &val
	return s
}

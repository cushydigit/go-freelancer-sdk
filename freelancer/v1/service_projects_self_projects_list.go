package freelancer

import (
	"context"
	"encoding/json"
	"net/http"
)

type selfProjectsListService struct {
	client      *Client
	status      *ProjectStatusType
	role        *RoleType
	types       []TypeType
	query       *string
	sortField   *SortField
	reverseSort *bool
	recruiter   *bool
	offset      *int
	limit       *int
}

func (s *selfProjectsListService) Do(ctx context.Context) (*BaseResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: string(PROJECTS_SELF),
	}
	if s.status != nil {
		r.addParam("status", *s.status)
	}
	if s.role != nil {
		r.addParam("role", *s.role)
	}
	for _, val := range s.types {
		r.addParam("types[]", val)
	}
	if s.query != nil {
		r.addParam("query", *s.query)
	}
	if s.sortField != nil {
		r.addParam("sort_field", *s.sortField)
	}
	if s.reverseSort != nil {
		r.addParam("reverse_sort", *s.reverseSort)
	}
	if s.recruiter != nil {
		r.addParam("recruiter", *s.recruiter)
	}
	if s.offset != nil {
		r.addParam("offset", *s.offset)
	}
	if s.limit != nil {
		r.addParam("limit", *s.limit)
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

func (s *selfProjectsListService) SetStatus(val ProjectStatusType) *selfProjectsListService {
	s.status = &val
	return s
}

func (s *selfProjectsListService) SetRole(val RoleType) *selfProjectsListService {
	s.role = &val
	return s
}

func (s *selfProjectsListService) SetTypes(vals []TypeType) *selfProjectsListService {
	s.types = vals
	return s
}

func (s *selfProjectsListService) SetQuery(val string) *selfProjectsListService {
	s.query = &val
	return s
}

func (s *selfProjectsListService) SetSortField(val SortField) *selfProjectsListService {
	s.sortField = &val
	return s
}

func (s *selfProjectsListService) SetReverseSort(val bool) *selfProjectsListService {
	s.reverseSort = &val
	return s
}

func (s *selfProjectsListService) SetRecruiter(val bool) *selfProjectsListService {
	s.recruiter = &val
	return s
}

func (s *selfProjectsListService) SetOffset(val int) *selfProjectsListService {
	s.offset = &val
	return s
}

func (s *selfProjectsListService) SetLimit(val int) *selfProjectsListService {
	s.limit = &val
	return s
}

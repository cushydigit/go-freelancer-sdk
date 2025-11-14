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

// SetStatus sets the status filter for the projects/contests.
// Example: awarded. Valid values: active, all, awarded, open, past.
func (s *selfProjectsListService) SetStatus(val ProjectStatusType) *selfProjectsListService {
	s.status = &val
	return s
}

// SetRole sets the role of the user in the projects/contests.
// Example: employer. Valid values: freelancer, employer.
func (s *selfProjectsListService) SetRole(val RoleType) *selfProjectsListService {
	s.role = &val
	return s
}

// SetTypes sets the type filter for the results, e.g., projects, contests, or both.
// Default: projects, contests. Example: projects.
func (s *selfProjectsListService) SetTypes(vals []TypeType) *selfProjectsListService {
	s.types = vals
	return s
}

// SetQuery sets the search terms to filter projects/contests by title, description, or usernames.
// Example: build website.
func (s *selfProjectsListService) SetQuery(val string) *selfProjectsListService {
	s.query = &val
	return s
}

// SetSortField sets the field by which to sort the results.
// Default: project_id. Valid values: project_id, submitdate.
func (s *selfProjectsListService) SetSortField(val SortField) *selfProjectsListService {
	s.sortField = &val
	return s
}

// SetReverseSort sets whether the sorting should be in ascending order.
// Example: true. If true, results are sorted in ascending order.
func (s *selfProjectsListService) SetReverseSort(val bool) *selfProjectsListService {
	s.reverseSort = &val
	return s
}

// SetRecruiter sets the filter to return only recruiter projects.
// Example: true. If true, only recruiter projects are returned.
func (s *selfProjectsListService) SetRecruiter(val bool) *selfProjectsListService {
	s.recruiter = &val
	return s
}

// SetOffset sets the number of results to skip (pagination).
// Default: 0. Example: 30.
func (s *selfProjectsListService) SetOffset(val int) *selfProjectsListService {
	s.offset = &val
	return s
}

// SetLimit sets the maximum number of results to return.
// Default: 100. Example: 10.
func (s *selfProjectsListService) SetLimit(val int) *selfProjectsListService {
	s.limit = &val
	return s
}

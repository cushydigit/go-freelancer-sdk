package freelancer

import (
	"context"
	"encoding/json"
	"net/http"
)

type ListProjectsActiveServicesService struct {
	client       *Client
	query        *string
	sort         *SortType
	reverseSort  *bool
	extraDetails *bool
	fileDetails  *bool
	jobDetails   *bool
	userDetails  *bool
	offset       *int
	compact      *bool
}

func (s *ListProjectsActiveServicesService) Do(ctx context.Context, id int64) (*BaseResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: string(PROJECTS_SERVICES),
	}

	if s.query != nil {
		r.addParam("query", *s.query)
	}
	if s.sort != nil {
		r.addParam("sort", *s.sort)
	}
	if s.reverseSort != nil {
		r.addParam("reverse_sort", *s.reverseSort)
	}
	if s.extraDetails != nil {
		r.addParam("extra_details", *s.extraDetails)
	}
	if s.fileDetails != nil {
		r.addParam("file_details", *s.fileDetails)
	}
	if s.jobDetails != nil {
		r.addParam("job_details", *s.jobDetails)
	}
	if s.userDetails != nil {
		r.addParam("user_details", *s.userDetails)
	}
	if s.offset != nil {
		r.addParam("offset", *s.offset)
	}
	if s.compact != nil {
		r.addParam("compact", *s.compact)
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

func (s *ListProjectsActiveServicesService) SetQuery(val string) *ListProjectsActiveServicesService {
	s.query = &val
	return s
}

func (s *ListProjectsActiveServicesService) SetSort(val SortType) *ListProjectsActiveServicesService {
	s.sort = &val
	return s
}

func (s *ListProjectsActiveServicesService) SetReverseSort(val bool) *ListProjectsActiveServicesService {
	s.reverseSort = &val
	return s
}

func (s *ListProjectsActiveServicesService) SetExtraDetails(val bool) *ListProjectsActiveServicesService {
	s.extraDetails = &val
	return s
}

func (s *ListProjectsActiveServicesService) SetFileDetails(val bool) *ListProjectsActiveServicesService {
	s.fileDetails = &val
	return s
}

func (s *ListProjectsActiveServicesService) SetJobDetails(val bool) *ListProjectsActiveServicesService {
	s.jobDetails = &val
	return s
}

func (s *ListProjectsActiveServicesService) SetUserDetails(val bool) *ListProjectsActiveServicesService {
	s.userDetails = &val
	return s
}

func (s *ListProjectsActiveServicesService) SetOffset(offset int) *ListProjectsActiveServicesService {
	s.offset = &offset
	return s
}

func (s *ListProjectsActiveServicesService) SetCompact(val bool) *ListProjectsActiveServicesService {
	s.compact = &val
	return s
}

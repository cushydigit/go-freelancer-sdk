package freelancer

import (
	"context"
	"encoding/json"
	"net/http"
)

type jobsListService struct {
	client                    *Client
	jobs                      []int
	jobNames                  []string
	seoUrls                   []string
	categories                []int
	onlyLocal                 *bool
	activeProjectCountDetails *bool
	seoDetails                *bool
	seoCountryName            *string
	lang                      *string
}

func (s *jobsListService) Do(ctx context.Context) (*BaseResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: string(PROJECTS_JOBS),
	}

	for _, val := range s.jobs {
		r.addParam("jobs[]", val)
	}
	for _, val := range s.jobNames {
		r.addParam("job_names[]", val)
	}
	for _, val := range s.seoUrls {
		r.addParam("seo_urls[]", val)
	}
	for _, val := range s.categories {
		r.addParam("categories[]", val)
	}
	if s.onlyLocal != nil {
		r.addParam("only_local", *s.onlyLocal)
	}
	if s.activeProjectCountDetails != nil {
		r.addParam("active_project_count_details", *s.activeProjectCountDetails)
	}
	if s.seoDetails != nil {
		r.addParam("seo_details", *s.seoDetails)
	}
	if s.seoCountryName != nil {
		r.addParam("seo_country_name", *s.seoCountryName)
	}
	if s.lang != nil {
		r.addParam("lang", *s.lang)
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
func (s *jobsListService) SetJobs(vals []int) *jobsListService {
	s.jobs = vals
	return s
}
func (s *jobsListService) SetJobNames(vals []string) *jobsListService {
	s.jobNames = vals
	return s
}

func (s *jobsListService) SetSeoUrls(vals []string) *jobsListService {
	s.seoUrls = vals
	return s
}

func (s *jobsListService) SetCategories(vals []int) *jobsListService {
	s.categories = vals
	return s
}

func (s *jobsListService) SetOnlyLocal(val bool) *jobsListService {
	s.onlyLocal = &val
	return s
}

func (s *jobsListService) SetActiveProjectCountDetails(val bool) *jobsListService {
	s.activeProjectCountDetails = &val
	return s
}

func (s *jobsListService) SetSeoDetails(val bool) *jobsListService {
	s.seoDetails = &val
	return s
}

func (s *jobsListService) SetSeoCountryName(val string) *jobsListService {
	s.seoCountryName = &val
	return s
}

func (s *jobsListService) SetLang(val string) *jobsListService {
	s.lang = &val
	return s
}

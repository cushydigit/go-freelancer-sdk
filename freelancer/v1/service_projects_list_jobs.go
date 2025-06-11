package freelancer

import (
	"context"
	"encoding/json"
	"net/http"
)

type listJobsService struct {
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

func (s *listJobsService) Do(ctx context.Context) (*BaseResponse, error) {
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
func (s *listJobsService) SetJobs(vals []int) *listJobsService {
	s.jobs = vals
	return s
}
func (s *listJobsService) SetJobNames(vals []string) *listJobsService {
	s.jobNames = vals
	return s
}

func (s *listJobsService) SetSeoUrls(vals []string) *listJobsService {
	s.seoUrls = vals
	return s
}

func (s *listJobsService) SetCategories(vals []int) *listJobsService {
	s.categories = vals
	return s
}

func (s *listJobsService) SetOnlyLocal(val bool) *listJobsService {
	s.onlyLocal = &val
	return s
}

func (s *listJobsService) SetActiveProjectCountDetails(val bool) *listJobsService {
	s.activeProjectCountDetails = &val
	return s
}

func (s *listJobsService) SetSeoDetails(val bool) *listJobsService {
	s.seoDetails = &val
	return s
}

func (s *listJobsService) SetSeoCountryName(val string) *listJobsService {
	s.seoCountryName = &val
	return s
}

func (s *listJobsService) SetLang(val string) *listJobsService {
	s.lang = &val
	return s
}

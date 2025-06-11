package freelancer

import (
	"context"
	"encoding/json"
	"net/http"
)

type searchJobsService struct {
	client                    *Client
	jobs                      []int
	jobNames                  []string
	seoUrls                   []string
	seoTexts                  []string
	categories                []int
	onlyLocal                 *bool
	activeProjectCountDetails *bool
	seoDetails                *bool
	seoCountryName            *string
	lang                      *string
}

func (s *searchJobsService) Do(ctx context.Context) (*BaseResponse, error) {
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
	for _, val := range s.seoTexts {
		r.addParam("seo_texts[]", val)
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
func (s *searchJobsService) SetJobs(vals []int) *searchJobsService {
	s.jobs = vals
	return s
}
func (s *searchJobsService) SetJobNames(vals []string) *searchJobsService {
	s.jobNames = vals
	return s
}

func (s *searchJobsService) SetSeoUrls(vals []string) *searchJobsService {
	s.seoUrls = vals
	return s
}

func (s *searchJobsService) SetSeoTexts(vals []string) *searchJobsService {
	s.seoTexts = vals
	return s
}

func (s *searchJobsService) SetCategories(vals []int) *searchJobsService {
	s.categories = vals
	return s
}

func (s *searchJobsService) SetOnlyLocal(val bool) *searchJobsService {
	s.onlyLocal = &val
	return s
}

func (s *searchJobsService) SetActiveProjectCountDetails(val bool) *searchJobsService {
	s.activeProjectCountDetails = &val
	return s
}

func (s *searchJobsService) SetSeoDetails(val bool) *searchJobsService {
	s.seoDetails = &val
	return s
}

func (s *searchJobsService) SetSeoCountryName(val string) *searchJobsService {
	s.seoCountryName = &val
	return s
}

func (s *searchJobsService) SetLang(val string) *searchJobsService {
	s.lang = &val
	return s
}

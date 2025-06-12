package freelancer

import (
	"context"
	"encoding/json"
	"net/http"
)

type JobsSearchService struct {
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

func (s *JobsSearchService) Do(ctx context.Context) (*BaseResponse, error) {
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
func (s *JobsSearchService) SetJobs(vals []int) *JobsSearchService {
	s.jobs = vals
	return s
}
func (s *JobsSearchService) SetJobNames(vals []string) *JobsSearchService {
	s.jobNames = vals
	return s
}

func (s *JobsSearchService) SetSeoUrls(vals []string) *JobsSearchService {
	s.seoUrls = vals
	return s
}

func (s *JobsSearchService) SetSeoTexts(vals []string) *JobsSearchService {
	s.seoTexts = vals
	return s
}

func (s *JobsSearchService) SetCategories(vals []int) *JobsSearchService {
	s.categories = vals
	return s
}

func (s *JobsSearchService) SetOnlyLocal(val bool) *JobsSearchService {
	s.onlyLocal = &val
	return s
}

func (s *JobsSearchService) SetActiveProjectCountDetails(val bool) *JobsSearchService {
	s.activeProjectCountDetails = &val
	return s
}

func (s *JobsSearchService) SetSeoDetails(val bool) *JobsSearchService {
	s.seoDetails = &val
	return s
}

func (s *JobsSearchService) SetSeoCountryName(val string) *JobsSearchService {
	s.seoCountryName = &val
	return s
}

func (s *JobsSearchService) SetLang(val string) *JobsSearchService {
	s.lang = &val
	return s
}

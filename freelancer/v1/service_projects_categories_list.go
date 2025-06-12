package freelancer

import (
	"context"
	"encoding/json"
	"net/http"
)

type categoriesListService struct {
	client                    *Client
	categories                []int
	jobDetails                bool
	lang                      string
	activeProjectCountDetails bool
	seoDetails                bool
}

func (s *categoriesListService) Do(ctx context.Context) (*BaseResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: string(PROJECTS_CATEGORIES),
	}
	if s.categories != nil {
		r.addParam("categories", s.categories)
	}
	if s.jobDetails {
		r.addParam("job_details", s.jobDetails)
	}
	if s.lang != "" {
		r.addParam("lang", s.lang)
	}
	if s.activeProjectCountDetails {
		r.addParam("active_project_count_details", s.activeProjectCountDetails)
	}
	if s.seoDetails {
		r.addParam("seo_details", s.seoDetails)
	}
	data, err := s.client.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := &BaseResponse{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *categoriesListService) SetCategories(categories []int) *categoriesListService {
	s.categories = categories
	return s
}

func (s *categoriesListService) SetJobDetails(jobDetails bool) *categoriesListService {
	s.jobDetails = jobDetails
	return s
}

func (s *categoriesListService) SetLang(lang string) *categoriesListService {
	s.lang = lang
	return s
}

func (s *categoriesListService) SetActiveProjectCountDetails(activeProjectCountDetails bool) *categoriesListService {
	s.activeProjectCountDetails = activeProjectCountDetails
	return s
}

func (s *categoriesListService) SetSeoDetails(seoDetails bool) *categoriesListService {
	s.seoDetails = seoDetails
	return s
}

package freelancer

import (
	"context"
	"encoding/json"
	"net/http"
)

type ListCategoriesService struct {
	client                    *Client
	categories                []int
	jobDetails                bool
	lang                      string
	activeProjectCountDetails bool
	seoDetails                bool
}

type ListCategoriesResponse struct {
	Status    string           `json:"status"`
	RequestID string           `json:"request_id,omitempty"` // Optional
	Result    CategoriesResult `json:"result"`
}

type CategoriesResult struct {
	Jobs       *Jobs      `json:"jobs,omitempty"` // Jobs can be an object or null, represented by a pointer
	Categories []Category `json:"categories"`
}

// Do perform GET request on endpoint "projects/0.1/categories/"
func (s *ListCategoriesService) Do(ctx context.Context) (*ListCategoriesResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/projects/0.1/categories",
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
	res := &ListCategoriesResponse{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *ListCategoriesService) SetCategories(categories []int) *ListCategoriesService {
	s.categories = categories
	return s
}

func (s *ListCategoriesService) SetJobDetails(jobDetails bool) *ListCategoriesService {
	s.jobDetails = jobDetails
	return s
}

func (s *ListCategoriesService) SetLang(lang string) *ListCategoriesService {
	s.lang = lang
	return s
}

func (s *ListCategoriesService) SetActiveProjectCountDetails(activeProjectCountDetails bool) *ListCategoriesService {
	s.activeProjectCountDetails = activeProjectCountDetails
	return s
}

func (s *ListCategoriesService) SetSeoDetails(seoDetails bool) *ListCategoriesService {
	s.seoDetails = seoDetails
	return s
}

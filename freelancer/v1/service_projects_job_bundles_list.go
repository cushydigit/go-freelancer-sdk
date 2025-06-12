package freelancer

import (
	"context"
	"encoding/json"
	"net/http"
)

type jobBundlesListService struct {
	client     *Client
	jobBundles []int
	categories []int
	lang       *string
}

func (s *jobBundlesListService) Do(ctx context.Context) (*BaseResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: string(PROJECTS_JOB_BUNDLES),
	}

	for _, val := range s.jobBundles {
		r.addParam("job_bundles[]", val)
	}
	for _, val := range s.categories {
		r.addParam("categories[]", val)
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

func (s *jobBundlesListService) SetJobBundles(vals []int) *jobBundlesListService {
	s.jobBundles = vals
	return s
}

func (s *jobBundlesListService) SetCategories(vals []int) *jobBundlesListService {
	s.categories = vals
	return s
}

func (s *jobBundlesListService) SetLang(val string) *jobBundlesListService {
	s.lang = &val
	return s
}

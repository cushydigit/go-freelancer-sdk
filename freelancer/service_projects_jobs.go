package freelancer

import (
	"context"
	"net/http"

	"github.com/cushydigit/go-freelancer-sdk/freelancer/internal/endpoints"
	"github.com/cushydigit/go-freelancer-sdk/freelancer/internal/query"
)

type ListJobsOptions struct {
	Jobs                      []int64  `url:"jobs[]"`
	JobNames                  []string `url:"job_names[]"`
	SeoUrls                   []string `url:"seo_urls[]"`
	Categories                []int64  `url:"categories[]"`
	OnlyLocal                 *bool    `url:"only_local"`
	ActiveProjectCountDetails *bool    `url:"active_project_count_details"`
	SeoDetails                *bool    `url:"seo_details"`
	SeoCountryName            *string  `url:"seo_country_name"`
	Lang                      *string  `url:"lang"`
}

// TODO: refine with typed response

// Returns a list of milestone requests.
// It maps to the `GET` `/projects/0.1/jobs` endpoint
func (s *JobsService) List(ctx context.Context, opts *ListJobsOptions) (*RawResponse, error) {
	p := endpoints.ProjectsJobs
	q := query.Values(opts)
	return execute[*RawResponse](ctx, s.client, http.MethodGet, p, q, nil)
}

type SearchJobsOptions struct {
	Jobs                      []int64  `url:"jobs[]"`
	JobNames                  []string `url:"job_names[]"`
	SeoUrls                   []string `url:"seo_urls[]"`
	SeoTexts                  []string `url:"seo_texts[]"`
	Categories                []int64  `url:"categories[]"`
	OnlyLocal                 *bool    `url:"only_local"`
	ActiveProjectCountDetails *bool    `url:"active_project_count_details"`
	SeoDetails                *bool    `url:"seo_details"`
	SeoCountryName            *string  `url:"seo_country_name"`
	Lang                      *string  `url:"lang"`
}

// TODO: refine with typed response

// Returns a list of jobs. Note: This performs a sub-string search for all the parameters specified on the jobs.
// It maps to the `GET` `/projects/0.1/jobs/search` endpoint
func (s *JobsService) Search(ctx context.Context, opts *SearchJobsOptions) (*RawResponse, error) {
	p := endpoints.ProjectsJobsSearch
	q := query.Values(opts)
	return execute[*RawResponse](ctx, s.client, http.MethodGet, p, q, nil)
}

type ListJobBundlesOptions struct {
	JobBundles []int64 `url:"job_bundles[]"`
	Categories []int64 `url:"categories[]"`
	Lang       *string `url:"lang"`
}

// TODO: refine with typed response

// Returns a list of job bundles. Note: Categories in this context are job bundle categories. These are not the same as job categories even though they share the same name.
// It maps to the `GET` `/projects/0.1/job_bundles` endpoint
func (s *JobsService) ListJobBundles(ctx context.Context, opts *ListJobBundlesOptions) (*RawResponse, error) {
	p := endpoints.ProjectsJobBundles
	q := query.Values(opts)
	return execute[*RawResponse](ctx, s.client, http.MethodGet, p, q, nil)
}

type ListJobBundleCategoriesOptions struct {
	JobBundles []int64 `url:"job_bundles[]"`
	Categories []int64 `url:"categories[]"`
	Lang       *string `url:"lang"`
}

// TODO: refine with typed response

// Returns a list of job bundle categories.
// It maps to the `GET` `/projects/0.1/job_bundle_categories` endpoint
func (s *JobsService) ListJobBundleCategories(ctx context.Context, opts *ListJobBundleCategoriesOptions) (*RawResponse, error) {
	p := endpoints.ProjectsJobBundleCategories
	q := query.Values(opts)
	return execute[*RawResponse](ctx, s.client, http.MethodGet, p, q, nil)
}

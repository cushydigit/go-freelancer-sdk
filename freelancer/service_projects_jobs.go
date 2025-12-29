package freelancer

import (
	"context"
	"net/http"
	"net/url"

	"github.com/cushydigit/go-freelancer-sdk/freelancer/endpoints"
)

type ListJobsOptions struct {
	Jobs                      []int64
	JobNames                  []string
	SeoUrls                   []string
	Categories                []int64
	OnlyLocal                 *bool
	ActiveProjectCountDetails *bool
	SeoDetails                *bool
	SeoCountryName            *string
	Lang                      *string
}

// TODO: refine with typed response

// Returns a list of milestone requests.
// It maps to the `GET` `/projects/0.1/jobs` endpoint
func (s *JobsService) List(ctx context.Context, opts *ListJobsOptions) (*RawResponse, error) {
	path := endpoints.ProjectsJobs
	query := url.Values{}
	if opts != nil {
		addInt64Slice(query, "jobs[]", opts.Jobs)
		addStringSlice(query, "job_names[]", opts.JobNames)
		addStringSlice(query, "seo_urls[]", opts.SeoUrls)
		addInt64Slice(query, "categories[]", opts.Categories)
		addBool(query, "only_local", opts.OnlyLocal)
		addBool(query, "active_project_count_details", opts.ActiveProjectCountDetails)
		addBool(query, "seo_details", opts.SeoDetails)
		addString(query, "seo_country_name", opts.SeoCountryName)
		addString(query, "lang", opts.Lang)
	}
	return execute[*RawResponse](ctx, s.client, http.MethodGet, path, query, nil)
}

type SearchJobsOptions struct {
	Jobs                      []int64
	JobNames                  []string
	SeoUrls                   []string
	SeoTexts                  []string
	Categories                []int64
	OnlyLocal                 *bool
	ActiveProjectCountDetails *bool
	SeoDetails                *bool
	SeoCountryName            *string
	Lang                      *string
}

// TODO: refine with typed response

// Returns a list of jobs. Note: This performs a sub-string search for all the parameters specified on the jobs.
// It maps to the `GET` `/projects/0.1/jobs/search` endpoint
func (s *JobsService) Search(ctx context.Context, opts *SearchJobsOptions) (*RawResponse, error) {
	path := endpoints.ProjectsJobsSearch
	query := url.Values{}
	if opts != nil {
		addInt64Slice(query, "jobs[]", opts.Jobs)
		addStringSlice(query, "job_names[]", opts.JobNames)
		addStringSlice(query, "seo_urls[]", opts.SeoUrls)
		addStringSlice(query, "seo_texts[]", opts.SeoTexts)
		addInt64Slice(query, "categories[]", opts.Categories)
		addBool(query, "only_local", opts.OnlyLocal)
		addBool(query, "active_project_count_details", opts.ActiveProjectCountDetails)
		addBool(query, "seo_details", opts.SeoDetails)
		addString(query, "seo_country_name", opts.SeoCountryName)
		addString(query, "lang", opts.Lang)
	}
	return execute[*RawResponse](ctx, s.client, http.MethodGet, path, query, nil)
}

type ListJobBundlesOptions struct {
	JobBundles []int64
	Categories []int64
	Lang       *string
}

// TODO: refine with typed response

// Returns a list of job bundles. Note: Categories in this context are job bundle categories. These are not the same as job categories even though they share the same name.
// It maps to the `GET` `/projects/0.1/job_bundles` endpoint
func (s *JobsService) ListJobBundles(ctx context.Context, opts *ListJobBundlesOptions) (*RawResponse, error) {
	path := endpoints.ProjectsJobBundles
	query := url.Values{}
	if opts != nil {
		addInt64Slice(query, "job_bundles[]", opts.JobBundles)
		addInt64Slice(query, "categories[]", opts.Categories)
		addString(query, "lang", opts.Lang)
	}
	return execute[*RawResponse](ctx, s.client, http.MethodGet, path, query, nil)
}

type ListJobBundleCategoriesOptions struct {
	JobBundles []int64
	Categories []int64
	Lang       *string
}

// TODO: refine with typed response

// Returns a list of job bundle categories.
// It maps to the `GET` `/projects/0.1/job_bundle_categories` endpoint
func (s *JobsService) ListJobBundleCategories(ctx context.Context, opts *ListJobBundleCategoriesOptions) (*RawResponse, error) {
	path := endpoints.ProjectsJobBundleCategories
	query := url.Values{}
	if opts != nil {
		addInt64Slice(query, "job_bundles[]", opts.JobBundles)
		addInt64Slice(query, "categories[]", opts.Categories)
		addString(query, "lang", opts.Lang)
	}
	return execute[*RawResponse](ctx, s.client, http.MethodGet, path, query, nil)
}

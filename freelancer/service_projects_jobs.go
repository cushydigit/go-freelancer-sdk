package freelancer

import (
	"context"
	"net/http"
	"net/url"
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
	query := url.Values{}
	if opts != nil {
		for _, id := range opts.Jobs {
			addInt64(query, "jobs[]", &id)
		}
		for _, name := range opts.JobNames {
			addString(query, "job_names[]", &name)
		}
		for _, url := range opts.SeoUrls {
			addString(query, "seo_urls[]", &url)
		}
		for _, id := range opts.Categories {
			addInt64(query, "categories[]", &id)
		}
		addBool(query, "only_local", opts.OnlyLocal)
		addBool(query, "active_project_count_details", opts.ActiveProjectCountDetails)
		addBool(query, "seo_details", opts.SeoDetails)
		addString(query, "seo_country_name", opts.SeoCountryName)
		addString(query, "lang", opts.Lang)
	}

	return execute[*RawResponse](ctx, s.client, http.MethodGet, string(ProjectsJobs), query, nil)
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
	query := url.Values{}
	if opts != nil {
		for _, id := range opts.Jobs {
			addInt64(query, "jobs[]", &id)
		}
		for _, name := range opts.JobNames {
			addString(query, "job_names[]", &name)
		}
		for _, url := range opts.SeoUrls {
			addString(query, "seo_urls[]", &url)
		}
		for _, text := range opts.SeoTexts {
			addString(query, "seo_texts[]", &text)
		}
		for _, id := range opts.Categories {
			addInt64(query, "categories[]", &id)
		}
		addBool(query, "only_local", opts.OnlyLocal)
		addBool(query, "active_project_count_details", opts.ActiveProjectCountDetails)
		addBool(query, "seo_details", opts.SeoDetails)
		addString(query, "seo_country_name", opts.SeoCountryName)
		addString(query, "lang", opts.Lang)
	}
	return execute[*RawResponse](ctx, s.client, http.MethodGet, string(ProjectsJobsSearch), query, nil)
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
	query := url.Values{}
	if opts != nil {
		for _, id := range opts.JobBundles {
			addInt64(query, "job_bundles[]", &id)
		}
		for _, id := range opts.Categories {
			addInt64(query, "categories[]", &id)
		}
		addString(query, "lang", opts.Lang)
	}
	return execute[*RawResponse](ctx, s.client, http.MethodGet, string(ProjectsJobBundles), query, nil)
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
	query := url.Values{}
	if opts != nil {
		for _, id := range opts.JobBundles {
			addInt64(query, "job_bundles[]", &id)
		}
		for _, id := range opts.Categories {
			addInt64(query, "categories[]", &id)
		}
		addString(query, "lang", opts.Lang)
	}
	return execute[*RawResponse](ctx, s.client, http.MethodGet, string(ProjectsJobBundleCategories), query, nil)
}

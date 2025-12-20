package freelancer

import (
	"context"
	"encoding/json"
	"net/http"
)

type listJobs struct {
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

func (s *listJobs) Do(ctx context.Context) (*BaseResponse, error) {
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

// SetJobs sets the IDs of the jobs to filter the results.
// Only jobs with these IDs will be returned by the API.
func (s *listJobs) SetJobs(vals []int) *listJobs {
	s.jobs = vals
	return s
}

// SetJobNames sets the names of the jobs to filter the results.
// Only jobs matching these names will be returned.
func (s *listJobs) SetJobNames(vals []string) *listJobs {
	s.jobNames = vals
	return s
}

// SetSeoUrls sets the SEO URLs of the jobs to filter the results.
// Only jobs matching these SEO URLs will be returned.
func (s *listJobs) SetSeoUrls(vals []string) *listJobs {
	s.seoUrls = vals
	return s
}

// SetCategories sets the category IDs to filter the jobs.
// Only jobs within these categories will be returned.
func (s *listJobs) SetCategories(vals []int) *listJobs {
	s.categories = vals
	return s
}

// SetOnlyLocal filters the results to include only local jobs if set to true.
func (s *listJobs) SetOnlyLocal(val bool) *listJobs {
	s.onlyLocal = &val
	return s
}

// SetActiveProjectCountDetails enables or disables returning counts of active projects for each job type.
func (s *listJobs) SetActiveProjectCountDetails(val bool) *listJobs {
	s.activeProjectCountDetails = &val
	return s
}

// SetSeoDetails enables or disables returning SEO-related information for the jobs.
func (s *listJobs) SetSeoDetails(val bool) *listJobs {
	s.seoDetails = &val
	return s
}

// SetSeoCountryName sets the country name for SEO text population.
func (s *listJobs) SetSeoCountryName(val string) *listJobs {
	s.seoCountryName = &val
	return s
}

// SetLang sets the language code for the returned job values.
func (s *listJobs) SetLang(val string) *listJobs {
	s.lang = &val
	return s
}

type searchJobs struct {
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

func (s *searchJobs) Do(ctx context.Context) (*BaseResponse, error) {
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

// SetJobs sets the list of job IDs to filter the search results.
// Only jobs matching these IDs will be returned.
func (s *searchJobs) SetJobs(vals []int) *searchJobs {
	s.jobs = vals
	return s
}

// SetJobNames sets the list of job name substrings to filter the search results.
// Returns jobs where the job name contains any of the specified strings.
func (s *searchJobs) SetJobNames(vals []string) *searchJobs {
	s.jobNames = vals
	return s
}

// SetSeoUrls sets the list of SEO URL substrings to filter the search results.
// Returns jobs where the SEO URL contains any of the specified strings.
func (s *searchJobs) SetSeoUrls(vals []string) *searchJobs {
	s.seoUrls = vals
	return s
}

// SetSeoTexts sets the list of SEO text substrings to filter the search results.
// Returns jobs where the SEO text contains any of the specified strings.
func (s *searchJobs) SetSeoTexts(vals []string) *searchJobs {
	s.seoTexts = vals
	return s
}

// SetCategories sets the list of category IDs to filter the search results.
// Only jobs within the specified categories will be returned.
func (s *searchJobs) SetCategories(vals []int) *searchJobs {
	s.categories = vals
	return s
}

// SetOnlyLocal filters the search results to return only local jobs if true.
func (s *searchJobs) SetOnlyLocal(val bool) *searchJobs {
	s.onlyLocal = &val
	return s
}

// SetActiveProjectCountDetails enables returning additional active project count details for jobs.
func (s *searchJobs) SetActiveProjectCountDetails(val bool) *searchJobs {
	s.activeProjectCountDetails = &val
	return s
}

// SetSeoDetails enables returning SEO information for the returned jobs.
func (s *searchJobs) SetSeoDetails(val bool) *searchJobs {
	s.seoDetails = &val
	return s
}

// SetSeoCountryName specifies the country name for SEO details.
// The returned SEO text will include information specific to this country.
func (s *searchJobs) SetSeoCountryName(val string) *searchJobs {
	s.seoCountryName = &val
	return s
}

// SetLang sets the language code for the returned values.
func (s *searchJobs) SetLang(val string) *searchJobs {
	s.lang = &val
	return s
}

type listJobBundles struct {
	client     *Client
	jobBundles []int
	categories []int
	lang       *string
}

func (s *listJobBundles) Do(ctx context.Context) (*BaseResponse, error) {
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

// SetJobBundles sets a list of job bundle IDs to filter the results.
// These IDs correspond to specific job bundles that the API should return.
// Example: []int{1, 2}
// Returns the updated listJobBundles instance for chaining.
func (s *listJobBundles) SetJobBundles(vals []int) *listJobBundles {
	s.jobBundles = vals
	return s
}

// SetCategories sets a list of job bundle category IDs to filter the results.
// Note: These are categories of job bundles, not individual job categories.
// Example: []int{1, 2}
// Returns the updated listJobBundles instance for chaining.
func (s *listJobBundles) SetCategories(vals []int) *listJobBundles {
	s.categories = vals
	return s
}

// SetLang sets the projection language code for the returned values.
// Example: "en"
// Returns the updated listJobBundles instance for chaining.
func (s *listJobBundles) SetLang(val string) *listJobBundles {
	s.lang = &val
	return s
}

type listJobBundleCategories struct {
	client     *Client
	jobBundles []int
	categories []int
	lang       *string
}

func (s *listJobBundleCategories) Do(ctx context.Context) (*BaseResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: string(PROJECTS_JOB_BUNDLE_CATEGORIES),
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

// SetJobBundles sets a list of job bundle IDs to filter the results.
// These IDs correspond to specific job bundles that the API should return.
// Example: []int{1, 2}
// Returns the updated listJobBundleCategories instance for chaining.
func (s *listJobBundleCategories) SetJobBundles(vals []int) *listJobBundleCategories {
	s.jobBundles = vals
	return s
}

// SetCategories sets a list of job bundle category IDs to filter the results.
// Note: These are categories of job bundles, not individual job categories.
// Example: []int{1, 2}
// Returns the updated listJobBundleCategories instance for chaining.
func (s *listJobBundleCategories) SetCategories(vals []int) *listJobBundleCategories {
	s.categories = vals
	return s
}

// SetLang sets the projection language code for the returned values.
// Example: "en"
// Returns the updated listJobBundleCategories instance for chaining.
func (s *listJobBundleCategories) SetLang(val string) *listJobBundleCategories {
	s.lang = &val
	return s
}

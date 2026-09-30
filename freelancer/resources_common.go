package freelancer

import (
	"context"
	"net/http"

	"github.com/cushydigit/go-freelancer-sdk/freelancer/internal/endpoints"
	rr "github.com/cushydigit/go-freelancer-sdk/freelancer/reqres"
)

// --------------------------------------
// Resource-Common
// --------------------------------------

// ListCountries fetches a list of countries from the Freelancer API.
// It maps to the `GET` `/common/0.1/countries` endpoint.
func (r *Common) ListCountries(
	ctx context.Context,
	opts *rr.ListCountriesOptions,
) (*rr.ListCountriesResponse, *ResponseMeta, error) {
	return execute[*rr.ListCountriesResponse](
		ctx,
		r.client,
		http.MethodGet,
		endpoints.CommonCountries,
		opts,
		nil,
	)
}

// ListTimezones fetches a list of timezones from the Freelancer API.
// It maps to the `GET` `/common/0.1/timezones` endpoint.
func (r *Common) ListTimezones(
	ctx context.Context,
	opts *rr.ListTimezonesOptions,
) (*rr.ListTimezonesResponse, *ResponseMeta, error) {
	return execute[*rr.ListTimezonesResponse](
		ctx,
		r.client,
		http.MethodGet,
		endpoints.CommonTimezones,
		opts,
		nil,
	)
}

// Returns a list of currencies.currency_codes and currency_ids are incompatible with each other.
// It maps to the `GET` `/projects/0.1/currencies` endpoint
func (r *Common) ListCurrencies(
	ctx context.Context,
	opts *rr.ListCurrenciesOptions,
) (*rr.ListCurrenciesResponse, *ResponseMeta, error) {
	return execute[*rr.ListCurrenciesResponse](
		ctx,
		r.client,
		http.MethodGet,
		endpoints.Currencies,
		opts,
		nil,
	)
}

// Returns a list of categories. If job_details is set, a map of category IDs to jobs in those categories.
// it maps to the `GET` `/projects/0.1/categories` endpoint
func (r *Common) ListCategories(
	ctx context.Context,
	opts *rr.ListCategoriesOptions,
) (*rr.ListCategoriesResponse, *ResponseMeta, error) {
	return execute[*rr.ListCategoriesResponse](
		ctx,
		r.client,
		http.MethodGet,
		endpoints.Categories,
		opts,
		nil,
	)
}

// Returns a list of budgets with the specified currencies.currency_codes and currency_ids are incompatible with each other.
// It maps to the `GET` `/projects/0.1/budgets` endpoint
func (r *Common) ListBudgets(
	ctx context.Context,
	opts *rr.ListBudgetsOptions,
) (*rr.ListBudgetsResponse, *ResponseMeta, error) {
	return execute[*rr.ListBudgetsResponse](
		ctx,
		r.client,
		http.MethodGet,
		endpoints.Budgets,
		opts,
		nil,
	)
}

// Returns a list of jobs.
// It maps to the `GET` `/projects/0.1/jobs` endpoint
func (r *Common) ListJobs(
	ctx context.Context,
	opts *rr.ListJobsOptions,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		r.client,
		http.MethodGet,
		endpoints.Jobs,
		opts,
		nil,
	)
}

// Returns a list of jobs. Note: This performs a sub-string search for all the parameters specified on the jobs.
// It maps to the `GET` `/projects/0.1/jobs/search` endpoint
func (r *Common) SearchJobs(
	ctx context.Context,
	opts *rr.SearchJobsOptions,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		r.client,
		http.MethodGet,
		endpoints.JobsSearch,
		opts,
		nil,
	)
}

// Returns a list of job bundles. Note: Categories in this context are job bundle categories. These are not the same as job categories even though they share the same name.
// It maps to the `GET` `/projects/0.1/job_bundles` endpoint
func (r *Common) ListJobBundles(
	ctx context.Context,
	opts *rr.ListJobBundlesOptions,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		r.client,
		http.MethodGet,
		endpoints.JobBundles,
		opts,
		nil,
	)
}

// Returns a list of job bundle categories.
// It maps to the `GET` `/projects/0.1/job_bundle_categories` endpoint
func (r *Common) ListJobBundleCategories(
	ctx context.Context,
	opts *rr.ListJobBundleCategoriesOptions,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		r.client,
		http.MethodGet,
		endpoints.JobBundleCategories,
		opts,
		nil,
	)
}

package freelancer

import (
	"context"
	"net/http"

	"github.com/cushydigit/go-freelancer-sdk/freelancer/internal/endpoints"
	rr "github.com/cushydigit/go-freelancer-sdk/freelancer/reqres"
)

// ListCountries fetches a list of countries from the Freelancer API.
// It maps to the `GET` `/common/0.1/countries` endpoint.
func (s *CommonService) ListCountries(
	ctx context.Context,
	opts *rr.ListCountriesOptions,
) (*rr.ListCountriesResponse, *ResponseMeta, error) {
	return execute[*rr.ListCountriesResponse](
		ctx,
		s.client,
		http.MethodGet,
		endpoints.CommonCountries,
		opts,
		nil,
	)
}

// ListTimezones fetches a list of timezones from the Freelancer API.
// It maps to the `GET` `/common/0.1/timezones` endpoint.
func (s *CommonService) ListTimezones(
	ctx context.Context,
	opts *rr.ListTimezonesOptions,
) (*rr.ListTimezonesResponse, *ResponseMeta, error) {
	return execute[*rr.ListTimezonesResponse](
		ctx,
		s.client,
		http.MethodGet,
		endpoints.CommonTimezones,
		opts,
		nil,
	)
}

// Returns a list of currencies.currency_codes and currency_ids are incompatible with each other.
// It maps to the `GET` `/projects/0.1/currencies` endpoint
func (s *CommonService) ListCurrencies(
	ctx context.Context,
	opts *rr.ListCurrenciesOptions,
) (*rr.ListCurrenciesResponse, *ResponseMeta, error) {
	return execute[*rr.ListCurrenciesResponse](
		ctx,
		s.client,
		http.MethodGet,
		endpoints.Currencies,
		opts,
		nil,
	)
}

// Returns a list of categories. If job_details is set, a map of category IDs to jobs in those categories.
// it maps to the `GET` `/projects/0.1/categories` endpoint
func (s *CommonService) ListCategories(
	ctx context.Context,
	opts *rr.ListCategoriesOptions,
) (*rr.ListCategoriesResponse, *ResponseMeta, error) {
	return execute[*rr.ListCategoriesResponse](
		ctx,
		s.client,
		http.MethodGet,
		endpoints.Categories,
		opts,
		nil,
	)
}

// Returns a list of budgets with the specified currencies.currency_codes and currency_ids are incompatible with each other.
// It maps to the `GET` `/projects/0.1/budgets` endpoint
func (s *CommonService) ListBudgets(
	ctx context.Context,
	opts *rr.ListBudgetsOptions,
) (*rr.ListBudgetsResponse, *ResponseMeta, error) {
	return execute[*rr.ListBudgetsResponse](
		ctx,
		s.client,
		http.MethodGet,
		endpoints.Budgets,
		opts,
		nil,
	)
}

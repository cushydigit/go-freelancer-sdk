package freelancer

import (
	"context"
	"net/http"

	"github.com/cushydigit/go-freelancer-sdk/freelancer/internal/endpoints"
	"github.com/cushydigit/go-freelancer-sdk/freelancer/internal/query"
	rr "github.com/cushydigit/go-freelancer-sdk/freelancer/reqres"
)

// ListCountries fetches a list of countries from the Freelancer API.
// It maps to the `GET` `/common/0.1/countries` endpoint.
func (s *CommonService) ListCountries(ctx context.Context, opts *rr.ListCountriesOptions) (*rr.ListCountriesResponse, error) {
	p := endpoints.CommonCountries
	q := query.Values(opts)
	return execute[*rr.ListCountriesResponse](ctx, s.client, http.MethodGet, p, q, nil)
}

// ListTimezones fetches a list of timezones from the Freelancer API.
// It maps to the `GET` `/common/0.1/timezones` endpoint.
func (s *CommonService) ListTimezones(ctx context.Context, opts *rr.ListTimezonesOptions) (*rr.ListTimezonesResponse, error) {
	p := endpoints.CommonTimezones
	q := query.Values(opts)
	return execute[*rr.ListTimezonesResponse](ctx, s.client, http.MethodGet, p, q, nil)
}

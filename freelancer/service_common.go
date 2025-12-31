package freelancer

import (
	"context"
	"net/http"

	"github.com/cushydigit/go-freelancer-sdk/freelancer/internal/endpoints"
	"github.com/cushydigit/go-freelancer-sdk/freelancer/internal/query"
)

// CountriesOptions holds optional filters for the ListCountries request.
type ListCountriesOptions struct {
	ExtraDetails *bool `url:"extra_details"`
}

// ListCountries fetches a list of countries from the Freelancer API.
// It maps to the `GET` `/common/0.1/countries` endpoint.
func (s *CommonService) ListCountries(ctx context.Context, opts *ListCountriesOptions) (*ListCountriesResponse, error) {
	p := endpoints.CommonCountries
	q := query.Values(opts)
	return execute[*ListCountriesResponse](ctx, s.client, http.MethodGet, p, q, nil)
}

// TimezonesOptions holds optional filters for the ListTimezones request.
type ListTimezonesOptions struct {
	Timezones     []int64  `url:"timezones[]"`
	TimezoneNames []string `url:"timezone_names[]"`
}

// ListTimezones fetches a list of timezones from the Freelancer API.
// It maps to the `GET` `/common/0.1/timezones` endpoint.
func (s *CommonService) ListTimezones(ctx context.Context, opts *ListTimezonesOptions) (*ListTimezonesResponse, error) {
	p := endpoints.CommonTimezones
	q := query.Values(opts)
	return execute[*ListTimezonesResponse](ctx, s.client, http.MethodGet, p, q, nil)
}

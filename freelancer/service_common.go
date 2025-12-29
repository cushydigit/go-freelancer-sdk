package freelancer

import (
	"context"
	"net/http"
	"net/url"

	"github.com/cushydigit/go-freelancer-sdk/freelancer/endpoints"
)

// CountriesOptions holds optional filters for the ListCountries request.
type ListCountriesOptions struct {
	ExtraDetails *bool
}

// ListCountries fetches a list of countries from the Freelancer API.
// It maps to the `GET` `/common/0.1/countries` endpoint.
func (s *CommonService) ListCountries(ctx context.Context, opts *ListCountriesOptions) (*ListCountriesResponse, error) {
	path := endpoints.CommonCountries
	query := url.Values{}
	if opts != nil {
		addBool(query, "extra_details", opts.ExtraDetails)
	}
	return execute[*ListCountriesResponse](ctx, s.client, http.MethodGet, path, query, nil)
}

// TimezonesOptions holds optional filters for the ListTimezones request.
type ListTimezonesOptions struct {
	Timezones     []int64
	TimezoneNames []string
}

// ListTimezones fetches a list of timezones from the Freelancer API.
// It maps to the `GET` `/common/0.1/timezones` endpoint.
func (s *CommonService) ListTimezones(ctx context.Context, opts *ListTimezonesOptions) (*ListTimezonesResponse, error) {
	path := endpoints.CommonTimezones
	query := url.Values{}
	if opts != nil {
		addInt64Slice(query, "timezones[]", opts.Timezones)
		addStringSlice(query, "timezone_names[]", opts.TimezoneNames)
	}

	return execute[*ListTimezonesResponse](ctx, s.client, http.MethodGet, path, query, nil)
}

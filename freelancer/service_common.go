package freelancer

import (
	"context"
	"net/http"
	"net/url"
	"strconv"

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
		if opts.ExtraDetails != nil {
			query.Add("extra_details", strconv.FormatBool(*opts.ExtraDetails))
		}
	}
	return execute[*ListCountriesResponse](ctx, s.client, http.MethodGet, path, query, nil)
}

// TimezonesOptions holds optional filters for the ListTimezones request.
type ListTimezonesOptions struct {
	Timezones     []int
	TimezoneNames []string
}

// ListTimezones fetches a list of timezones from the Freelancer API.
// It maps to the `GET` `/common/0.1/timezones` endpoint.
func (s *CommonService) ListTimezones(ctx context.Context, opts *ListTimezonesOptions) (*ListTimezonesResponse, error) {
	path := endpoints.CommonTimezones
	query := url.Values{}
	if opts != nil {
		for _, val := range opts.Timezones {
			query.Add("timezones[]", strconv.Itoa(val))
		}
		for _, val := range opts.TimezoneNames {
			query.Add("timezone_names[]", val)
		}
	}

	return execute[*ListTimezonesResponse](ctx, s.client, http.MethodGet, path, query, nil)
}

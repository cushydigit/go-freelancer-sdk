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

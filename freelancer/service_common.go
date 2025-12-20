package freelancer

import (
	"context"
	"net/http"
)

type listCountries struct {
	client       *Client
	extraDetails *bool
}

func (s *listCountries) Do(ctx context.Context) (*ListCountriesResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: string(COMMON_COUNTRIES),
	}
	if s.extraDetails != nil {
		r.setParam("extra_details", *s.extraDetails)
	}
	return execute[*ListCountriesResponse](ctx, s.client, r)
}

// SetExtraDetails enables or disables the `extra_details` query parameter.
//
// When set to true, the API response will include additional country metadata
// such as phone codes, SEO URLs, and language information.
func (c *listCountries) SetExtraDetails(val bool) *listCountries {
	c.extraDetails = &val
	return c
}

type listTimezones struct {
	client        *Client
	timezones     []int
	timezoneNames []string
}

func (s *listTimezones) Do(ctx context.Context) (*ListTimezonesResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: string(COMMON_TIMEZONES),
	}
	if len(s.timezones) > 0 {
		r.setParam("timezones", s.timezones)
	}
	if len(s.timezoneNames) > 0 {
		r.setParam("timezone_names", s.timezoneNames)
	}
	return execute[*ListTimezonesResponse](ctx, s.client, r)
}

// SetTimezones sets an optional filter to return only timezones
// matching the given timezone IDs.
//
// This corresponds to the `timezones[]` query parameter.
func (s *listTimezones) SetTimezones(vals []int) *listTimezones {
	s.timezones = vals
	return s
}

// SetTimezoneNames sets an optional filter to return only timezones
// matching the given timezone names (e.g. "Australia/Sydney").
//
// This corresponds to the `timezone_names[]` query parameter.
func (s *listTimezones) SetTimezoneNames(vals []string) *listTimezones {
	s.timezoneNames = vals
	return s
}

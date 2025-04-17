package freelancer

import (
	"context"
	"encoding/json"
	"net/http"
)

type ListCountriesService struct {
	client       *Client
	extraDetails bool
}

type ListCountriesResponse struct {
	Status    string          `json:"status"`
	RequestID string          `json:"request_id,omitempty"` // Optional
	Result    CountriesResult `json:"result"`
}
type CountriesResult struct {
	Countries []Country `json:"countries"`
}

// Do perform GET request on endpoint "common/0.1/countries/"
func (c *ListCountriesService) Do(ctx context.Context) (*ListCountriesResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/common/0.1/countries",
	}
	if c.extraDetails {
		r.setParam("extra_details", c.extraDetails)
	}
	data, err := c.client.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := &ListCountriesResponse{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *ListCountriesService) SetExtraDetails(extraDetails bool) *ListCountriesService {
	c.extraDetails = extraDetails
	return c
}

type ListTimezonesService struct {
	client        *Client
	timezones     []int
	timezoneNames []string
}

type ListTimezonesResponse struct {
	Status    string          `json:"status"`
	RequestID string          `json:"request_id,omitempty"` // Optional
	Result    TimezonesResult `json:"result"`
}
type TimezonesResult struct {
	ID       int    `json:"id"`                 // "number" represented as float64
	Country  string `json:"country,omitempty"`  // Optional
	Timezone string `json:"timezone,omitempty"` // Optional
	Offset   int    `json:"offset,omitempty"`   // Optional, "Decimal"
}

// Do perform GET request on endpoint "common/0.1/timezones/"
func (c *ListTimezonesService) Do(ctx context.Context) (*ListTimezonesResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/common/0.1/timezones",
	}
	if len(c.timezones) > 0 {
		r.setParam("timezones", c.timezones)
	}
	if len(c.timezoneNames) > 0 {
		r.setParam("timezone_names", c.timezoneNames)
	}
	data, err := c.client.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := &ListTimezonesResponse{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *ListTimezonesService) Timezones(timezones []int) *ListTimezonesService {
	s.timezones = timezones
	return s
}

func (s *ListTimezonesService) TimezoneNames(timezoneNames []string) *ListTimezonesService {
	s.timezoneNames = timezoneNames
	return s
}

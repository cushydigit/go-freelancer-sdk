package freelancer

import (
	"context"
	"encoding/json"
	"net/http"
)

type timezonesListService struct {
	client        *Client
	timezones     []int
	timezoneNames []string
}

func (c *timezonesListService) Do(ctx context.Context) (*BaseResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: string(COMMON_TIMEZONES),
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
	res := &BaseResponse{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// optional - filter Returns timezones with the given timezone IDs.
func (s *timezonesListService) SetTimezones(vals []int) *timezonesListService {
	s.timezones = vals
	return s
}

// optional - filter Returns timezones with the given timezone names.
func (s *timezonesListService) SetTimezoneNames(vals []string) *timezonesListService {
	s.timezoneNames = vals
	return s
}

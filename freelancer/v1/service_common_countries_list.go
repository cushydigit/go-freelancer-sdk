package freelancer

import (
	"context"
	"encoding/json"
	"net/http"
)

type countriesListService struct {
	client       *Client
	extraDetails *bool
}

func (s *countriesListService) Do(ctx context.Context) (*BaseResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: string(COMMON_COUNTRIES),
	}
	if s.extraDetails != nil {
		r.setParam("extra_details", *s.extraDetails)
	}
	data, err := s.client.callAPI(ctx, r)
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

func (c *countriesListService) SetExtraDetails(val bool) *countriesListService {
	c.extraDetails = &val
	return c
}

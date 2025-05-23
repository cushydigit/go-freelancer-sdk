package freelancer

import (
	"context"
	"encoding/json"
	"net/http"
)

// Returns the project fees for a given list of currencies
type ListBidUpgradeFeesService struct {
	client      *Client
	currencies  []int
	taxIncluded *bool
}

func (s *ListBidUpgradeFeesService) Do(ctx context.Context) (*BaseResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: string(PROJECTS_BIDS_FEES),
	}

	for _, currency := range s.currencies {
		r.addParam("currencies[]", currency)
	}
	if s.taxIncluded != nil {
		r.addParam("tax_included", *s.taxIncluded)
	}

	data, err := s.client.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}

	resp := &BaseResponse{}
	if err := json.Unmarshal(data, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *ListBidUpgradeFeesService) SetCurrencies(vals []int) *ListBidUpgradeFeesService {
	s.currencies = vals
	return s
}

func (s *ListBidUpgradeFeesService) SetTaxIncluded(val bool) *ListBidUpgradeFeesService {
	s.taxIncluded = &val
	return s
}

package freelancer

import (
	"context"
	"encoding/json"
	"net/http"
)

type bidUpgradeFeesListService struct {
	client      *Client
	currencies  []int
	taxIncluded *bool
}

func (s *bidUpgradeFeesListService) Do(ctx context.Context) (*BaseResponse, error) {
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

func (s *bidUpgradeFeesListService) SetCurrencies(vals []int) *bidUpgradeFeesListService {
	s.currencies = vals
	return s
}

func (s *bidUpgradeFeesListService) SetTaxIncluded(val bool) *bidUpgradeFeesListService {
	s.taxIncluded = &val
	return s
}

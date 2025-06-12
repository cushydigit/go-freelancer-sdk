package freelancer

import (
	"context"
	"encoding/json"
	"net/http"
)

type currenciesListService struct {
	client                    *Client
	currencyCodes             []string
	currencyIDs               []int
	includeExternalCurrencies *bool
}

func (s *currenciesListService) Do(ctx context.Context) (*BaseResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/projects/0.1/currencies",
	}
	if s.currencyCodes != nil {
		r.addParam("currency_codes", s.currencyCodes)
	}
	if s.currencyIDs != nil {
		r.addParam("currency_ids", s.currencyIDs)
	}
	if s.includeExternalCurrencies != nil {
		r.addParam("include_external_currencies", s.includeExternalCurrencies)
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

func (s *currenciesListService) SetCurrencyCodes(vals []string) *currenciesListService {
	s.currencyCodes = vals
	return s
}

func (s *currenciesListService) SetCurrencyIDs(vals []int) *currenciesListService {
	s.currencyIDs = vals
	return s
}

func (s *currenciesListService) SetIncludeExternalCurrencies(val bool) *currenciesListService {
	s.includeExternalCurrencies = &val
	return s
}

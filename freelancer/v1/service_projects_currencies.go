package freelancer

import (
	"context"
	"encoding/json"
	"net/http"
)

type ListCurrenciesService struct {
	client                    *Client
	currencyCodes             []string
	currencyIDs               []int
	includeExternalCurrencies bool
}

type ListCurrenciesResponse struct {
	Status    string           `json:"status"`
	RequestID string           `json:"request_id,omitempty"` // Optional
	Result    CurrenciesResult `json:"result"`
}

type CurrenciesResult struct {
	Currencies []Currency `json:"currencies"`
}

// Do perform GET request on endpoint "projects/0.1/currencies/"
func (s *ListCurrenciesService) Do(ctx context.Context) (*ListCurrenciesResponse, error) {
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
	if s.includeExternalCurrencies {
		r.addParam("include_external_currencies", s.includeExternalCurrencies)
	}
	data, err := s.client.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := &ListCurrenciesResponse{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *ListCurrenciesService) SetCurrencyCodes(codes []string) *ListCurrenciesService {
	s.currencyCodes = codes
	return s
}

func (s *ListCurrenciesService) SetCurrencyIDs(ids []int) *ListCurrenciesService {
	s.currencyIDs = ids
	return s
}

func (s *ListCurrenciesService) SetIncludeExternalCurrencies(include bool) *ListCurrenciesService {
	s.includeExternalCurrencies = include
	return s
}

package freelancer

import (
	"context"
	"encoding/json"
	"net/http"
)

type budgetsListService struct {
	client          *Client
	currencyCodes   []string
	currencyIDs     []int
	projectType     ProjectType
	lang            string
	currencyDetails bool
}

func (s *budgetsListService) Do(ctx context.Context) (*BaseResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/projects/0.1/budgets",
	}
	if s.currencyCodes != nil {
		r.addParam("currency_codes", s.currencyCodes)
	}
	if s.currencyIDs != nil {
		r.addParam("currency_ids", s.currencyIDs)
	}
	if s.projectType != "" {
		r.addParam("project_type", s.projectType)
	}
	if s.lang != "" {
		r.addParam("lang", s.lang)
	}
	if s.currencyDetails {
		r.addParam("currency_details", s.currencyDetails)
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

func (s *budgetsListService) SetCurrencyCodes(currencyCodes []string) *budgetsListService {
	s.currencyCodes = currencyCodes
	return s
}

func (s *budgetsListService) SetCurrencyIDs(currencyIDs []int) *budgetsListService {
	s.currencyIDs = currencyIDs
	return s
}

func (s *budgetsListService) SetProjectType(projectType ProjectType) *budgetsListService {
	s.projectType = projectType
	return s
}

func (s *budgetsListService) SetLang(lang string) *budgetsListService {
	s.lang = lang
	return s
}

func (s *budgetsListService) SetCurrencyDetails(currencyDetails bool) *budgetsListService {
	s.currencyDetails = currencyDetails
	return s
}

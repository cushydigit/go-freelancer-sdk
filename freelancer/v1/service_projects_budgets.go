package freelancer

import (
	"context"
	"encoding/json"
	"net/http"
)

type ListBudgetsService struct {
	client          *Client
	currencyCodes   []string
	currencyIDs     []int
	projectType     ProjectType
	lang            string
	currencyDetails bool
}

type ListBudgetsResponse struct {
	Status    string        `json:"status"`
	RequestID string        `json:"request_id,omitempty"` // Optional
	Result    BudgetsResult `json:"result"`
}

type BudgetsResult struct {
	Budgets []Budget `json:"budgets"`
}

// Do perform GET request on endpoint "projects/0.1/budgets/"
func (s *ListBudgetsService) Do(ctx context.Context) (*ListBudgetsResponse, error) {
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
	res := &ListBudgetsResponse{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *ListBudgetsService) SetCurrencyCodes(currencyCodes []string) *ListBudgetsService {
	s.currencyCodes = currencyCodes
	return s
}

func (s *ListBudgetsService) SetCurrencyIDs(currencyIDs []int) *ListBudgetsService {
	s.currencyIDs = currencyIDs
	return s
}

func (s *ListBudgetsService) SetProjectType(projectType ProjectType) *ListBudgetsService {
	s.projectType = projectType
	return s
}

func (s *ListBudgetsService) SetLang(lang string) *ListBudgetsService {
	s.lang = lang
	return s
}

func (s *ListBudgetsService) SetCurrencyDetails(currencyDetails bool) *ListBudgetsService {
	s.currencyDetails = currencyDetails
	return s
}

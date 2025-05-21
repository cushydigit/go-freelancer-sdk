package freelancer

import (
	"context"
	"encoding/json"
	"net/http"
)

// Returns the project fees for a given list of currencies
type ListUpgradeFeesService struct {
	client             *Client
	currencies         []int
	project            *int64
	freeUpgradeDetails *bool
	taxIncluded        *bool
}

func (s *ListUpgradeFeesService) Do(ctx context.Context) (*BaseResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: string(PROJECTS_PROJECTS_FEES),
	}

	for _, currency := range s.currencies {
		r.addParam("currencies[]", currency)
	}
	if s.project != nil {
		r.addParam("project", *s.project)
	}
	if s.freeUpgradeDetails != nil {
		r.addParam("fee_upgrade_details", *s.freeUpgradeDetails)
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

func (s *ListUpgradeFeesService) SetCurrencies(vals []int) *ListUpgradeFeesService {
	s.currencies = vals
	return s
}

func (s *ListUpgradeFeesService) SetProject(val int64) *ListUpgradeFeesService {
	s.project = &val
	return s
}

func (s *ListUpgradeFeesService) SetFreeUpgradeDetails(val bool) *ListUpgradeFeesService {
	s.freeUpgradeDetails = &val
	return s
}

func (s *ListUpgradeFeesService) SetTaxIncluded(val bool) *ListUpgradeFeesService {
	s.taxIncluded = &val
	return s
}

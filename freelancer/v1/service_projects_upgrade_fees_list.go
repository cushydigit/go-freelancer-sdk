package freelancer

import (
	"context"
	"encoding/json"
	"net/http"
)

// Returns the project fees for a given list of currencies
type upgradeFeesListService struct {
	client             *Client
	currencies         []int
	project            *int64
	freeUpgradeDetails *bool
	taxIncluded        *bool
}

func (s *upgradeFeesListService) Do(ctx context.Context) (*BaseResponse, error) {
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

func (s *upgradeFeesListService) SetCurrencies(vals []int) *upgradeFeesListService {
	s.currencies = vals
	return s
}

func (s *upgradeFeesListService) SetProject(val int64) *upgradeFeesListService {
	s.project = &val
	return s
}

func (s *upgradeFeesListService) SetFreeUpgradeDetails(val bool) *upgradeFeesListService {
	s.freeUpgradeDetails = &val
	return s
}

func (s *upgradeFeesListService) SetTaxIncluded(val bool) *upgradeFeesListService {
	s.taxIncluded = &val
	return s
}

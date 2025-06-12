package freelancer

import (
	"context"
	"encoding/json"
	"net/http"
)

type hourlyContractInfoGetService struct {
	client            *Client
	projectIDs        []int64
	bidderIDs         []int64
	hourlyContractIDs []int
	projectOwnerIDs   []int64
	// teamsFilter  / not know what exactly this parameter do or is
	billingDetails *bool
	invoiceDetails *bool
}

func (s *hourlyContractInfoGetService) Do(ctx context.Context) (*BaseResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: string(PROJECTS_HOURLY_CONTRACTS),
	}

	for _, val := range s.projectIDs {
		r.addParam("project_ids[]", val)
	}
	for _, val := range s.bidderIDs {
		r.addParam("bidder_ids[]", val)
	}
	for _, val := range s.hourlyContractIDs {
		r.addParam("hourly_contract_ids[]", val)
	}
	for _, val := range s.projectOwnerIDs {
		r.addParam("project_owner_ids[]", val)
	}
	if s.billingDetails != nil {
		r.addParam("billing_details", *s.billingDetails)
	}
	if s.invoiceDetails != nil {
		r.addParam("invoice_details", *s.invoiceDetails)
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

func (s *hourlyContractInfoGetService) SetProjectIDs(vals []int64) *hourlyContractInfoGetService {
	s.projectIDs = vals
	return s
}

func (s *hourlyContractInfoGetService) SetBidderIDs(vals []int64) *hourlyContractInfoGetService {
	s.bidderIDs = vals
	return s
}

func (s *hourlyContractInfoGetService) SetProjectOwnerIDs(vals []int64) *hourlyContractInfoGetService {
	s.projectOwnerIDs = vals
	return s
}

func (s *hourlyContractInfoGetService) SetHourlyContractIDs(vals []int) *hourlyContractInfoGetService {
	s.hourlyContractIDs = vals
	return s
}

func (s *hourlyContractInfoGetService) SetBillingDetails(val bool) *hourlyContractInfoGetService {
	s.billingDetails = &val
	return s
}

func (s *hourlyContractInfoGetService) SetInvoiceDetails(val bool) *hourlyContractInfoGetService {
	s.invoiceDetails = &val
	return s
}

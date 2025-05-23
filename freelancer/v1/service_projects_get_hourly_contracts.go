package freelancer

import (
	"context"
	"encoding/json"
	"net/http"
)

// Fetch the hourly contract matching the desired query
type GetHourlyContractInformationService struct {
	client            *Client
	projectIDs        []int64
	bidderIDs         []int64
	hourlyContractIDs []int
	porjectOwnerIDs   []int64
	// teamsFilter  / not know what exactly this parameter do or is
	billingDetails *bool
	invoiceDetails *bool
}

func (s *GetHourlyContractInformationService) Do(ctx context.Context) (*BaseResponse, error) {
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
	for _, val := range s.porjectOwnerIDs {
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

func (s *GetHourlyContractInformationService) SetProjectIDs(vals []int64) *GetHourlyContractInformationService {
	s.projectIDs = vals
	return s
}

func (s *GetHourlyContractInformationService) SetBidderIDs(vals []int64) *GetHourlyContractInformationService {
	s.bidderIDs = vals
	return s
}

func (s *GetHourlyContractInformationService) SetHourlyContractIDs(vals []int) *GetHourlyContractInformationService {
	s.hourlyContractIDs = vals
	return s
}

func (s *GetHourlyContractInformationService) SetBillingDetails(val bool) *GetHourlyContractInformationService {
	s.billingDetails = &val
	return s
}

func (s *GetHourlyContractInformationService) SetInvoiceDetails(val bool) *GetHourlyContractInformationService {
	s.invoiceDetails = &val
	return s
}

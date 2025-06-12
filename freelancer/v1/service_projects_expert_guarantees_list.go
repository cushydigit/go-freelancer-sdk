package freelancer

import (
	"context"
	"encoding/json"
	"net/http"
)

type expertGuaranteesListService struct {
	client           *Client
	expertGuarantees []int64
	projects         []int64
	projectOwners    []int64
	bidders          []int64
	bids             []int
	statuses         []ExpertGuaranteesStatus
	fromTime         *int64
	toTime           *int64
	offset           *int
	limit            *int
}

func (s *expertGuaranteesListService) Do(ctx context.Context) (*BaseResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: string(PROJECTS_EXPERT_GUARANTEES),
	}

	for _, val := range s.bids {
		r.addParam("bids[]", val)
	}
	for _, val := range s.projects {
		r.addParam("projects[]", val)
	}
	for _, val := range s.bidders {
		r.addParam("bidders[]", val)
	}
	for _, val := range s.projectOwners {
		r.addParam("project_owners[]", val)
	}
	for _, val := range s.bids {
		r.addParam("bids[]", val)
	}
	for _, val := range s.statuses {
		r.addParam("statuses[]", val)
	}
	if s.fromTime != nil {
		r.addParam("from_time", *s.fromTime)
	}
	if s.toTime != nil {
		r.addParam("to_time", *s.toTime)
	}
	if s.limit != nil {
		r.addParam("limit", *s.limit)
	}
	if s.offset != nil {
		r.addParam("offset", *s.offset)
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

func (s *expertGuaranteesListService) SetExpertGuarantees(vals []int64) *expertGuaranteesListService {
	s.expertGuarantees = vals
	return s
}

func (s *expertGuaranteesListService) SetBids(vals []int) *expertGuaranteesListService {
	s.bids = vals
	return s
}

func (s *expertGuaranteesListService) SetProjects(vals []int64) *expertGuaranteesListService {
	s.projects = vals
	return s
}

func (s *expertGuaranteesListService) SetBidders(vals []int64) *expertGuaranteesListService {
	s.bidders = vals
	return s
}

func (s *expertGuaranteesListService) SetProjectOwners(vals []int64) *expertGuaranteesListService {
	s.projectOwners = vals
	return s
}

func (s *expertGuaranteesListService) SetStatuses(vals []ExpertGuaranteesStatus) *expertGuaranteesListService {
	s.statuses = vals
	return s
}

func (s *expertGuaranteesListService) SetFromTime(val int64) *expertGuaranteesListService {
	s.fromTime = &val
	return s
}

func (s *expertGuaranteesListService) SetToTime(val int64) *expertGuaranteesListService {
	s.toTime = &val
	return s
}

func (s *expertGuaranteesListService) SetLimit(val int) *expertGuaranteesListService {
	s.limit = &val
	return s
}

func (s *expertGuaranteesListService) SetOffset(val int) *expertGuaranteesListService {
	s.offset = &val
	return s
}

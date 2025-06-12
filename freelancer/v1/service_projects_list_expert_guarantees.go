package freelancer

import (
	"context"
	"encoding/json"
	"net/http"
)

type listExpertGuaranteesService struct {
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

func (s *listExpertGuaranteesService) Do(ctx context.Context) (*BaseResponse, error) {
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

func (s *listExpertGuaranteesService) SetExpertGuarantees(vals []int64) *listExpertGuaranteesService {
	s.expertGuarantees = vals
	return s
}

func (s *listExpertGuaranteesService) SetBids(vals []int) *listExpertGuaranteesService {
	s.bids = vals
	return s
}

func (s *listExpertGuaranteesService) SetProjects(vals []int64) *listExpertGuaranteesService {
	s.projects = vals
	return s
}

func (s *listExpertGuaranteesService) SetBidders(vals []int64) *listExpertGuaranteesService {
	s.bidders = vals
	return s
}

func (s *listExpertGuaranteesService) SetProjectOwners(vals []int64) *listExpertGuaranteesService {
	s.projectOwners = vals
	return s
}

func (s *listExpertGuaranteesService) SetStatuses(vals []ExpertGuaranteesStatus) *listExpertGuaranteesService {
	s.statuses = vals
	return s
}

func (s *listExpertGuaranteesService) SetFromTime(val int64) *listExpertGuaranteesService {
	s.fromTime = &val
	return s
}

func (s *listExpertGuaranteesService) SetToTime(val int64) *listExpertGuaranteesService {
	s.toTime = &val
	return s
}

func (s *listExpertGuaranteesService) SetLimit(val int) *listExpertGuaranteesService {
	s.limit = &val
	return s
}

func (s *listExpertGuaranteesService) SetOffset(val int) *listExpertGuaranteesService {
	s.offset = &val
	return s
}

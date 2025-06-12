package freelancer

import (
	"context"
	"encoding/json"
	"net/http"
)

type reputationListService struct {
	client       *Client
	users        []int64
	jobs         []int32
	role         *RoleType
	jobHistory   *bool
	projectStats *bool
	rehireRates  *bool
}

func (s *reputationListService) Do(ctx context.Context) (*BaseResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: string(USERS_REPUTATIONS),
	}

	for _, id := range s.users {
		r.addParam("users[]", id)
	}
	for _, id := range s.jobs {
		r.addParam("jobs[]", id)
	}
	if s.role != nil {
		r.addParam("role", *s.role)
	}
	if s.jobHistory != nil {
		r.addParam("job_history", *s.jobHistory)
	}
	if s.projectStats != nil {
		r.addParam("project_stats", *s.projectStats)
	}
	if s.rehireRates != nil {
		r.addParam("rehire_rates", *s.rehireRates)
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

func (s *reputationListService) SetUsers(users []int64) *reputationListService {
	s.users = users
	return s
}

func (s *reputationListService) SetJobs(jobs []int32) *reputationListService {
	s.jobs = jobs
	return s
}

func (s *reputationListService) SetRole(val RoleType) *reputationListService {
	s.role = &val
	return s
}

func (s *reputationListService) SetJobHistory(val bool) *reputationListService {
	s.jobHistory = &val
	return s
}

func (s *reputationListService) SetProjectState(val bool) *reputationListService {
	s.projectStats = &val
	return s
}

func (s *reputationListService) SetRehireRates(val bool) *reputationListService {
	s.rehireRates = &val
	return s
}

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

// SetUsers sets the filter user IDs for the reputation list service.
func (s *reputationListService) SetUsers(users []int64) *reputationListService {
	s.users = users
	return s
}

// SetJobs limits users to a set of job IDs as a filter.
func (s *reputationListService) SetJobs(jobs []int32) *reputationListService {
	s.jobs = jobs
	return s
}

// SetRole sets the type of reputation filter (e.g., freelancer or employer).
func (s *reputationListService) SetRole(val RoleType) *reputationListService {
	s.role = &val
	return s
}

// SetJobHistory sets whether to return job history projection for each user.
func (s *reputationListService) SetJobHistory(val bool) *reputationListService {
	s.jobHistory = &val
	return s
}

// SetProjectState sets whether to return project stats projection for each user.
func (s *reputationListService) SetProjectState(val bool) *reputationListService {
	s.projectStats = &val
	return s
}

// SetRehireRates sets whether to return rehire rates projection for a single freelancer.
func (s *reputationListService) SetRehireRates(val bool) *reputationListService {
	s.rehireRates = &val
	return s
}

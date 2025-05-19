package freelancer

import (
	"context"
	"encoding/json"
	"net/http"
)

// Gets the reputations for a list of users
type ListUsersReputationsService struct {
	clinet       *Client
	users        []int64
	jobs         []int32
	role         *RoleType
	jobHistory   *bool
	projectStats *bool
	rehireRates  *bool
}

type ListUsersReputationsResponse struct {
	Status    string                `json:"status"`
	RequsetID string                `json:"request_id"`
	Result    map[string]Reputation `json:"result"`
}

func (s *ListUsersReputationsService) DO(ctx context.Context) (*ListUsersReputationsResponse, error) {
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

	data, err := s.clinet.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	resp := &ListUsersReputationsResponse{}
	if err := json.Unmarshal(data, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *ListUsersReputationsService) SetUsers(users []int64) *ListUsersReputationsService {
	s.users = users
	return s
}

func (s *ListUsersReputationsService) SetJobs(jobs []int32) *ListUsersReputationsService {
	s.jobs = jobs
	return s
}

func (s *ListUsersReputationsService) SetRole(val RoleType) *ListUsersReputationsService {
	s.role = &val
	return s
}

func (s *ListUsersReputationsService) SetJobHistory(val bool) *ListUsersReputationsService {
	s.jobHistory = &val
	return s
}

func (s *ListUsersReputationsService) SetProjectState(val bool) *ListUsersReputationsService {
	s.projectStats = &val
	return s
}

func (s *ListUsersReputationsService) SetRehireRates(val bool) *ListUsersReputationsService {
	s.rehireRates = &val
	return s
}

package freelancer

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

// Creats a new project
// Note the service implemented for creating a fixed an hourly projects
// HireMe and Local not implemented yet
type CreateProjectService struct {
	client *Client
}

type CreateProjectRequestBody struct {
	// required
	Title string `json:"title"`
	// required
	Description string `json:"description"`
	// required with minimum amount budget
	Budget Budget `json:"budget"`
	// required
	Jobs              []int32           `json:"jobs"`
	Type              ProjectType       `json:"type"`
	HourlyProjectInfo HourlyProjectInfo `json:"hourly_project_info"`
}

type HourlyProjectInfo struct {
	Commitment Commitment `json:"commitment"`
}

type Commitment struct {
	Hours    int          `json:"hours"`
	Interval IntervalType `json:"interval"`
}

type HireMeInitialBid struct {
	BidderId int64 `json:"bidder_id"`
	Amount   int   `json:"amount"`
	Period   int   `json:"period"`
}

func (s *CreateProjectService) Do(ctx context.Context, b CreateProjectRequestBody) (*BaseResponse, error) {
	m, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}

	r := &request{
		method:   http.MethodPost,
		endpoint: string(PROJECTS_PROJECTS),
		body:     bytes.NewBuffer(m),
	}

	data, err := s.client.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}

	resp := &BaseResponse{}
	if err := json.Unmarshal(data, resp); err != nil {
		return nil, err
	}
	return resp, err
}

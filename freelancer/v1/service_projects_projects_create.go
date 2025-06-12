package freelancer

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

type projectsCreateService struct {
	client *Client
}

// Title, Description, Budget, Jobs are required
type CreateProjectBody struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Budget      struct {
		Minimum    float64  `json:"minimum"`
		Maximum    *float64 `json:"maximum,omitempty"`
		CurrencyID *int64   `json:"currency_id,omitempty"`
	} `json:"budget"`
	Jobs              []int64      `json:"jobs"`
	Type              *ProjectType `json:"type,omitempty"`
	HourlyProjectInfo *struct {
		Commitment struct {
			Hours    int          `json:"hours"`
			Interval IntervalType `json:"interval"`
		} `json:"commitment"`
	} `json:"hourly_project_info,omitempty"`
	HirMe            *bool `json:"hire_me,omitempty"`
	HiremeInitialBid *struct {
		BidderID int64   `json:"bidder_id"`
		Amount   float64 `json:"amount"`
		Period   int64   `json:"period"`
	} `json:"hireme_initial_bid,omitempty"`
}

func (s *projectsCreateService) Do(ctx context.Context, b CreateProjectBody) (*BaseResponse, error) {
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

package freelancer

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// Returns bids for a single project.
type bidTimeTrackingGetService struct {
	client                *Client
	fromTime              *int64
	toTime                *int64
	dailyAggregateDetails *bool
	invoiced              *bool
}

func (s *bidTimeTrackingGetService) Do(ctx context.Context, bidID int) (*BaseResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("%s/%d/time_tracking_sessions", string(PROJECTS_BIDS), bidID),
	}

	if s.fromTime != nil {
		r.addParam("from_time", *s.fromTime)
	}
	if s.toTime != nil {
		r.addParam("to_time", *s.toTime)
	}
	if s.dailyAggregateDetails != nil {
		r.addParam("daily_aggregate_details", *s.dailyAggregateDetails)
	}
	if s.invoiced != nil {
		r.addParam("invoiced", *s.invoiced)
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

func (s *bidTimeTrackingGetService) SetFromTime(val int64) *bidTimeTrackingGetService {
	s.fromTime = &val
	return s
}

func (s *bidTimeTrackingGetService) SetToTime(val int64) *bidTimeTrackingGetService {
	s.toTime = &val
	return s
}

func (s *bidTimeTrackingGetService) SetDailyAggreateDetails(val bool) *bidTimeTrackingGetService {
	s.dailyAggregateDetails = &val
	return s
}

func (s *bidTimeTrackingGetService) SetInvoiced(val bool) *bidTimeTrackingGetService {
	s.invoiced = &val
	return s
}

package freelancer

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// Returns bids for a single project.
type getBidTimeTrackingService struct {
	client                *Client
	fromTime              *int64
	toTime                *int64
	dailyAggregateDetails *bool
	invoiced              *bool
}

func (s *getBidTimeTrackingService) Do(ctx context.Context, bidID int) (*BaseResponse, error) {
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

func (s *getBidTimeTrackingService) SetFromTime(val int64) *getBidTimeTrackingService {
	s.fromTime = &val
	return s
}

func (s *getBidTimeTrackingService) SetToTime(val int64) *getBidTimeTrackingService {
	s.toTime = &val
	return s
}

func (s *getBidTimeTrackingService) SetDailyAggreateDetails(val bool) *getBidTimeTrackingService {
	s.dailyAggregateDetails = &val
	return s
}

func (s *getBidTimeTrackingService) SetInvoiced(val bool) *getBidTimeTrackingService {
	s.invoiced = &val
	return s
}

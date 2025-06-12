package freelancer

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type milestonesActionsService struct {
	client *Client
}

type ActionMilestoneBody struct {
	Action      ActionBid             `json:"action"`
	Amount      int                   `json:"amount"`
	Reason      MilestoneActionReason `json:"reason"`
	ReasonText  string                `json:"reason_text"`
	OtherReason string                `json:"other_reason"`
}

func (s *milestonesActionsService) Do(ctx context.Context, milestoneID int, b ActionMilestoneBody) (*BaseResponse, error) {
	m, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodPut,
		endpoint: fmt.Sprintf("%s/%d", string(PROJECTS_MILESTONES), milestoneID),
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

	return resp, nil

}

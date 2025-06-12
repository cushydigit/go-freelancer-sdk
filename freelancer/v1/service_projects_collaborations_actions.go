package freelancer

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type collaborationsActionsService struct {
	client *Client
}
type ActionProjectCollaborationsBody struct {
	// required
	Action ProjectCollaborationAction `json:"action"`
	// required
	Permissions struct {
		Chat     bool `json:"CHAT"`
		BidAward bool `json:"BID_AWARD"`
	}
}

func (s *collaborationsActionsService) Do(ctx context.Context, id int64, collaborationID int, b ActionProjectCollaborationsBody) (*BaseResponse, error) {
	m, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodPut,
		endpoint: fmt.Sprintf("%s/%s/collaborations/%d", string(PROJECTS_PROJECTS), strconv.FormatInt(id, 10), collaborationID),
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

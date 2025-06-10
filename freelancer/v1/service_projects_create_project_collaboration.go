package freelancer

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type CreateProjectCollaborationService struct {
	client *Client
}
type CreateProjectCollaborationRequestBody struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Title    string `json:"title"`
	// required
	Permissions struct {
		Chat     bool `json:"CHAT"`
		BidAward bool `jsong:"BID_AWARD"`
	}
}

func (s *CreateProjectCollaborationService) Do(ctx context.Context, id int64, b CreateProjectCollaborationRequestBody) (*BaseResponse, error) {
	m, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodPost,
		endpoint: fmt.Sprintf("%s/%s/collaborations/", string(PROJECTS_PROJECTS), strconv.FormatInt(id, 10)),
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

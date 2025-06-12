package freelancer

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type collaborationsCreateService struct {
	client *Client
}
type CreateProjectCollaborationsBody struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Title    string `json:"title"`
	// required
	Permissions struct {
		Chat     bool `json:"CHAT"`
		BidAward bool `json:"BID_AWARD"`
	}
}

func (s *collaborationsCreateService) Do(ctx context.Context, id int64, b CreateProjectCollaborationsBody) (*BaseResponse, error) {
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

package freelancer

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

type reviewsCreateService struct {
	client *Client
}

type CreateReviewBody struct {
	ProjectID  int64      `json:"project_id"`
	ToUserID   int64      `json:"to_user_id"`
	FromUserID int64      `json:"from_user_id"`
	ReviewType ReviewType `json:"review_type"`
	Comment    string     `json:"comment"`
	Role       RoleType   `json:"role"`
}

func (s *reviewsCreateService) Do(ctx context.Context, b CreateReviewBody) (*BaseResponse, error) {
	m, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodPost,
		endpoint: string(PROJECTS_REVIEWS),
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

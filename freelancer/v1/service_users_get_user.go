package freelancer

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// Returns information about a specific user
type GetUserService struct {
	client *Client
}

type GetUserResponse struct {
	Status string `json:"status"`
	Result User   `json:"result"`
}

func (s *GetUserService) Do(ctx context.Context, userID int64) (*GetUserResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("%s/%s", string(USERS_USERS), strconv.FormatInt(userID, 10)),
	}

	data, err := s.client.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}

	res := &GetUserResponse{}
	if err := json.Unmarshal(data, res); err != nil {
		return nil, err
	}

	return res, nil
}

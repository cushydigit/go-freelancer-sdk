package freelancer

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type usersGetByIDService struct {
	client *Client
}

func (s *usersGetByIDService) Do(ctx context.Context, userID int64) (*BaseResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("%s/%s", string(USERS_USERS), strconv.FormatInt(userID, 10)),
	}

	data, err := s.client.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}

	res := &BaseResponse{}
	if err := json.Unmarshal(data, res); err != nil {
		return nil, err
	}

	return res, nil
}

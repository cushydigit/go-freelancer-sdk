package freelancer

import (
	"context"
	"encoding/json"
	"net/http"
)

type selfDevicesListService struct {
	client *Client
}

func (s *selfDevicesListService) Do(ctx context.Context) (*BaseResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: string(USERS_SELF_DEVICES),
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

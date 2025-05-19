package freelancer

import (
	"context"
	"encoding/json"
	"net/http"
)

// Return a list of current user's recent logged in devices
type ListSelfLoginDevicesService struct {
	client *Client
}

type ListSelfLoginDevicesResponse struct {
	Status    string        `json:"status"`
	RequestID string        `json:"requset_id"`
	Result    DevicesResult `json:"result"`
}

type DevicesResult struct {
	Devices []Device `json:"devices"`
}

type Device struct {
	UserAgent string `json:"user_agent"`
	Platform  string `json:"platform"`
	City      string `json:"city"`
	Country   string `json:"country"`
	LastLogin int64  `json:"last_login"`
}

func (s *ListSelfLoginDevicesService) DO(ctx context.Context) (*ListSelfLoginDevicesResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: string(USERS_SELF_DEVICES),
	}

	data, err := s.client.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := &ListSelfLoginDevicesResponse{}
	if err := json.Unmarshal(data, res); err != nil {
		return nil, err
	}
	return res, nil
}

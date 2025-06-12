package freelancer

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type servicesOrderService struct {
	client *Client
}

func (s *servicesOrderService) Do(ctx context.Context, id int, serviceType ServiceType) (*BaseResponse, error) {
	r := &request{
		method:   http.MethodPut,
		endpoint: fmt.Sprintf("%s/%s/%d/order", string(PROJECTS_SERVICES), string(serviceType), id),
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

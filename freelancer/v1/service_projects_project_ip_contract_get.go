package freelancer

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type projectIPContractInfoGetService struct {
	client *Client
}

func (s *projectIPContractInfoGetService) Do(ctx context.Context, id int64) (*BaseResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("%s/%s/ip_contracts", string(PROJECTS_PROJECTS), strconv.FormatInt(id, 10)),
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

package freelancer

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// Returns information for posting bids on project
type GetProjectBidInformationService struct {
	client    *Client
	projectID *int64
}

func (s *GetProjectBidInformationService) Do(ctx context.Context, id int64) (*BaseResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("%s/%s/bidinfo", string(PROJECTS_PROJECTS), strconv.FormatInt(id, 10)),
	}
	if s.projectID != nil {
		r.addParam("project_id", *s.projectID)
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

func (s *GetProjectBidInformationService) SetProjectID(val int64) *GetProjectBidInformationService {
	s.projectID = &val
	return s
}

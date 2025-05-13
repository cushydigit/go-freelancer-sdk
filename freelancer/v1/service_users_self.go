package freelancer

import (
	"context"
	"encoding/json"
	"net/http"
)

type ListSelfDevicesService struct {
	client *Client
}

type ListSelfDevicesResponse struct {
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

func (s *ListSelfDevicesService) Do(ctx context.Context) (*ListSelfDevicesResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: string(USERS_SELF_DEVICES),
	}
	data, err := s.client.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := &ListSelfDevicesResponse{}
	if err := json.Unmarshal(data, res); err != nil {
		return nil, err
	}
	return res, nil
}

type GetSelfInfoService struct {
	client                       *Client
	avatar                       bool
	countryDetails               bool
	profileDescription           bool
	displayInfo                  bool
	jobs                         bool
	balanceDetails               bool
	qualificationDetails         bool
	membershipDetails            bool
	financialDetails             bool
	locationDetails              bool
	portfolioDetails             bool
	preferredDetails             bool
	badgeDetails                 bool
	status                       bool
	reputation                   bool
	employerReputation           bool
	reputationExtra              bool
	employerReputationExtra      bool
	coverImage                   bool
	pastCoverImages              bool
	mobileTracking               bool
	bidQualityDetails            bool
	depositMethods               bool
	userRecommendations          bool
	marketingMobileNumber        bool
	sanctionDetials              bool
	limitedAccount               bool
	completedUserRelavantJobCont bool
	equipmentGroupDetails        bool
	jonRanks                     bool
	jobSeoDetails                bool
	risingStar                   bool
	shareholderDetails           bool
	staffDetils                  bool
	limit                        int
	offset                       int
	compact                      bool
}

type GetSelfInfoResponse struct {
	Status string `json:"status"`
	Result User   `json:"result"`
}

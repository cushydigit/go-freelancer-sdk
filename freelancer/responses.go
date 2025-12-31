package freelancer

import "encoding/json"

type base struct {
	Status    string `json:"status"`
	RequestID string `json:"request_id"`
}
type RawResponse struct {
	base
	Result json.RawMessage `json:"result"`
}

// Common Responses
type ListCountriesResponse struct {
	base
	Result struct {
		Countries []*Country `json:"countries"`
	} `json:"result"`
}

type ListTimezonesResponse struct {
	base
	Result struct {
		Timezones []*Timezone `json:"timezones"`
	} `json:"result"`
}

type ListCategoriesResponse struct {
	base
	Result struct {
		Categories []*Category `json:"categories"`
		Jobs       []*Job      `json:"jobs,omitempty"`
	} `json:"result"`
}

type ListProjectsResponse struct {
	base
	Result struct {
		Projects   []*Project `json:"projects"`
		TotalCount int        `json:"total_count"`
	} `json:"result"`
}

type ListUsersResponse struct {
	base
	Result struct {
		Users map[string]*User `json:"users"`
	} `json:"result"`
}

type GetProjectResponse struct {
	base
	Result *Project `json:"result"`
}

type CreateProjectResponse struct {
	base
	Result *Project `json:"result"`
}
type GetUserResponse struct {
	base
	Result *User `json:"result"`
}

type ListUsersReputationsResponse struct {
	base
	Result map[string]*Reputation `json:"result"`
}

type SearchFreelancersResponse struct {
	base
	Result struct {
		Users      []*User `json:"users"`
		TotalCount int     `json:"total_count"`
	} `json:"result"`
}

type GetSelfInfoResponse struct {
	base
	Result *User `json:"result"`
}

type ListBudgetsResponse struct {
	base
	Result struct {
		Budgets []*Budget `json:"budgets"`
	} `json:"result"`
}

type ListCurrenciesResponse struct {
	base
	Result struct {
		Currencies []*Currency `json:"currencies"`
	} `json:"result"`
}
type ListUsersPortfoliosResponse struct {
	base
	Result json.RawMessage `json:"result"`
}

type ListSelfLoginDevicesResponse struct {
	base
	Result struct {
		Devices []*Device `json:"devices"`
	} `json:"result"`
}

type ListBidEditRequestsResponse struct {
	base
	Result struct {
		BidEditRequests []*BidEditRequest `json:"bid_edit_requests"`
	} `json:"result"`
}

type CreateBidEditRequestResponse struct {
	base
	Result *BidEditRequest `json:"result"`
}

type ActionBidEditRequestResponse struct {
	base
	Result *BidEditRequest `json:"result"`
}

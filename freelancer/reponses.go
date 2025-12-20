package freelancer

type BaseResponse struct {
	Status    string `json:"status"`
	RequestID string `json:"request_id"`
}

// Common Responses
type ListCountriesResponse struct {
	BaseResponse
	Result struct {
		Countries []Country `json:"countries"`
	} `json:"result"`
}

type ListTimezonesResponse struct {
	BaseResponse
	Result struct {
		Timezones []Timezone `json:"timezones"`
	} `json:"result"`
}

type ListCategoriesResponse struct {
	BaseResponse
	Result struct {
		Categories []Category `json:"categories"`
		Jobs       []Job      `json:"jobs,omitempty"`
	} `json:"result"`
}

type ListProjectsResponse struct {
	BaseResponse
	Result struct {
		Projects   []Project `json:"projects"`
		TotalCount int       `json:"total_count"`
	} `json:"result"`
}

type ListUsersResponse struct {
	BaseResponse
	Result struct {
		Users map[string]User `json:"users"`
	} `json:"result"`
}

type GetProjectResponse struct {
	BaseResponse
	Result Project `json:"result"`
}

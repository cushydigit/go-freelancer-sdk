package responseHelper

type ListCountriesResponse struct {
	Status    string          `json:"status"`
	RequestID string          `json:"request_id,omitempty"` // Optional
	Result    CountriesResult `json:"result"`
}

type CountriesResult struct {
	Countries []Country `json:"countries"`
}

type ListTimezonesResponse struct {
	Status    string          `json:"status"`
	RequestID string          `json:"request_id,omitempty"` // Optional
	Result    TimezonesResult `json:"result"`
}
type TimezonesResult struct {
	Timezones []Timezone `json:"timezones"` // Optional, "Decimal"
}

type Timezone struct {
	ID       int     `json:"id"`                 // "number" represented as float64
	Country  string  `json:"country,omitempty"`  // Optional
	Timezone string  `json:"timezone,omitempty"` // Optional
	Offset   float32 `json:"offset,omitempty"`
}

type ListCategoriesResponse struct {
	Status    string           `json:"status"`
	RequestID string           `json:"request_id,omitempty"` // Optional
	Result    CategoriesResult `json:"result"`
}

type CategoriesResult struct {
	// Jobs       []Job      `json:"jobs,omitempty"`
	Categories []Category `json:"categories"`
}

type GetProjectResponse struct {
	Status    string  `json:"status"`
	RequestID string  `json:"request_id"`
	Result    Project `json:"result"`
}

type ListProjectsResponse struct {
	Status    string         `json:"status"`
	RequestID string         `json:"request_id"`
	Result    ProjectsResult `json:"result"`
}

type ProjectsResult struct {
	Projects   []Project `json:"projects"`
	TotalCount int       `json:"total_count"`
}

type SearchActiveProjectsResponse struct {
	Status    string               `json:"status"`
	RequestID string               `json:"request_id"`
	Result    ActiveProjectsResult `json:"result"`
}

type ActiveProjectsResult struct {
	Projects   []Project `json:"projects"`
	TotalCount int       `json:"total_count"`
}
type GetUserResponse struct {
	Status string `json:"status"`
	Result User   `json:"result"`
}

type ListUsersResponse struct {
	Status    string      `json:"status"`
	RequestID string      `json:"request_id"`
	Result    UsersResult `json:"result"`
}

type UsersResult struct {
	Users map[string]User `json:"users"`
}

// Gets the reputations for a list of users

type ListUsersReputationsResponse struct {
	Status    string                `json:"status"`
	RequestID string                `json:"request_id"`
	Result    map[string]Reputation `json:"result"`
}

type SearchFreelancersResponse struct {
	Status    string           `json:"status"`
	RequestID string           `json:"request_id"`
	Result    FreelancerResult `json:"result"`
}

type FreelancerResult struct {
	Users      []User `json:"users"`
	TotalCount int    `json:"total_count"`
}

type GetSelfInfoResponse struct {
	Status string `json:"status"`
	Result User   `json:"result"`
}

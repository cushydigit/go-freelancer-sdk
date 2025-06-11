package freelancer

type Endpoint string

const (
	COMMON_COUNTRIES Endpoint = "/common/0.1/countries"
	COMMON_TIMEZONES Endpoint = "/common/0.1/timezones"

	PROJECTS_PROJECTS           Endpoint = "/projects/0.1/projects"
	PROJECTS_SELF               Endpoint = "/projects/0.1/self"
	PROJECTS_BUDGETS            Endpoint = "/projects/0.1/budgets"
	PROJECTS_CATEGORIES         Endpoint = "/projects/0.1/categories"
	PROJECTS_CURRENCIES         Endpoint = "/projects/0.1/currencies"
	PROJECTS_BIDS               Endpoint = "/projects/0.1/bids"
	PROJECTS_BIDS_EDIT_REQUESTS Endpoint = "/projects/0.1/bids/bid_edit_requests"
	PROJECTS_REVIEWS            Endpoint = "/projects/0.1/reviews"
	PROJECTS_MILESTONRE         Endpoint = "/projects/0.1/milestonte"
	PROJECTS_HOURLY_CONTRACTS   Endpoint = "/projects/0.1/hourly_contracts"
	PROJECTS_PROJECTS_ACTIVE    Endpoint = "/projects/0.1/projects/active"
	PROJECTS_PROJECTS_FEES      Endpoint = "/projects/0.1/projects/fees"
	PROJECTS_JOBS_SEARCH        Endpoint = "/projects/0.1/jobs/search"
	PROJECTS_BIDS_FEES          Endpoint = "/projects/0.1/bids/fees"
	PROJECTS_MILESTONE_REQUESTS Endpoint = "/projects/0.1/milestonte_requests/"
	PROJECTS_COLLABORATIONS     Endpoint = "/projects/0.1/projects/collaborations"
	PROJECTS_SERVICES           Endpoint = "/projects/0.1/services"

	USERS_USERS       Endpoint = "/users/0.1/users"
	USERS_FREELANCERS Endpoint = "/users/0.1/users/directory"
	USERS_REPUTATIONS Endpoint = "/users/0.1/reputations"
	USERS_PORTFOLIOS  Endpoint = "/users/0.1/portfolios"
	USERS_PROFILES    Endpoint = "/users/0.1/profiles"
	USERS_ENTERPRISES Endpoint = "/users/0.1/enterprises"
	USERS_VIOLATION   Endpoint = "/users/0.1/violation_reports"
	USERS_POOLS       Endpoint = "/users/0.1/pools"

	USERS_SELF         Endpoint = "/users/0.1/self"
	USERS_SELF_DEVICES Endpoint = "/users/0.1/self/devices"
	USERS_SELF_JOBS    Endpoint = "/users/0.1/self/jobs"
)

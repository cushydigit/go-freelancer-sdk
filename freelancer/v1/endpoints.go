package freelancer

type Endpoint string

const (
	COMMON_COUNTRIES Endpoint = "/common/0.1/countries"
	COMMON_TIMEZONES Endpoint = "/common/0.1/timezones"

	// POST: create fixed, hourly, local projects
	PROJECTS_PROJECTS   Endpoint = "/projects/0.1/projects"
	PROJECTS_BUDGETS    Endpoint = "/projects/0.1/budgets"
	PROJECTS_CATEGORIES Endpoint = "/projects/0.1/categories"
	PROJECTS_CURRENCIES Endpoint = "/projects/0.1/currencies"
	// POST: create a bid on a projects
	// GET: list bids
	PROJECTS_BIDS Endpoint = "/projects/0.1/bids"
	// GET: search porjects
	PROJECTS_PROJECT_ACTIVE Endpoint = "/projects/0.1/projects/active"
	// GET: to get retrieve th Job ID for a skill
	PROJECTS_JOBS_SEARCH Endpoint = "/projects/0.1/jobs/search"
	// POST: posting a review on completed project
	PROJECTS_REVIEWS = "/projects/0.1/reviews"
	// POST: create a milestone on fixed project
	// \\tPUT: release a milestone {id}
	PROJECTS_MILESTONRE = "/projects/0.1/milestonte"
	// PUT: accespt a milestotne
	// POST: requset a milestone
	PROJECTS_MILESTONE_REQUESTS = "projects/0.1/milestonte_requests/{id}"

	USERS_USERS       Endpoint = "/users/0.1/users"
	USERS_FREELANCERS Endpoint = "/users/0.1/users/directory"
	USERS_REPUTATIONS Endpoint = "/users/0.1/reputations"

	USERS_SELF         Endpoint = "/users/0.1/self"
	USERS_SELF_DEVICES Endpoint = "/users/0.1/self/devices"
	// POST, PUT, DELETE for add, replace or delete jobs
	USERS_SELF_JOBS = "/users/0.1/self/jobs"
)

package freelancer

type Endpoint string

const (
	COMMON_COUNTRIES Endpoint = "/common/0.1/countries"
	COMMON_TIMEZONES Endpoint = "/common/0.1/timezones"

	PROJECT_PROJECT        Endpoint = "/projects/0.1/projects"
	PROJECT_BUDGETS        Endpoint = "/projects/0.1/budgets"
	PROJECT_CATEGORIES     Endpoint = "/projects/0.1/categories"
	PROJECT_CURRENCIES     Endpoint = "/projects/0.1/currencies"
	PROJECT_PROJECT_ACTIVE Endpoint = "/projects/0.1/projects/active"

	USERS_USERS Endpoint = "/users/0.1/users"
)

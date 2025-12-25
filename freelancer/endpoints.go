package freelancer

type Endpoint string

const (
	CommonCountries Endpoint = "/common/0.1/countries"
	CommonTimezones Endpoint = "/common/0.1/timezones"

	ProjectsProjects            Endpoint = "/projects/0.1/projects"
	ProjectsSelf                Endpoint = "/projects/0.1/self"
	ProjectsBudgets             Endpoint = "/projects/0.1/budgets"
	ProjectsCategories          Endpoint = "/projects/0.1/categories"
	ProjectsCurrencies          Endpoint = "/projects/0.1/currencies"
	ProjectsBids                Endpoint = "/projects/0.1/bids"
	ProjectsBidsBidEditRequests Endpoint = "/projects/0.1/bids/bid_edit_requests"
	ProjectsBidRatings          Endpoint = "/projects/0.1/bid_ratings"
	ProjectsReviews             Endpoint = "/projects/0.1/reviews"
	ProjectsMilestones          Endpoint = "/projects/0.1/milestones"
	ProjectsMilestoneRequests   Endpoint = "/projects/0.1/milestone_requests/"
	ProjectsHourlyContract      Endpoint = "/projects/0.1/hourly_contracts"
	ProjectsProjectsActive      Endpoint = "/projects/0.1/projects/active"
	ProjectsProjectsAll         Endpoint = "/projects/0.1/projects/all"
	ProjectsProjectsFees        Endpoint = "/projects/0.1/projects/fees"
	ProjectsJobs                Endpoint = "/projects/0.1/jobs"
	ProjectsJobBundles          Endpoint = "/projects/0.1/job_bundles"
	ProjectsJobBundleCategories Endpoint = "/projects/0.1/job_bundle_categories"
	ProjectsJobsSearch          Endpoint = "/projects/0.1/jobs/search"
	ProjectsBidsFees            Endpoint = "/projects/0.1/bids/fees"
	ProjectsCollaborations      Endpoint = "/projects/0.1/projects/collaborations"
	ProjectsServices            Endpoint = "/projects/0.1/services"
	ProjectsServicesActive      Endpoint = "/projects/0.1/services/active"
	ProjectsExpertGuarantees    Endpoint = "/projects/0.1/expert_guarantees"

	UsersUsers            Endpoint = "/users/0.1/users"
	UsersFreelancers      Endpoint = "/users/0.1/users/directory"
	UsersReputations      Endpoint = "/users/0.1/reputations"
	UsersPortfolios       Endpoint = "/users/0.1/portfolios"
	UsersProfiles         Endpoint = "/users/0.1/profiles"
	UsersEnterprises      Endpoint = "/users/0.1/enterprises"
	UsersViolationReports Endpoint = "/users/0.1/violation_reports"
	UsersPools            Endpoint = "/users/0.1/pools"

	UsersSelf        Endpoint = "/users/0.1/self"
	UsersSelfDevices Endpoint = "/users/0.1/self/devices"
	UsersSelfJobs    Endpoint = "/users/0.1/self/jobs"
)

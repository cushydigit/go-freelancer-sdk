package endpoints

import "fmt"

const (
	Users            Endpoint = "/users/0.1/users"
	Freelancers      Endpoint = "/users/0.1/users/directory"
	Reputations      Endpoint = "/users/0.1/reputations"
	Portfolios       Endpoint = "/users/0.1/portfolios"
	Profiles         Endpoint = "/users/0.1/profiles"
	Enterprises      Endpoint = "/users/0.1/enterprises"
	ViolationReports Endpoint = "/users/0.1/violation_reports"
	Pools            Endpoint = "/users/0.1/pools"
	Self             Endpoint = "/users/0.1/self"
	Devices          Endpoint = "/users/0.1/self/devices"
	SelfJobs         Endpoint = "/users/0.1/self/jobs"
)

// it maps to the `/users/0.1/users/{user_id}`
func User(id int64) Endpoint {
	return Endpoint(fmt.Sprintf("%s/%d", Users, id))
}

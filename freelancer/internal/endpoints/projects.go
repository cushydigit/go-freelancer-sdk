package endpoints

import (
	"fmt"
)

const (
	Projects               Endpoint = "/projects/0.1/projects"
	ProjectsSelf           Endpoint = "/projects/0.1/self"
	ProjectsActive         Endpoint = "/projects/0.1/projects/active"
	ProjectsAll            Endpoint = "/projects/0.1/projects/all"
	ProjectsFees           Endpoint = "/projects/0.1/projects/fees"
	ProjectsCollaborations Endpoint = "/projects/0.1/projects/collaborations"
	Budgets                Endpoint = "/projects/0.1/budgets"
	Categories             Endpoint = "/projects/0.1/categories"
	Currencies             Endpoint = "/projects/0.1/currencies"
	Bids                   Endpoint = "/projects/0.1/bids"
	BidsEditRequests       Endpoint = "/projects/0.1/bids/bid_edit_requests"
	BidRatings             Endpoint = "/projects/0.1/bid_ratings"
	Reviews                Endpoint = "/projects/0.1/reviews"
	Milestones             Endpoint = "/projects/0.1/milestones"
	MilestoneRequests      Endpoint = "/projects/0.1/milestone_requests"
	HourlyContractInfo     Endpoint = "/projects/0.1/hourly_contract_info"
	Jobs                   Endpoint = "/projects/0.1/jobs"
	JobBundles             Endpoint = "/projects/0.1/job_bundles"
	JobBundleCategories    Endpoint = "/projects/0.1/job_bundle_categories"
	JobsSearch             Endpoint = "/projects/0.1/jobs/search"
	BidsFees               Endpoint = "/projects/0.1/bids/fees"
	Services               Endpoint = "/projects/0.1/services"
	ServicesActive         Endpoint = "/projects/0.1/services/active"
	ExpertGuarantees       Endpoint = "/projects/0.1/expert_guarantees"
)

// it maps to the `/projects/0.1/projects/{project_id}`
func Project(id int64) Endpoint {
	return Endpoint(fmt.Sprintf("%s/%d", Projects, id))
}

// it maps to the `/projects/0.1/projects/{project_id}/invite`
func ProjectInvite(id int64) Endpoint {
	return Endpoint(fmt.Sprintf("%s/%d/invite", Projects, id))
}

// it maps to the `/projects/0.1/projects/{project_id}/bids`
func ProjectBids(id int64) Endpoint {
	return Endpoint(fmt.Sprintf("%s/%d/bids", Projects, id))
}

// it maps to the `/projects/0.1/projects/{project_id}/bids_info`
func ProjectBidsInfo(id int64) Endpoint {
	return Endpoint(fmt.Sprintf("%s/%d/bids_info", Projects, id))
}

// it maps to the `/projects/0.1/projects/{project_id}/milestones`
func ProjectMilestones(id int64) Endpoint {
	return Endpoint(fmt.Sprintf("%s/%d/milestones", Projects, id))
}

// it maps to the `/projects/0.1/projects/{project_id}/milestone_requests`
func ProjectMilestoneRequests(id int64) Endpoint {
	return Endpoint(fmt.Sprintf("%s/%d/milestone_requests", Projects, id))
}

// it maps to the `/projects/0.1/projects/{project_id}/ip_contract_info`
func ProjectIPContractInfo(id int64) Endpoint {
	return Endpoint(fmt.Sprintf("%s/%d/ip_contract_info", Projects, id))
}

// it maps to the `/projects/0.1/projects/{project_id}/collaborations`
func ProjectCollaborations(id int64) Endpoint {
	return Endpoint(fmt.Sprintf("%s/%d/collaborations", Projects, id))
}

// it maps to the `/projects/0.1/projects/{projects_id}/collaborations/{collaboration_id}/actions`
func ProjectCollaborationsActions(id, collaborationID int64) Endpoint {
	return Endpoint(fmt.Sprintf("%s/%d/collaborations/%d/actions", Projects, id, collaborationID))
}

// it maps to the `/projects/0.1/services/{service_type}/{service_id}/order`
func ServicesOrder(serviceType string, serviceID int64) Endpoint {
	return Endpoint(fmt.Sprintf("%s/%s/%d/order", Services, serviceType, serviceID))
}

// it maps to the `/projects/0.1/bids/{bid_id}`
func Bid(bidID int64) Endpoint {
	return Endpoint(fmt.Sprintf("%s/%d", Bids, bidID))
}

// it maps to the `/projects/0.1/bids/{bid_id}/time_tracking`
func BidTimeTracking(bidID int64) Endpoint {
	return Endpoint(fmt.Sprintf("%s/%d/time_tracking", Bids, bidID))
}

// it maps to the `/projects/0.1/bids/{bid_id}/edit_requests`
func BidEditRequests(bidID int64) Endpoint {
	return Endpoint(fmt.Sprintf("%s/%d/edit_requests", Bids, bidID))
}

// it maps to the `/projects/0.1/bids/{bid_id}/edit_requests/{bid_edit_request_id}`
func BidEditRequest(bidID int64, bidEditRequestID int64) Endpoint {
	return Endpoint(fmt.Sprintf("%s/%d/edit_requests/%d", Bids, bidID, bidEditRequestID))
}

// it maps to the `/projects/0.1/bids/{bid_id}/bid_ratings`
func BidsRatings(bidID int64) Endpoint {
	return Endpoint(fmt.Sprintf("%s/%d/bid_ratings", Bids, bidID))
}

// it maps to the `/projects/0.1/bids/{bid_id}/bid_ratings/{bid_rating_id}`
func BidRating(bidID int64, bidRatingID int64) Endpoint {
	return Endpoint(fmt.Sprintf("%s/%d/bid_ratings/%d", Bids, bidID, bidRatingID))
}

// it maps to the `/projects/0.1/milestones/{milestone_id}`
func Milestone(milestoneID int64) Endpoint {
	return Endpoint(fmt.Sprintf("%s/%d", Milestones, milestoneID))
}

// it maps to the `/projects/0.1/milestone_requests/{milestone_request_id}`
func MilestoneRequest(milestoneRequestID int64) Endpoint {
	return Endpoint(fmt.Sprintf("%s/%d", MilestoneRequests, milestoneRequestID))
}

// it maps to the `/projects/0.1/reviews/{review_id}`
func Review(reviewID int64) Endpoint {
	return Endpoint(fmt.Sprintf("%s/%d", Reviews, reviewID))
}

// it maps to the `/projects/0.1/expert_guarantees/{expert_guarantees_id}`
func ExpertGuarantee(expertGuaranteesID int64) Endpoint {
	return Endpoint(fmt.Sprintf("%s/%d", ExpertGuarantees, expertGuaranteesID))
}

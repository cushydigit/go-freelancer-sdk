package utils

// TODO: check this service functionality so the projectActionBody should have better fields
// check : https://developers.freelancer.com/docs/projects/projects#projects-put
// Performs an action on a project.
type ProjectActionBody struct {
	ProjectID int64         `json:"project_id"`
	Action    ProjectAction `json:"action"`
}

type ActionProjectCollaborationsBody struct {
	// required
	Action ProjectCollaborationAction `json:"action"`
	// required
	Permissions struct {
		Chat     bool `json:"CHAT"`
		BidAward bool `json:"BID_AWARD"`
	}
}

type ActionBidBody struct {
	Action BidAction `json:"action"`
}

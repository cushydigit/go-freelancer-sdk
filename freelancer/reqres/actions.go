package reqres

// check : https://developers.freelancer.com/docs/projects/projects#projects-put
// Performs an action on a project.
type ActionProjectBody struct {
	ProjectID int64         `json:"project_id"`
	Action    ProjectAction `json:"action"`
}

// action an permissions are required
type ActionCollaborationBody struct {
	Action      CollaborationAction `json:"action"`
	Permissions Permissions         `json:"permissions"`
}

type Permissions struct {
	Chat     bool `json:"CHAT"`
	BidAward bool `json:"BID_AWARD"`
}

type ActionBidBody struct {
	Action BidAction `json:"action"`
}

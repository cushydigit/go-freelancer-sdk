package reqres

type ActionProjectBody struct {
	Action ProjectAction `json:"action"`
}

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

type ActionReviewBody struct {
	Action     ReviewAction `json:"action"`
	ReviewType ReviewType   `json:"review_type"`
}

type ActionExpertGuaranteesBody struct {
	Action ExpertGuaranteesAction `json:"action"`
}

type ActionMilestoneBody struct {
	Action      MilestoneAction       `json:"action"`
	Amount      int                   `json:"amount"`
	Reason      MilestoneActionReason `json:"reason"`
	ReasonText  string                `json:"reason_text"`
	OtherReason string                `json:"other_reason"`
}

type ActionBidEditRequestBody struct {
	Action BidEditRequestAction `json:"action"`
}

type ActionMilestoneRequestBody struct {
	Action MilestoneActionRequest `json:"action"`
}

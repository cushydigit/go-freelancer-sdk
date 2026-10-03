package reqres

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
	Action      MilestoneAction              `json:"action"`
	Amount      int                          `json:"amount"`
	Reason      MilestoneRequestCancelReason `json:"reason"`
	ReasonText  string                       `json:"reason_text"`
	OtherReason string                       `json:"other_reason"`
}

type ActionBidEditRequestBody struct {
	Action BidEditRequestAction `json:"action"`
}

type ActionMilestoneRequestBody struct {
	Action MilestoneActionRequest `json:"action"`
}

type ActionProjectSignNDA struct {
	Action   ProjectAction `json:"action"`
	FullName string        `json:"fullname"`
	Address  string        `json:"address"`
	City     string        `json:"city"`
	State    string        `json:"state"`
	Phone    string        `json:"phone"`
	Country  string        `json:"country"`
}

type ActionProjectUpdate struct {
	Action      ProjectAction `json:"action"`
	Description *string       `json:"description"`
	JonIDs      []int64       `json:"job_ids"`
}

type ActionProjectUpgrade struct {
	Action   ProjectAction        `json:"action"`
	Upgrades []ProjectUpgradeType `json:"upgrades"`
}

type ActionProjectClose struct {
	Action ProjectAction `json:"action"`
}

type ActionProjectEnd struct {
	Action ProjectAction        `json:"action"`
	BidID  int64                `json:"bid_id"`
	Status ProjectEndStatusType `json:"status"`
}

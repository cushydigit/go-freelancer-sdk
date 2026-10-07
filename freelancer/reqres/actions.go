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

type ActionProjectBody struct {
	// General Action
	Action ProjectAction `json:"action"`

	// SignNDA
	FullName string `json:"fullname"`
	Address  string `json:"address"`
	City     string `json:"city"`
	State    string `json:"state"`
	Phone    string `json:"phone"`
	Country  string `json:"country"`

	// Update
	Description *string `json:"description"`
	JonIDs      []int64 `json:"job_ids"`

	// Upgrade
	Upgrades []ProjectUpgradeType `json:"upgrades"`

	// End
	BidID  int64                `json:"bid_id"`
	Status ProjectEndStatusType `json:"status"`
}

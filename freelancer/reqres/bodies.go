package reqres

// Title, Description, Budget, Jobs are required
type CreateProjectBody struct {
	Title             string             `json:"title"`
	Description       string             `json:"description"`
	Budget            Budget             `json:"budget"`
	Jobs              []int64            `json:"jobs"`
	Type              *ProjectType       `json:"type,omitempty"`
	HourlyProjectInfo *HourlyProjectInfo `json:"hourly_project_info,omitempty"`
	HirMe             *bool              `json:"hire_me,omitempty"`
	HiremeInitialBid  *HiremeInitialBid  `json:"hireme_initial_bid,omitempty"`
}

type HourlyProjectInfo struct {
	Commitment Commitment `json:"commitment"`
}

type Commitment struct {
	Hours    int          `json:"hours"`
	Interval IntervalType `json:"interval"`
}

type HiremeInitialBid struct {
	BidderID int64   `json:"bidder_id"`
	Amount   float64 `json:"amount"`
	Period   int64   `json:"period"`
}

type InviteFreelancersBody struct {
	FreelancerID int64 `json:"freelancer_id"`
}

type CreateProjectCollaborationsBody struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Title    string `json:"title"`
	// required
	Permissions struct {
		Chat     bool `json:"CHAT"`
		BidAward bool `json:"BID_AWARD"`
	}
}

// ProjectID, BidderID, Amount, Period (days), MilestonePercentage (0-100) required
type CreateBidBody struct {
	ProjectID           int64  `json:"project_id"`
	BidderID            int64  `json:"bidder_id"`
	Amount              int    `json:"amount"`
	Period              int    `json:"period"`
	MilestonePercentage int    `json:"milestone_percentage"`
	Description         string `json:"description"`
	ProfileID           int    `json:"profile_id"`
}

type UpdateBidBody struct {
	Amount              int    `json:"amount"`
	MilestonePercentage int    `json:"milestone_percentage"`
	Description         string `json:"description"`
}

// StartTime and Seconds (Duration of session in seconds) required
type CreateTimeTrackingBody struct {
	StartTime int64  `json:"time_start"`
	Seconds   int    `json:"seconds"`
	Note      string `json:"note"`
}

// Required: BidID, NewAmount, NewPeriod
type CreateBidEditRequestsBody struct {
	BidID     int64 `json:"bid_id"`
	NewAmount int   `json:"new_amount"`
	NewPeriod int   `json:"new_period"`
	Comment   int   `json:"comment"`
}

type ActionBidEditRequestsBody struct {
	Action BidEditRequestAction `json:"action"`
}

// Rating required
type CreateBidRatingsBody struct {
	Rating  int    `json:"rating"`
	Comment string `json:"comment"`
}

type UpdateBidRatingBody struct {
	Rating  int    `json:"rating"`
	Comment string `json:"comment"`
}

type CreateMilestoneBody struct {
	ProjectID   int64                 `json:"project_id"`
	BidderID    int64                 `json:"bidder_id"`
	Amount      int                   `json:"amount"`
	Reason      MilestoneCreateReason `json:"reason"`
	Description string                `json:"description"`
}

type ActionMilestoneBody struct {
	Action      MilestoneAction       `json:"action"`
	Amount      int                   `json:"amount"`
	Reason      MilestoneActionReason `json:"reason"`
	ReasonText  string                `json:"reason_text"`
	OtherReason string                `json:"other_reason"`
}

// ProjectID, BidID, Amount and Description are required
type CreateMilestoneRequestBody struct {
	ProjectID   int64  `json:"project_id"`
	BidID       int    `json:"bid_id"`
	Amount      int    `json:"amount"`
	Description string `json:"description"`
}

type ActionMilestoneRequestBody struct {
	Action MilestoneActionRequest `json:"action"`
}

type CreateReviewBody struct {
	ProjectID  int64      `json:"project_id"`
	ToUserID   int64      `json:"to_user_id"`
	FromUserID int64      `json:"from_user_id"`
	ReviewType ReviewType `json:"review_type"`
	Comment    string     `json:"comment"`
	Role       RoleType   `json:"role"`
}

type ReviewActionBody struct {
	Action     ReviewAction `json:"action"`
	ReviewType ReviewType   `json:"review_type"`
}

type ExpertGuaranteesActionRequestBody struct {
	Action ExpertGuaranteesAction `json:"action"`
}

type JobsBody struct {
	Jobs []int64 `json:"jobs[]"`
}

type CreateProfileBody struct {
	Tagline     string `json:"tagline"`
	HourlyRate  int    `json:"hourly_Rate"`
	Description string `json:"description"`
	ProfileName string `json:"profile_name,omitempty"`
	SkillIDs    []int  `json:"skill_ids,omitempty"`
}

type UpdateProfileBody struct {
	ProfileID   int    `json:"profile_id"`
	Tagline     string `json:"tagline,omitempty"`
	HourlyRate  int    `json:"hourly_Rate,omitempty"`
	Description string `json:"description,omitempty"`
	ProfileName string `json:"profile_name,omitempty"`
	SkillIDs    []int  `json:"skill_ids,omitempty"`
}

type CreateViolationBody struct {
	ContextID        int                       `json:"context_id"`
	ContextType      ViolationContext          `json:"context_type"`
	ViolatorUserID   int64                     `json:"violator_user_id"`
	Reason           ViolationReason           `json:"reason"`
	AdditionalReason ViolationAdditionalReason `json:"additional_reason,omitempty"`
	Comments         string                    `json:"comments,omitempty"`
	Url              string                    `json:"url,omitempty"`
}

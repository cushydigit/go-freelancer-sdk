package freelancer

type TypeType string
type ProjectType string
type ProjectUpgradeType string
type ProjectStatusType string
type IntervalType string
type ContestUpgradeType string
type SortFieldsType string
type RoleType string
type ViolationContextType string
type ViolationReasonType string
type ViolationAdditionalReasonType string
type MilestoneStatusType string
type ProjectCollaborationActionType string
type ServiceType string
type ServiceStatusType string
type SortType string
type AwardStatusType string
type PaidStatusType string
type CompleteStatusType string
type FrontendBidStatusType string
type ActionBid string
type BidStatus string
type ActionBidEditRequest string

const (
	Projects TypeType = "projects"
	Contests TypeType = "contests"

	ProjectFixed  ProjectType = "fixed"
	ProjectHourly ProjectType = "hourly"

	IntervalWeek  IntervalType = "WEEK"
	IntervalMonth IntervalType = "MONTH"

	ProjectStatusActive  ProjectStatusType = "active"
	ProjectStatusAll     ProjectStatusType = "all"
	ProjectStatusAwarded ProjectStatusType = "awarded"
	ProjectStatusOpen    ProjectStatusType = "open"
	ProjectStatusPast    ProjectStatusType = "close"

	ProjectUpgradeFeatured   ProjectUpgradeType = "featured"
	ProjectUpgradeSealed     ProjectUpgradeType = "sealed"
	ProjectUpgradeNonpublic  ProjectUpgradeType = "nonpublic"
	ProjectUpgradeFulltime   ProjectUpgradeType = "fulltime"
	ProjectUpgradeUrgent     ProjectUpgradeType = "urgent"
	ProjectUpgradeQualified  ProjectUpgradeType = "qualified"
	ProjectUpgradeNDA        ProjectUpgradeType = "NDA"
	ProjectUpgradeAssisted   ProjectUpgradeType = "assisted"
	ProjectUpgradePFOnly     ProjectUpgradeType = "pf_only"
	ProjectUpgradeIpContract ProjectUpgradeType = "ip_contract"
	ProjectUpgradeNonCompete ProjectUpgradeType = "non_compete"

	ContestUpgradeFeatured   ContestUpgradeType = "featured"
	ContestUpgradeSealed     ContestUpgradeType = "sealed"
	ContestUpgradeNonpublic  ContestUpgradeType = "nonpublic"
	ContestUpgradeHighlight  ContestUpgradeType = "highlight"
	ContestUpgradeGuaranteed ContestUpgradeType = "guaranteed"

	SortFieldsTimeUpdated SortFieldsType = "time_updated"
	SortFieldsBidCount    SortFieldsType = "bid_count"
	SortFieldsBidEndDate  SortFieldsType = "bid_enddate"
	SortFieldsBidAvgUsd   SortFieldsType = "bid_avg_usd"

	SortNeweset     SortType = "newest"
	SortQuickest    SortType = "quickest"
	SortRecommended SortType = "recommended"

	RoleFreelancer RoleType = "freelancer"
	RoleEmployer   RoleType = "employer"

	ViolationContextProject               ViolationContextType = "project"
	ViolationContextBid                   ViolationContextType = "bid"
	ViolationContextMessage               ViolationContextType = "message"
	ViolationContextProfile               ViolationContextType = "profile"
	ViolationContextPrivateMessage        ViolationContextType = "private_message"
	ViolationContextContestEntry          ViolationContextType = "contest_entry"
	ViolationContextContestComment        ViolationContextType = "contest_comment"
	ViolationContextCopyrightInfringement ViolationContextType = "copyright_infringement"

	ViolationReasonContact     ViolationReasonType = "contacts"
	ViolationReasonAdvertising ViolationReasonType = "advertising"
	ViolationReasonFake        ViolationReasonType = "fake"
	ViolationReasonHarassment  ViolationReasonType = "harassment"
	ViolationReasonNonfeatured ViolationReasonType = "nonfeatured"
	ViolationReasonOther       ViolationReasonType = "other"

	ViolationAdditionalReasonOffsitting                     ViolationAdditionalReasonType = "offsiting_payment"
	ViolationAdditionalReasonOffsitingCommunication         ViolationAdditionalReasonType = "offsiting_communication"
	ViolationAdditionalReasonPublicDisplay                  ViolationAdditionalReasonType = "public_display_of_communication_details"
	ViolationAdditionalReasonFakeProfile                    ViolationAdditionalReasonType = "fake_profile"
	ViolationAdditionalReasonMisleadingAbilities            ViolationAdditionalReasonType = "misleading_about_own_abilities"
	ViolationAdditionalReasonMisleadingProfileContent       ViolationAdditionalReasonType = "misleading_content_on_profile"
	ViolationAdditionalReasonOrganizationNotPerson          ViolationAdditionalReasonType = "user_is_an_organization_not_a_person"
	ViolationAdditionalReasonThreatsOfViolence              ViolationAdditionalReasonType = "sending_threats_of_violence"
	ViolationAdditionalReasonInappropriateLanguage          ViolationAdditionalReasonType = "inappropriate_language"
	ViolationAdditionalReasonIssueWithUserProject           ViolationAdditionalReasonType = "issue_on_project_with_user"
	ViolationAdditionalReasonIssueWithOthersProject         ViolationAdditionalReasonType = "issue_on_project_with_someone_else"
	ViolationAdditionalReasonCopiedFromUser                 ViolationAdditionalReasonType = "copied_work_from_user"
	ViolationAdditionalReasonCopiedFromSomeoneElse          ViolationAdditionalReasonType = "copied_work_from_someone_else"
	ViolationAdditionalReasonExplicitOrInappropriateContent ViolationAdditionalReasonType = "explicit_or_inappropriate_content"
	ViolationAdditionalReasonLowQualityWork                 ViolationAdditionalReasonType = "low_quality_work"
	ViolationAdditionalReasonWorkDoesNotMatchRequirements   ViolationAdditionalReasonType = "work_does_not_match_requirements"
	ViolationAdditionalReasonOther                          ViolationAdditionalReasonType = "other"

	MilestoneStatusDeleted  MilestoneStatusType = "deleted"
	MilestoneStatusRejected MilestoneStatusType = "rejected"
	MilestoneStatusPending  MilestoneStatusType = "pending"
	MilestoneStatusCreated  MilestoneStatusType = "created"

	ProjectCollaborationActionRevoke            ProjectCollaborationActionType = "revoke"
	ProjectCollaborationActionUpdatePermissions ProjectCollaborationActionType = "update_permissions"

	ServiceRegular ServiceType = "regular"
	ServiceLocal   ServiceType = "local"

	ServiceStatusPending ServiceStatusType = "pending"
	ServiceStatusActive  ServiceStatusType = "active"
	ServiceStatusClosed  ServiceStatusType = "closed"

	AwardStatusAwarded  AwardStatusType = "awarded"
	AwardStatusRejected AwardStatusType = "rejected"
	AwardStatusRevoked  AwardStatusType = "revoked"
	AwardStatusPending  AwardStatusType = "pending"
	AwardStatusCanceled AwardStatusType = "canceled"

	PaidStatusPending    PaidStatusType = "pending"
	PaidStatusIncomplete PaidStatusType = "incomplete"
	PaidStatusComplete   PaidStatusType = "complete"

	FrontendBidStatusActive     FrontendBidStatusType = "active"
	FrontendBidStatusInprogress FrontendBidStatusType = "in_progress"
	FrontendBidStatusComplete   FrontendBidStatusType = "complete"

	ActionBidAccept                ActionBid = "accept"
	ActionBidDeny                  ActionBid = "deny"
	ActionBidRetract               ActionBid = "retract"
	ActionBidHighlight             ActionBid = "highlight"
	ActionBidSponser               ActionBid = "sponser"
	ActionBidAward                 ActionBid = "award"
	ActionBidRevoke                ActionBid = "revoke"
	ActionBidShortlist             ActionBid = "shortlist"
	ActionBidHide                  ActionBid = "hide"
	ActionBidUnhide                ActionBid = "unhide"
	ActionBidRquestLocationSharing ActionBid = "request_location_sharing"

	BidStatusPending  BidStatus = "pending"
	BidStatusAccepted BidStatus = "accepted"

	ActionBidEditRequestAccept  ActionBidEditRequest = "accept"
	ActionBidEditRequestDecline ActionBidEditRequest = "decline"
)

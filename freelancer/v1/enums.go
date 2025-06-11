package freelancer

type TypeType string
type ProjectType string
type ProjectUpgradeType string
type ProjectStatusType string
type IntervalType string
type ContestUpgradeType string
type SortField string
type RoleType string
type ViolationContext string
type ViolationReason string
type ViolationAdditionalReason string
type ProjectCollaborationAction string
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
type MilestoneStatus string

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

	SortFieldsTimeUpdated  SortField = "time_updated"
	SortFieldsBidCount     SortField = "bid_count"
	SortFieldsBidEndDate   SortField = "bid_enddate"
	SortFieldsBidAvgUsd    SortField = "bid_avg_usd"
	SortFieldsTimeReleased SortField = "time_released"

	SortNeweset     SortType = "newest"
	SortQuickest    SortType = "quickest"
	SortRecommended SortType = "recommended"

	RoleFreelancer RoleType = "freelancer"
	RoleEmployer   RoleType = "employer"

	ViolationContextProject               ViolationContext = "project"
	ViolationContextBid                   ViolationContext = "bid"
	ViolationContextMessage               ViolationContext = "message"
	ViolationContextProfile               ViolationContext = "profile"
	ViolationContextPrivateMessage        ViolationContext = "private_message"
	ViolationContextContestEntry          ViolationContext = "contest_entry"
	ViolationContextContestComment        ViolationContext = "contest_comment"
	ViolationContextCopyrightInfringement ViolationContext = "copyright_infringement"

	ViolationReasonContact     ViolationReason = "contacts"
	ViolationReasonAdvertising ViolationReason = "advertising"
	ViolationReasonFake        ViolationReason = "fake"
	ViolationReasonHarassment  ViolationReason = "harassment"
	ViolationReasonNonfeatured ViolationReason = "nonfeatured"
	ViolationReasonOther       ViolationReason = "other"

	ViolationAdditionalReasonOffsitting                     ViolationAdditionalReason = "offsiting_payment"
	ViolationAdditionalReasonOffsitingCommunication         ViolationAdditionalReason = "offsiting_communication"
	ViolationAdditionalReasonPublicDisplay                  ViolationAdditionalReason = "public_display_of_communication_details"
	ViolationAdditionalReasonFakeProfile                    ViolationAdditionalReason = "fake_profile"
	ViolationAdditionalReasonMisleadingAbilities            ViolationAdditionalReason = "misleading_about_own_abilities"
	ViolationAdditionalReasonMisleadingProfileContent       ViolationAdditionalReason = "misleading_content_on_profile"
	ViolationAdditionalReasonOrganizationNotPerson          ViolationAdditionalReason = "user_is_an_organization_not_a_person"
	ViolationAdditionalReasonThreatsOfViolence              ViolationAdditionalReason = "sending_threats_of_violence"
	ViolationAdditionalReasonInappropriateLanguage          ViolationAdditionalReason = "inappropriate_language"
	ViolationAdditionalReasonIssueWithUserProject           ViolationAdditionalReason = "issue_on_project_with_user"
	ViolationAdditionalReasonIssueWithOthersProject         ViolationAdditionalReason = "issue_on_project_with_someone_else"
	ViolationAdditionalReasonCopiedFromUser                 ViolationAdditionalReason = "copied_work_from_user"
	ViolationAdditionalReasonCopiedFromSomeoneElse          ViolationAdditionalReason = "copied_work_from_someone_else"
	ViolationAdditionalReasonExplicitOrInappropriateContent ViolationAdditionalReason = "explicit_or_inappropriate_content"
	ViolationAdditionalReasonLowQualityWork                 ViolationAdditionalReason = "low_quality_work"
	ViolationAdditionalReasonWorkDoesNotMatchRequirements   ViolationAdditionalReason = "work_does_not_match_requirements"
	ViolationAdditionalReasonOther                          ViolationAdditionalReason = "other"

	MilestoneStatusDeleted         MilestoneStatus = "deleted"
	MilestoneStatusRejected        MilestoneStatus = "rejected"
	MilestoneStatusPending         MilestoneStatus = "pending"
	MilestoneStatusCreated         MilestoneStatus = "created"
	MilestoneStatusDisputed        MilestoneStatus = "disputed"
	MilestoneStatusRquestedRelease MilestoneStatus = "requested_release"
	MilestoneStatusCleared         MilestoneStatus = "cleared"
	MilestoneStatusFrozen          MilestoneStatus = "cleared"
	MilestoneStatusCanceled        MilestoneStatus = "canceled"

	ProjectCollaborationActionRevoke            ProjectCollaborationAction = "revoke"
	ProjectCollaborationActionUpdatePermissions ProjectCollaborationAction = "update_permissions"

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

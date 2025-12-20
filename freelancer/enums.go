package freelancer

type TypeType string
type ProjectType string
type ContextType string
type ProjectUpgradeType string
type ProjectStatusType string
type ProjectAction string
type IntervalType string
type ContestUpgradeType string
type SortField string
type SortDirection string
type RoleType string
type ViolationContext string
type ViolationReason string
type ViolationAdditionalReason string
type ProjectCollaborationAction string
type ServiceType string
type ServiceStatusType string
type SortType string
type BidAwardStatus string
type BidPaidStatus string
type BidCompleteStatus string
type BidFrontendStatus string
type BidAction string
type BidStatus string
type BidEditRequestAction string
type MilestoneStatus string
type MilestoneCreateReason string
type MilestoneAction string
type MilestoneActionReason string
type MilestoneActionRequest string
type ExpertGuaranteesStatus string
type ExpertGuaranteesAction string
type ReviewType string
type CompletionStatus string
type ReviewAction string

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

	SortDirectionAsc  SortDirection = "asc"
	SortDirectionDesc SortDirection = "desc"

	SortNewest      SortType = "newest"
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

	ViolationAdditionalReasonOffsiting                      ViolationAdditionalReason = "offsiting_payment"
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

	MilestoneStatusDeleted          MilestoneStatus = "deleted"
	MilestoneStatusRejected         MilestoneStatus = "rejected"
	MilestoneStatusPending          MilestoneStatus = "pending"
	MilestoneStatusCreated          MilestoneStatus = "created"
	MilestoneStatusDisputed         MilestoneStatus = "disputed"
	MilestoneStatusRequestedRelease MilestoneStatus = "requested_release"
	MilestoneStatusCleared          MilestoneStatus = "cleared"
	MilestoneStatusFrozen           MilestoneStatus = "cleared"
	MilestoneStatusCanceled         MilestoneStatus = "canceled"

	MilestoneCreateReasonFullPayment     MilestoneCreateReason = "full_payment"
	MilestoneCreateReasonPartialPayment  MilestoneCreateReason = "partial_payment"
	MilestoneCreateReasonTaskDescription MilestoneCreateReason = "task_description"
	MilestoneCreateReasonOther           MilestoneCreateReason = "other"

	MilestoneActionRelease        MilestoneAction = "release"
	MilestoneActionUpdate         MilestoneAction = "update"
	MilestoneActionRequestCancel  MilestoneAction = "request_cancel"
	MilestoneActionRequestRelease MilestoneAction = "request_release"
	MilestoneActionCancel         MilestoneAction = "cancel"
	MilestoneActionRejectCancel   MilestoneAction = "reject_cancel"

	MilestoneActionReasonAccidentallyCreated          MilestoneActionReason = "accidentally_created"
	MilestoneActionReasonNoLongerNeeded               MilestoneActionReason = "no_longer_needed"
	MilestoneActionReasonFreelancerDidNotMeetDeadline MilestoneActionReason = "freelancer_did_not_meet_deadline"
	MilestoneActionReasonFreelancerDidNotCompleteTask MilestoneActionReason = "freelancer_did_not_complete_task"

	MilestoneActionRequestAccept MilestoneActionRequest = "accept"
	MilestoneActionRequestReject MilestoneActionRequest = "reject"
	MilestoneActionRequestDelete MilestoneActionRequest = "delete"

	ProjectCollaborationActionRevoke            ProjectCollaborationAction = "revoke"
	ProjectCollaborationActionUpdatePermissions ProjectCollaborationAction = "update_permissions"

	ServiceRegular ServiceType = "regular"
	ServiceLocal   ServiceType = "local"

	ServiceStatusPending ServiceStatusType = "pending"
	ServiceStatusActive  ServiceStatusType = "active"
	ServiceStatusClosed  ServiceStatusType = "closed"

	BidAwardStatusAwarded  BidAwardStatus = "awarded"
	BidAwardStatusRejected BidAwardStatus = "rejected"
	BidAwardStatusRevoked  BidAwardStatus = "revoked"
	BidAwardStatusPending  BidAwardStatus = "pending"
	BidAwardStatusCanceled BidAwardStatus = "canceled"

	BidCompleteStatusPending    BidCompleteStatus = "pending"
	BidCompleteStatusComplete   BidCompleteStatus = "complete"
	BidCompleteStatusIncomplete BidCompleteStatus = "incomplete"

	BidPaidStatusPending    BidPaidStatus = "pending"
	BidPaidStatusIncomplete BidPaidStatus = "incomplete"
	BidPaidStatusComplete   BidPaidStatus = "complete"

	BidFrontendStatusActive     BidFrontendStatus = "active"
	BidFrontendStatusInprogress BidFrontendStatus = "in_progress"
	BidFrontendStatusComplete   BidFrontendStatus = "complete"

	BidActionAccept                 BidAction = "accept"
	BidActionDeny                   BidAction = "deny"
	BidActionRetract                BidAction = "retract"
	BidActionHighlight              BidAction = "highlight"
	BidActionSponsor                BidAction = "sponsor"
	BidActionAward                  BidAction = "award"
	BidActionRevoke                 BidAction = "revoke"
	BidActionShortlist              BidAction = "shortlist"
	BidActionUnshortlist            BidAction = "unshortlist"
	BidActionHide                   BidAction = "hide"
	BidActionUnhide                 BidAction = "unhide"
	BidActionRequestLocationSharing BidAction = "request_location_sharing"

	BidStatusPending  BidStatus = "pending"
	BidStatusAccepted BidStatus = "accepted"

	BidEditRequestActionAccept  BidEditRequestAction = "accept"
	BidEditRequestActionDecline BidEditRequestAction = "decline"

	ExpertGuaranteesStatusDisputed         ExpertGuaranteesStatus = "disputed"
	ExpertGuaranteesStatusRequestedRelease ExpertGuaranteesStatus = "requested_release"
	ExpertGuaranteesStatusLocked           ExpertGuaranteesStatus = "locked"
	ExpertGuaranteesStatusReleased         ExpertGuaranteesStatus = "released"
	ExpertGuaranteesStatusLiquidated       ExpertGuaranteesStatus = "liquidated"
	ExpertGuaranteesStatusAdminRefunded    ExpertGuaranteesStatus = "refunded"
	ExpertGuaranteesActionRelease          ExpertGuaranteesAction = "release"
	ExpertGuaranteesActionRequestRelease   ExpertGuaranteesAction = "request_release"

	ReviewTypeProject ReviewType = "project"
	ReviewTypeContest ReviewType = "contest"

	CompletionStatusComplete   CompletionStatus = "complete"
	CompletionStatusIncomplete CompletionStatus = "incomplete"

	ReviewActionFeature   ReviewAction = "feature"
	ReviewActionUnfeature ReviewAction = "unfeature"

	ContextTypeProject ContextType = "project"
	ContextTypeContest ContextType = "contest"
	ContextTypeGeneral ContextType = "general"
)

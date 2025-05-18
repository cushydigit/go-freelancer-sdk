package freelancer

type ProjectType string
type ProjectUpgradeType string
type IntervalType string
type ContestUpgradeType string
type SortFieldsType string
type RoleType string
type ViolationContextType string
type ViolationReasonType string
type ViolationAdditionalReasonType string

const (
	ProjectTypeFixed  ProjectType = "fixed"
	ProjectTypeHourly ProjectType = "hourly"

	IntervalTypeWeek  IntervalType = "WEEK"
	IntervalTypeMonth IntervalType = "MONTH"

	ProjectUpgradeTypeFeatured   ProjectUpgradeType = "featured"
	ProjectUpgradeTypeSealed     ProjectUpgradeType = "sealed"
	ProjectUpgradeTypeNonpublic  ProjectUpgradeType = "nonpublic"
	ProjectUpgradeTypeFulltime   ProjectUpgradeType = "fulltime"
	ProjectUpgradeTypeUrgent     ProjectUpgradeType = "urgent"
	ProjectUpgradeTypeQualified  ProjectUpgradeType = "qualified"
	ProjectUpgradeTypeNDA        ProjectUpgradeType = "NDA"
	ProjectUpgradeTypeAssisted   ProjectUpgradeType = "assisted"
	ProjectUpgradeTypePFOnly     ProjectUpgradeType = "pf_only"
	ProjectUpgradeTypeIpContract ProjectUpgradeType = "ip_contract"
	ProjectUpgradeTypeNonCompete ProjectUpgradeType = "non_compete"

	ContestUpgradeTypeFeatured   ContestUpgradeType = "featured"
	ContestUpgradeTypeSealed     ContestUpgradeType = "sealed"
	ContestUpgradeTypeNonpublic  ContestUpgradeType = "nonpublic"
	ContestUpgradeTypeHighlight  ContestUpgradeType = "highlight"
	ContestUpgradeTypeGuaranteed ContestUpgradeType = "guaranteed"

	SortFieldsTypeTimeUpdated SortFieldsType = "time_updated"
	SortFieldsTypeBidCount    SortFieldsType = "bid_count"
	SortFieldsTypeBidEndDate  SortFieldsType = "bid_enddate"
	SortFieldsTypeBidAvgUsd   SortFieldsType = "bid_avg_usd"

	RoleTypeFreelancer RoleType = "freelancer"
	RoleTypeEmployer   RoleType = "employer"

	ViolationContextTypeProject               ViolationContextType = "project"
	ViolationContextTypeBid                   ViolationContextType = "bid"
	ViolationContextTypeMessage               ViolationContextType = "message"
	ViolationContextTypeProfile               ViolationContextType = "profile"
	ViolationContextTypePrivateMessage        ViolationContextType = "private_message"
	ViolationContextTypeContestEntry          ViolationContextType = "contest_entry"
	ViolationContextTypeContestComment        ViolationContextType = "contest_comment"
	ViolationContextTypeCopyrightInfringement ViolationContextType = "copyright_infringement"

	ViolationReasonTypeContact     ViolationReasonType = "contacts"
	ViolationReasonTypeAdvertising ViolationReasonType = "advertising"
	ViolationReasonTypeFake        ViolationReasonType = "fake"
	ViolationReasonTypeHarassment  ViolationReasonType = "harassment"
	ViolationReasonTypeNonfeatured ViolationReasonType = "nonfeatured"
	ViolationReasonTypeOther       ViolationReasonType = "other"

	ViolationAdditionalReasonTypeOffsitting                     ViolationAdditionalReasonType = "offsiting_payment"
	ViolationAdditionalReasonTypeOffsitingCommunication         ViolationAdditionalReasonType = "offsiting_communication"
	ViolationAdditionalReasonTypePublicDisplay                  ViolationAdditionalReasonType = "public_display_of_communication_details"
	ViolationAdditionalReasonTypeFakeProfile                    ViolationAdditionalReasonType = "fake_profile"
	ViolationAdditionalReasonTypeMisleadingAbilities            ViolationAdditionalReasonType = "misleading_about_own_abilities"
	ViolationAdditionalReasonTypeMisleadingProfileContent       ViolationAdditionalReasonType = "misleading_content_on_profile"
	ViolationAdditionalReasonTypeOrganizationNotPerson          ViolationAdditionalReasonType = "user_is_an_organization_not_a_person"
	ViolationAdditionalReasonTypeThreatsOfViolence              ViolationAdditionalReasonType = "sending_threats_of_violence"
	ViolationAdditionalReasonTypeInappropriateLanguage          ViolationAdditionalReasonType = "inappropriate_language"
	ViolationAdditionalReasonTypeIssueWithUserProject           ViolationAdditionalReasonType = "issue_on_project_with_user"
	ViolationAdditionalReasonTypeIssueWithOthersProject         ViolationAdditionalReasonType = "issue_on_project_with_someone_else"
	ViolationAdditionalReasonTypeCopiedFromUser                 ViolationAdditionalReasonType = "copied_work_from_user"
	ViolationAdditionalReasonTypeCopiedFromSomeoneElse          ViolationAdditionalReasonType = "copied_work_from_someone_else"
	ViolationAdditionalReasonTypeExplicitOrInappropriateContent ViolationAdditionalReasonType = "explicit_or_inappropriate_content"
	ViolationAdditionalReasonTypeLowQualityWork                 ViolationAdditionalReasonType = "low_quality_work"
	ViolationAdditionalReasonTypeWorkDoesNotMatchRequirements   ViolationAdditionalReasonType = "work_does_not_match_requirements"
	ViolationAdditionalReasonTypeOther                          ViolationAdditionalReasonType = "other"
)

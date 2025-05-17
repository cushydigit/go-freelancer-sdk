package freelancer

type ProjectType string
type ProjectUpgradeType string
type ContestUpgradeType string
type SortFieldsType string
type RoleType string
type ViolationContextType string

const (
	ProjectTypeFixed  ProjectType = "fixed"
	ProjectTypeHourly ProjectType = "hourly"

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
)

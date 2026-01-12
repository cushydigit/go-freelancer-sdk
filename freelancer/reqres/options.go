package reqres

// CountriesOptions holds optional filters for the ListCountries request.
type ListCountriesOptions struct {
	ExtraDetails *bool `url:"extra_details"`
}

// TimezonesOptions holds optional filters for the ListTimezones request.
type ListTimezonesOptions struct {
	Timezones     []int64  `url:"timezones[]"`
	TimezoneNames []string `url:"timezone_names[]"`
}

type ListProjectsOptions struct {
	Projects                     []int64                 `url:"projects[]"`
	Owners                       []int64                 `url:"owners[]"`
	Bidders                      []int64                 `url:"bidders[]"`
	SeoUrls                      []string                `url:"seo_urls[]"`
	FromTime                     *int64                  `url:"from_time"`
	ToTime                       *int64                  `url:"to_time"`
	FrontendProjectStatuses      []ProjectFrontendStatus `url:"frontend_project_statuses[]"`
	Team                         *bool                   `url:"team"`
	IsNonHireMe                  *bool                   `url:"is_none_hire_me"`
	HasMilestone                 *bool                   `url:"has_milestone"`
	Count                        *bool                   `url:"count"`
	FullDescription              *bool                   `url:"full_description"`
	JobDetails                   *bool                   `url:"job_details"`
	UpgradeDetails               *bool                   `url:"upgrade_details"`
	AttachmentDetails            *bool                   `url:"attachment_details"`
	FileDetails                  *bool                   `url:"file_details"`
	QualificationDetails         *bool                   `url:"qualification_details"`
	SelectedBids                 *bool                   `url:"selected_bids"`
	HiremeDetails                *bool                   `url:"hireme_details"`
	UserDetails                  *bool                   `url:"user_details"`
	InvitedFreelancerDetails     *bool                   `url:"invited_freelancer_details"`
	RecommendedFreelancerDetails *bool                   `url:"recommended_freelancer_details"`
	SupportSessionDetails        *bool                   `url:"support_session_details"`
	LocationDetails              *bool                   `url:"location_details"`
	NdaSignatureDetails          *bool                   `url:"nda_signature_details"`
	ProjectCollaborationDetails  *bool                   `url:"project_collaboration_details"`
	ProximityDetails             *bool                   `url:"proximity_details"`
	ReviewAvailabilityDetails    *bool                   `url:"review_availability_details"`
	NegotiatedDetails            *bool                   `url:"negotiated_details"`
	DriveFileDetails             *bool                   `url:"drive_file_details"`
	NdaDetails                   *bool                   `url:"nda_details"`
	LocalDetails                 *bool                   `url:"local_details"`
	EquipmentDetails             *bool                   `url:"equipment_details"`
	ClientEngagementDetails      *bool                   `url:"client_engagement_details"`
	ServiceOfferingDetails       *bool                   `url:"service_offering_details"`
	UserAvatar                   *bool                   `url:"user_avatar"`
	UserCountryDetails           *bool                   `url:"user_country_details"`
	UserProfileDescription       *bool                   `url:"user_profile_description"`
	UserDisplayInfo              *bool                   `url:"user_display_info"`
	UserJobs                     *bool                   `url:"user_jobs"`
	UserBalanceDetails           *bool                   `url:"user_balance_details"`
	UserQualificationDetails     *bool                   `url:"user_qualification_details"`
	UserMembershipDetails        *bool                   `url:"user_membership_details"`
	UserFinancialDetails         *bool                   `url:"user_financial_details"`
	UserLocationDetails          *bool                   `url:"user_location_details"`
	UserPortfolioDetails         *bool                   `url:"user_portfolio_details"`
	UserPreferredDetails         *bool                   `url:"user_preferred_details"`
	UserBadgeDetails             *bool                   `url:"user_badge_details"`
	UserStatus                   *bool                   `url:"user_status"`
	UserReputation               *bool                   `url:"user_reputation"`
	UserEmployerReputation       *bool                   `url:"user_employer_reputation"`
	UserReputationExtra          *bool                   `url:"user_reputation_extra"`
	UserEmployerReputationExtra  *bool                   `url:"user_employer_reputation_extra"`
	UserCoverImage               *bool                   `url:"user_cover_image"`
	UserPastCoverImage           *bool                   `url:"user_past_cover_image"`
	UserRecommendations          *bool                   `url:"user_recommendations"`
	UserResponsiveness           *bool                   `url:"user_responsiveness"`
	CorporateUsers               *bool                   `url:"corporate_users"`
	MarketingMobileNumber        *bool                   `url:"marketing_mobile_number"`
	SanctionDetails              *bool                   `url:"sanction_details"`
	LimitedAccount               *bool                   `url:"limited_account"`
	EquipmentGroupDetails        *bool                   `url:"equipment_group_details"`
	Limit                        *int                    `url:"limit"`
	Offset                       *int                    `url:"offset"`
	Compact                      *bool                   `url:"compact"`
}

type ListSelfProjectsOptions struct {
	Status      *ProjectStatusType `url:"status"`
	Role        *RoleType          `url:"role"`
	Types       []ProjectType      `url:"type[]"`
	Query       *string            `url:"query"`
	SortField   *SortField         `url:"sort_field"`
	ReverseSort *bool              `url:"reverse_sort"`
	Recruiter   *bool              `url:"recruiter"`
	Offset      *int               `url:"offset"`
	Limit       *int               `url:"limit"`
}

type GetProjectOptions struct {
	FullDescription              *bool `url:"full_description"`
	JobDetails                   *bool `url:"job_details"`
	UpgradeDetails               *bool `url:"upgrade_details"`
	AttachmentDetails            *bool `url:"attached_details"`
	FileDetails                  *bool `url:"file_details"`
	QualificationDetails         *bool `url:"qualification_details"`
	SelectedBids                 *bool `url:"selected_bid"`
	HiremeDetails                *bool `url:"hireme_details"`
	UserDetails                  *bool `url:"user_details"`
	InvitedFreelancerDetails     *bool `url:"invited_freelancer_details"`
	RecommendedFreelancerDetails *bool `url:"recommended_freelancer_details"`
	HourlyDetails                *bool `url:"hourly_details"`
	SupportSessionDetails        *bool `url:"support_session_details"`
	LocationDetails              *bool `url:"local_details"`
	NdaSignatureDetails          *bool `url:"nda_signature_details"`
	ProjectCollaborationDetails  *bool `url:"project_collaboration_details"`
	TrackDetails                 *bool `url:"track_details"`
	ProximityDetails             *bool `url:"proximity_details"`
	ReviewAvailabilityDetails    *bool `url:"review_availability_details"`
	NegotiatedDetails            *bool `url:"negotiated_details"`
	DriveFileDetails             *bool `url:"drive_file_details"`
	NdaDetails                   *bool `url:"nda_details"`
	LocalDetails                 *bool `url:"local_details"`
	EquipmentDetails             *bool `url:"equipment_details"`
	ClientEngagementDetails      *bool `url:"client_engagement_details"`
	ServiceOfferingDetails       *bool `url:"service_offering_details"`
	UserAvatar                   *bool `url:"user_avatar"`
	UserCountryDetails           *bool `url:"user_country_details"`
	UserProfileDescription       *bool `url:"user_profile_description"`
	UserDisplayInfo              *bool `url:"user_display_info"`
	UserJobs                     *bool `url:"user_jobs"`
	UserBalanceDetails           *bool `url:"user_balance_details"`
	UserQualificationDetails     *bool `url:"user_qualification_details"`
	UserMembershipDetails        *bool `url:"user_membership_details"`
	UserFinancialDetails         *bool `url:"user_financial_details"`
	UserLocationDetails          *bool `url:"user_location_details"`
	UserPortfolioDetails         *bool `url:"user_portfolio_details"`
	UserPreferredDetails         *bool `url:"user_preferred_details"`
	UserBadgeDetails             *bool `url:"user_badge_details"`
	UserStatus                   *bool `url:"user_status"`
	UserReputation               *bool `url:"user_reputation"`
	UserEmployerReputation       *bool `url:"user_employer_reputation"`
	UserReputationExtra          *bool `url:"user_reputation_extra"`
	UserEmployerReputationExtra  *bool `url:"user_employer_reputation_extra"`
	UserCoverImage               *bool `url:"user_cover_image"`
	UserPastCoverImage           *bool `url:"user_past_cover_image"`
	UserRecommendations          *bool `url:"user_recommendations"`
	UserResponsiveness           *bool `url:"user_responsiveness"`
	CorporateUsers               *bool `url:"corporate_users"`
	MarketingMobileNumber        *bool `url:"marketing_mobile_number"`
	SanctionDetails              *bool `url:"sanction_details"`
	LimitedAccount               *bool `url:"limited_account"`
	EquipmentGroupDetails        *bool `url:"equipment_group_details"`
	Limit                        *int  `url:"limit"`
	Offset                       *int  `url:"offset"`
	Compact                      *bool `url:"compact"`
}

type SearchActiveProjectsOptions struct {
	Query                       *string              `url:"query"`
	ProjectTypes                []ProjectBudgetType  `url:"project_types[]"`
	ProjectUpgrades             []ProjectUpgradeType `url:"project_upgrades[]"`
	ContestUpgrades             []ContestUpgradeType `url:"contest_upgrades[]"`
	MinAvgPrice                 *float64             `url:"min_avg_price"`
	MaxAvgPrice                 *float64             `url:"max_avg_price"`
	MinAvgHourlyRate            *float64             `url:"min_avg_hourly_rate"`
	MaxAvgHourlyRate            *float64             `url:"max_avg_hourly_rate"`
	MinPrice                    *float64             `url:"min_price"`
	MaxPrice                    *float64             `url:"max_price"`
	MinHourlyRate               *float64             `url:"min_hourly_rate"`
	MaxHourlyRate               *float64             `url:"max_hourly_rate"`
	Jobs                        []int64              `url:"jobs[]"`
	Countries                   []string             `url:"countries[]"`
	Languages                   []string             `url:"languages[]"`
	Latitude                    *float64             `url:"latitude"`
	Longitude                   *float64             `url:"longitude"`
	FromTime                    *int64               `url:"from_time"`
	ToTime                      *int64               `url:"to_time"`
	SortField                   *SortField           `url:"sort_field"`
	ProjectIDs                  []int64              `url:"project_ids[]"`
	TopRightLatitude            *float64             `url:"top_right_latitude"`
	TopRightLongitude           *float64             `url:"top_right_longitude"`
	BottomLeftLatitude          *float64             `url:"bottom_left_latitude"`
	BottomLeftLongitude         *float64             `url:"bottom_left_longitude"`
	ReverseSort                 *bool                `url:"reverse_sort"`
	OrSearchQuery               *string              `url:"or_search_query"`
	HighlightPreTags            *string              `url:"highlight_pre_tags"`
	HighlightPostTags           *string              `url:"highlight_post_tags"`
	FullDescription             *bool                `url:"full_description"`
	JobDetails                  *bool                `url:"job_details"`
	UpgradeDetails              *bool                `url:"upgrade_details"`
	UserDetails                 *bool                `url:"user_details"`
	LocationDetails             *bool                `url:"location_details"`
	NDASignatureDetails         *bool                `url:"nda_signature_details"`
	ProjectCollaborationDetails *bool                `url:"project_collaboration_details"`
	UserAvatar                  *bool                `url:"user_avatar"`
	UserCountryDetails          *bool                `url:"user_country_details"`
	UserProfileDescription      *bool                `url:"user_profile_description"`
	UserDisplayInfo             *bool                `url:"user_display_info"`
	UserJobs                    *bool                `url:"user_jobs"`
	UserBalanceDetails          *bool                `url:"user_balance_details"`
	UserQualificationDetails    *bool                `url:"user_qualification_details"`
	UserMembershipDetails       *bool                `url:"user_membership_details"`
	UserFinancialDetails        *bool                `url:"user_financial_details"`
	UserLocationDetails         *bool                `url:"user_location_details"`
	UserPortfolioDetails        *bool                `url:"user_portfolio_details"`
	UserPreferredDetails        *bool                `url:"user_preferred_details"`
	UserBadgeDetails            *bool                `url:"user_badge_details"`
	UserStatus                  *bool                `url:"user_status"`
	UserReputation              *bool                `url:"user_reputation"`
	UserEmployerReputation      *bool                `url:"user_employer_reputation"`
	UserReputationExtra         *bool                `url:"user_reputation_extra"`
	UserEmployerReputationExtra *bool                `url:"user_employer_reputation_extra"`
	UserCoverImage              *bool                `url:"user_cover_image"`
	UserPastCoverImage          *bool                `url:"user_past_cover_image"`
	UserRecommendations         *bool                `url:"user_recommendations"`
	UserResponsiveness          *bool                `url:"user_responsiveness"`
	CorporateUsers              *bool                `url:"corporate_users"`
	MarketingMobileNumber       *bool                `url:"marketing_mobile_number"`
	SanctionDetails             *bool                `url:"sanction_details"`
	LimitedAccount              *bool                `url:"limited_account"`
	EquipmentGroupDetails       *bool                `url:"equipment_group_details"`
	Limit                       *int                 `url:"limit"`
	Offset                      *int                 `url:"offset"`
	Compact                     *bool                `url:"compact"`
}

type SearchAllProjectsOptions struct {
	Query                       *string                 `url:"query"`
	ProjectTypes                []ProjectBudgetType     `url:"project_types[]"`
	ProjectUpgrades             []ProjectUpgradeType    `url:"project_upgrades[]"`
	ContestUpgrades             []ContestUpgradeType    `url:"contest_upgrades[]"`
	MinAvgPrice                 *float64                `url:"min_avg_price"`
	MaxAvgPrice                 *float64                `url:"max_avg_price"`
	MinAvgHourlyRate            *float64                `url:"min_avg_hourly_rate"`
	MaxAvgHourlyRate            *float64                `url:"max_avg_hourly_rate"`
	MinPrice                    *float64                `url:"min_price"`
	MaxPrice                    *float64                `url:"max_price"`
	MinHourlyRate               *float64                `url:"min_hourly_rate"`
	MaxHourlyRate               *float64                `url:"max_hourly_rate"`
	Jobs                        []int64                 `url:"jobs[]"`
	Countries                   []string                `url:"countries[]"`
	Languages                   []string                `url:"languages[]"`
	Latitude                    *float64                `url:"latitude"`
	Longitude                   *float64                `url:"longitude"`
	FromTime                    *int64                  `url:"from_time"`
	ToTime                      *int64                  `url:"to_time"`
	SortField                   *SortField              `url:"sort_field"`
	BidAwardStatuses            []BidAwardStatus        `url:"bid_award_statuses[]"`
	BidCompleteStatuses         []BidCompleteStatus     `url:"bid_complete_statuses[]"`
	ProjectStatuses             []ProjectFrontendStatus `url:"project_statuses[]"`
	ProjectIDs                  []int64                 `url:"project_ids[]"`
	TopRightLatitude            *float64                `url:"top_right_latitude"`
	TopRightLongitude           *float64                `url:"top_right_longitude"`
	BottomLeftLatitude          *float64                `url:"bottom_left_latitude"`
	BottomLeftLongitude         *float64                `url:"bottom_left_longitude"`
	ReverseSort                 *bool                   `url:"reverse_sort"`
	OrSearchQuery               *string                 `url:"or_search_query"`
	HighlightPreTags            *string                 `url:"highlight_pre_tags"`
	HighlightPostTags           *string                 `url:"highlight_post_tags"`
	FullDescription             *bool                   `url:"full_description"`
	JobDetails                  *bool                   `url:"job_details"`
	UpgradeDetails              *bool                   `url:"upgrade_details"`
	UserDetails                 *bool                   `url:"user_details"`
	UserAvatar                  *bool                   `url:"user_avatar"`
	UserCountryDetails          *bool                   `url:"user_country_details"`
	UserProfileDescription      *bool                   `url:"user_profile_description"`
	UserDisplayInfo             *bool                   `url:"user_display_info"`
	UserJobs                    *bool                   `url:"user_jobs"`
	UserBalanceDetails          *bool                   `url:"user_balance_details"`
	UserQualificationDetails    *bool                   `url:"user_qualification_details"`
	UserMembershipDetails       *bool                   `url:"user_membership_details"`
	UserFinancialDetails        *bool                   `url:"user_financial_details"`
	UserLocationDetails         *bool                   `url:"user_location_details"`
	UserPortfolioDetails        *bool                   `url:"user_portfolio_details"`
	UserPreferredDetails        *bool                   `url:"user_preferred_details"`
	UserBadgeDetails            *bool                   `url:"user_badge_details"`
	UserStatus                  *bool                   `url:"user_status"`
	UserReputation              *bool                   `url:"user_reputation"`
	UserEmployerReputation      *bool                   `url:"user_employer_reputation"`
	UserReputationExtra         *bool                   `url:"user_reputation_extra"`
	UserEmployerReputationExtra *bool                   `url:"user_employer_reputation_extra"`
	UserCoverImage              *bool                   `url:"user_cover_image"`
	UserPastCoverImage          *bool                   `url:"user_past_cover_image"`
	UserRecommendations         *bool                   `url:"user_recommendations"`
	UserResponsiveness          *bool                   `url:"user_responsiveness"`
	CorporateUsers              *bool                   `url:"corporate_users"`
	MarketingMobileNumber       *bool                   `url:"marketing_mobile_number"`
	SanctionDetails             *bool                   `url:"sanction_details"`
	LimitedAccount              *bool                   `url:"limited_account"`
	EquipmentGroupDetails       *bool                   `url:"equipment_group_details"`
	Limit                       *int                    `url:"limit"`
	Offset                      *int                    `url:"offset"`
	Compact                     *bool                   `url:"compact"`
}

type ListUpgradesFeesOptions struct {
	Currencies         []int64 `url:"currencies[]"`
	Project            *int64  `url:"project"`
	FreeUpgradeDetails *bool   `url:"fee_upgrade_details"`
	TaxIncluded        *bool   `url:"tax_included"`
}

type ListProjectBidsOptions struct {
	IsShortlisted               *bool `url:"is_shortlisted"`
	Reputation                  *bool `url:"reputation"`
	RecommendedBid              *bool `url:"recommended_bid"`
	ShortlistedBid              *bool `url:"shortlisted_bid"`
	Distance                    *bool `url:"distance"`
	UserDetails                 *bool `url:"user_details"`
	ExpertGuarantees            *bool `url:"expert_guarantees"`
	UserAvatar                  *bool `url:"user_avatar"`
	UserCountryDetails          *bool `url:"user_country_details"`
	UserProfileDescription      *bool `url:"user_profile_description"`
	UserDisplayInfo             *bool `url:"user_display_info"`
	UserJobs                    *bool `url:"user_jobs"`
	UserBalanceDetails          *bool `url:"user_balance_details"`
	UserQualificationDetails    *bool `url:"user_qualification_details"`
	UserMembershipDetails       *bool `url:"user_membership_details"`
	UserFinancialDetails        *bool `url:"user_financial_details"`
	UserLocationDetails         *bool `url:"user_location_details"`
	UserPortfolioDetails        *bool `url:"user_portfolio_details"`
	UserPreferredDetails        *bool `url:"user_preferred_details"`
	UserBadgeDetails            *bool `url:"user_badge_details"`
	UserStatus                  *bool `url:"user_status"`
	UserReputation              *bool `url:"user_reputation"`
	UserEmployerReputation      *bool `url:"user_employer_reputation"`
	UserReputationExtra         *bool `url:"user_reputation_extra"`
	UserEmployerReputationExtra *bool `url:"user_employer_reputation_extra"`
	UserCoverImage              *bool `url:"user_cover_image"`
	UserPastCoverImage          *bool `url:"user_past_cover_image"`
	UserRecommendations         *bool `url:"user_recommendations"`
	UserResponsiveness          *bool `url:"user_responsiveness"`
	CorporateUsers              *bool `url:"corporate_users"`
	MarketingMobileNumber       *bool `url:"marketing_mobile_number"`
	SanctionDetails             *bool `url:"sanction_details"`
	LimitedAccount              *bool `url:"limited_account"`
	EquipmentGroupDetails       *bool `url:"equipment_group_details"`
	Limit                       *int  `url:"limit"`
	Offset                      *int  `url:"offset"`
	Compact                     *bool `url:"compact"`
	Quotations                  *bool `url:"quotations"`
}

type ListProjectMilestonesOptions struct {
	Statuses                    []MilestoneStatus `url:"statuses[]"`
	UserAvatar                  *bool             `url:"user_avatar"`
	UserCountryDetails          *bool             `url:"user_country_details"`
	UserProfileDescription      *bool             `url:"user_profile_description"`
	UserDisplayInfo             *bool             `url:"user_display_info"`
	UserJobs                    *bool             `url:"user_jobs"`
	UserBalanceDetails          *bool             `url:"user_balance_details"`
	UserQualificationDetails    *bool             `url:"user_qualification_details"`
	UserMembershipDetails       *bool             `url:"user_membership_details"`
	UserFinancialDetails        *bool             `url:"user_financial_details"`
	UserLocationDetails         *bool             `url:"user_location_details"`
	UserPortfolioDetails        *bool             `url:"user_portfolio_details"`
	UserPreferredDetails        *bool             `url:"user_preferred_details"`
	UserBadgeDetails            *bool             `url:"user_badge_details"`
	UserStatus                  *bool             `url:"user_status"`
	UserReputation              *bool             `url:"user_reputation"`
	UserEmployerReputation      *bool             `url:"user_employer_reputation"`
	UserReputationExtra         *bool             `url:"user_reputation_extra"`
	UserEmployerReputationExtra *bool             `url:"user_employer_reputation_extra"`
	UserCoverImage              *bool             `url:"user_cover_image"`
	UserPastCoverImage          *bool             `url:"user_past_cover_image"`
	UserRecommendations         *bool             `url:"user_recommendations"`
	UserResponsiveness          *bool             `url:"user_responsiveness"`
	CorporateUsers              *bool             `url:"corporate_users"`
	MarketingMobileNumber       *bool             `url:"marketing_mobile_number"`
	SanctionDetails             *bool             `url:"sanction_details"`
	LimitedAccount              *bool             `url:"limited_account"`
	EquipmentGroupDetails       *bool             `url:"equipment_group_details"`
}

type ListProjectsMilestoneRequestsOptions struct {
	Statuses                    []MilestoneStatus `url:"statuses[]"`
	UserAvatar                  *bool             `url:"user_avatar"`
	UserCountryDetails          *bool             `url:"user_country_details"`
	UserProfileDescription      *bool             `url:"user_profile_description"`
	UserDisplayInfo             *bool             `url:"user_display_info"`
	UserJobs                    *bool             `url:"user_jobs"`
	UserBalanceDetails          *bool             `url:"user_balance_details"`
	UserQualificationDetails    *bool             `url:"user_qualification_details"`
	UserMembershipDetails       *bool             `url:"user_membership_details"`
	UserFinancialDetails        *bool             `url:"user_financial_details"`
	UserLocationDetails         *bool             `url:"user_location_details"`
	UserPortfolioDetails        *bool             `url:"user_portfolio_details"`
	UserPreferredDetails        *bool             `url:"user_preferred_details"`
	UserBadgeDetails            *bool             `url:"user_badge_details"`
	UserStatus                  *bool             `url:"user_status"`
	UserReputation              *bool             `url:"user_reputation"`
	UserEmployerReputation      *bool             `url:"user_employer_reputation"`
	UserReputationExtra         *bool             `url:"user_reputation_extra"`
	UserEmployerReputationExtra *bool             `url:"user_employer_reputation_extra"`
	UserCoverImage              *bool             `url:"user_cover_image"`
	UserPastCoverImage          *bool             `url:"user_past_cover_image"`
	UserRecommendations         *bool             `url:"user_recommendations"`
	UserResponsiveness          *bool             `url:"user_responsiveness"`
	CorporateUsers              *bool             `url:"corporate_users"`
	MarketingMobileNumber       *bool             `url:"marketing_mobile_number"`
	SanctionDetails             *bool             `url:"sanction_details"`
	LimitedAccount              *bool             `url:"limited_account"`
	EquipmentGroupDetails       *bool             `url:"equipment_group_details"`
}

type GetHourlyContractInfoOptions struct {
	ProjectIDs        []int64 `url:"project_ids[]"`
	BidderIDs         []int64 `url:"bidder_ids[]"`
	HourlyContractIDs []int64 `url:"hourly_contract_ids[]"`
	ProjectOwnerIDs   []int64 `url:"project_owner_ids[]"`
	BillingDetails    *bool   `url:"billing_details"`
	InvoiceDetails    *bool   `url:"invoice_details"`
}

type ListServicesOptions struct {
	Services                    []int64             `url:"services[]"`
	Owners                      []int64             `url:"owners[]"`
	Statuses                    []ServiceStatusType `url:"statuses[]"`
	SubStatuses                 []string            `url:"sub_statuses[]"`
	Titles                      []string            `url:"titles[]"`
	SeoUrls                     []string            `url:"seo_urls[]"`
	ExtraDetails                *bool               `url:"extra_details"`
	FileDetails                 *bool               `url:"file_details"`
	JobDetails                  *bool               `url:"job_details"`
	UserDetails                 *bool               `url:"user_details"`
	FullDescription             *bool               `url:"full_description"`
	UserAvatar                  *bool               `url:"user_avatar"`
	UserCountryDetails          *bool               `url:"user_country_details"`
	UserProfileDescription      *bool               `url:"user_profile_description"`
	UserDisplayInfo             *bool               `url:"user_display_info"`
	UserJobs                    *bool               `url:"user_jobs"`
	UserBalanceDetails          *bool               `url:"user_balance_details"`
	UserQualificationDetails    *bool               `url:"user_qualification_details"`
	UserMembershipDetails       *bool               `url:"user_membership_details"`
	UserFinancialDetails        *bool               `url:"user_financial_details"`
	UserLocationDetails         *bool               `url:"user_location_details"`
	UserPortfolioDetails        *bool               `url:"user_portfolio_details"`
	UserPreferredDetails        *bool               `url:"user_preferred_details"`
	UserBadgeDetails            *bool               `url:"user_badge_details"`
	UserStatus                  *bool               `url:"user_status"`
	UserReputation              *bool               `url:"user_reputation"`
	UserEmployerReputation      *bool               `url:"user_employer_reputation"`
	UserReputationExtra         *bool               `url:"user_reputation_extra"`
	UserEmployerReputationExtra *bool               `url:"user_employer_reputation_extra"`
	UserCoverImage              *bool               `url:"user_cover_image"`
	UserPastCoverImage          *bool               `url:"user_past_cover_image"`
	UserRecommendations         *bool               `url:"user_recommendations"`
	UserResponsiveness          *bool               `url:"user_responsiveness"`
	CorporateUsers              *bool               `url:"corporate_users"`
	MarketingMobileNumber       *bool               `url:"marketing_mobile_number"`
	SanctionDetails             *bool               `url:"sanction_details"`
	LimitedAccount              *bool               `url:"limited_account"`
	EquipmentGroupDetails       *bool               `url:"equipment_group_details"`
	Limit                       *int                `url:"limit"`
	Offset                      *int                `url:"offset"`
	Compact                     *bool               `url:"compact"`
}

type SearchActiveServicesOptions struct {
	Query        *string   `url:"query"`
	Sort         *SortType `url:"sort"`
	ReverseSort  *bool     `url:"reverse_sort"`
	ExtraDetails *bool     `url:"extra_details"`
	FileDetails  *bool     `url:"file_details"`
	JobDetails   *bool     `url:"job_details"`
	UserDetails  *bool     `url:"user_details"`
	Offset       *int      `url:"offset"`
	Compact      *bool     `url:"compact"`
}

type ListBidsOptions struct {
	Bids                        []int64             `url:"bids[]"`
	Projects                    []int64             `url:"projects[]"`
	Bidders                     []int64             `url:"bidders[]"`
	ProjectOwners               []int64             `url:"project_owners[]"`
	AwardStatuses               []BidAwardStatus    `url:"award_statuses[]"`
	PaidStatuses                []BidPaidStatus     `url:"paid_statuses[]"`
	CompleteStatuses            []BidCompleteStatus `url:"complete_statuses[]"`
	FrontBidStatuses            []BidFrontendStatus `url:"front_bid_statuses[]"`
	FromTime                    *int64              `url:"from_time"`
	ToTime                      *int64              `url:"to_time"`
	Reputation                  *bool               `url:"reputation"`
	BuyerProjectFee             *bool               `url:"buyer_project_fee"`
	AwardStatusPossibilities    *bool               `url:"award_status_possibilities"`
	ProjectDetails              *bool               `url:"project_details"`
	UserDetails                 *bool               `url:"user_details"`
	UserAvatar                  *bool               `url:"user_avatar"`
	UserCountryDetails          *bool               `url:"user_country_details"`
	UserProfileDescription      *bool               `url:"user_profile_description"`
	UserDisplayInfo             *bool               `url:"user_display_info"`
	UserJobs                    *bool               `url:"user_jobs"`
	UserBalanceDetails          *bool               `url:"user_balance_details"`
	UserQualificationDetails    *bool               `url:"user_qualification_details"`
	UserMembershipDetails       *bool               `url:"user_membership_details"`
	UserFinancialDetails        *bool               `url:"user_financial_details"`
	UserLocationDetails         *bool               `url:"user_location_details"`
	UserPortfolioDetails        *bool               `url:"user_portfolio_details"`
	UserPreferredDetails        *bool               `url:"user_preferred_details"`
	UserBadgeDetails            *bool               `url:"user_badge_details"`
	UserStatus                  *bool               `url:"user_status"`
	UserReputation              *bool               `url:"user_reputation"`
	UserEmployerReputation      *bool               `url:"user_employer_reputation"`
	UserReputationExtra         *bool               `url:"user_reputation_extra"`
	UserEmployerReputationExtra *bool               `url:"user_employer_reputation_extra"`
	UserCoverImage              *bool               `url:"user_cover_image"`
	UserPastCoverImage          *bool               `url:"user_past_cover_image"`
	UserRecommendations         *bool               `url:"user_recommendations"`
	UserResponsiveness          *bool               `url:"user_responsiveness"`
	CorporateUsers              *bool               `url:"corporate_users"`
	MarketingMobileNumber       *bool               `url:"marketing_mobile_number"`
	SanctionDetails             *bool               `url:"sanction_details"`
	LimitedAccount              *bool               `url:"limited_account"`
	EquipmentGroupDetails       *bool               `url:"equipment_group_details"`
	Limit                       *int                `url:"limit"`
	Offset                      *int                `url:"offset"`
	Compact                     *bool               `url:"compact"`
}

type GetBidOptions struct {
	Reputation                  *bool `url:"reputation"`
	BuyerProjectFee             *bool `url:"buyer_project_fee"`
	AwardStatusPossibilities    *bool `url:"award_status_possibilities"`
	ProjectDetails              *bool `url:"project_details"`
	UserDetails                 *bool `url:"user_details"`
	UserAvatar                  *bool `url:"user_avatar"`
	UserCountryDetails          *bool `url:"user_country_details"`
	UserProfileDescription      *bool `url:"user_profile_description"`
	UserDisplayInfo             *bool `url:"user_display_info"`
	UserJobs                    *bool `url:"user_jobs"`
	UserBalanceDetails          *bool `url:"user_balance_details"`
	UserQualificationDetails    *bool `url:"user_qualification_details"`
	UserMembershipDetails       *bool `url:"user_membership_details"`
	UserFinancialDetails        *bool `url:"user_financial_details"`
	UserLocationDetails         *bool `url:"user_location_details"`
	UserPortfolioDetails        *bool `url:"user_portfolio_details"`
	UserPreferredDetails        *bool `url:"user_preferred_details"`
	UserBadgeDetails            *bool `url:"user_badge_details"`
	UserStatus                  *bool `url:"user_status"`
	UserReputation              *bool `url:"user_reputation"`
	UserEmployerReputation      *bool `url:"user_employer_reputation"`
	UserReputationExtra         *bool `url:"user_reputation_extra"`
	UserEmployerReputationExtra *bool `url:"user_employer_reputation_extra"`
	UserCoverImage              *bool `url:"user_cover_image"`
	UserPastCoverImage          *bool `url:"user_past_cover_image"`
	UserRecommendations         *bool `url:"user_recommendations"`
	UserResponsiveness          *bool `url:"user_responsiveness"`
	CorporateUsers              *bool `url:"corporate_users"`
	MarketingMobileNumber       *bool `url:"marketing_mobile_number"`
	SanctionDetails             *bool `url:"sanction_details"`
	LimitedAccount              *bool `url:"limited_account"`
	EquipmentGroupDetails       *bool `url:"equipment_group_details"`
	Limit                       *int  `url:"limit"`
	Offset                      *int  `url:"offset"`
	Compact                     *bool `url:"compact"`
	Quotations                  *bool `url:"quotations"`
}

type GetTimeTrackingOptions struct {
	FromTime              *int64 `url:"from_time"`
	ToTime                *int64 `url:"to_time"`
	DailyAggregateDetails *bool  `url:"daily_aggregate_details"`
	Invoiced              *bool  `url:"invoiced"`
}

type ListBidEditRequestsOptions struct {
	Statuses          []BidStatus `url:"statuses[]"`
	BidEditRequestIDs []int64     `url:"bid_edit_request_ids[]"`
}

type GetByListOfBidsOptions struct {
	Bids []int64 `url:"bids[]"`
}

type ListJobsOptions struct {
	Jobs                      []int64  `url:"jobs[]"`
	JobNames                  []string `url:"job_names[]"`
	SeoUrls                   []string `url:"seo_urls[]"`
	Categories                []int64  `url:"categories[]"`
	OnlyLocal                 *bool    `url:"only_local"`
	ActiveProjectCountDetails *bool    `url:"active_project_count_details"`
	SeoDetails                *bool    `url:"seo_details"`
	SeoCountryName            *string  `url:"seo_country_name"`
	Lang                      *string  `url:"lang"`
}

type SearchJobsOptions struct {
	Jobs                      []int64  `url:"jobs[]"`
	JobNames                  []string `url:"job_names[]"`
	SeoUrls                   []string `url:"seo_urls[]"`
	SeoTexts                  []string `url:"seo_texts[]"`
	Categories                []int64  `url:"categories[]"`
	OnlyLocal                 *bool    `url:"only_local"`
	ActiveProjectCountDetails *bool    `url:"active_project_count_details"`
	SeoDetails                *bool    `url:"seo_details"`
	SeoCountryName            *string  `url:"seo_country_name"`
	Lang                      *string  `url:"lang"`
}

type ListJobBundlesOptions struct {
	JobBundles []int64 `url:"job_bundles[]"`
	Categories []int64 `url:"categories[]"`
	Lang       *string `url:"lang"`
}

type ListJobBundleCategoriesOptions struct {
	JobBundles []int64 `url:"job_bundles[]"`
	Categories []int64 `url:"categories[]"`
	Lang       *string `url:"lang"`
}

type ListMilestonesOptions struct {
	Projects                    []int64           `url:"projects[]"`
	ProjectOwners               []int64           `url:"project_owners[]"`
	Bidders                     []int64           `url:"bidders[]"`
	Users                       []int64           `url:"users[]"`
	Bids                        []int64           `url:"bids[]"`
	Statuses                    []MilestoneStatus `url:"statuses[]"`
	SortField                   *SortField        `url:"sort_field"`
	SortDirection               *SortDirection    `url:"sort_direction"`
	ExcludedMilestones          *bool             `url:"excluded_milestones"`
	UserAvatar                  *bool             `url:"user_avatar"`
	UserCountryDetails          *bool             `url:"user_country_details"`
	UserProfileDescription      *bool             `url:"user_profile_description"`
	UserDisplayInfo             *bool             `url:"user_display_info"`
	UserJobs                    *bool             `url:"user_jobs"`
	UserBalanceDetails          *bool             `url:"user_balance_details"`
	UserQualificationDetails    *bool             `url:"user_qualification_details"`
	UserMembershipDetails       *bool             `url:"user_membership_details"`
	UserFinancialDetails        *bool             `url:"user_financial_details"`
	UserLocationDetails         *bool             `url:"user_location_details"`
	UserPortfolioDetails        *bool             `url:"user_portfolio_details"`
	UserPreferredDetails        *bool             `url:"user_preferred_details"`
	UserBadgeDetails            *bool             `url:"user_badge_details"`
	UserStatus                  *bool             `url:"user_status"`
	UserReputation              *bool             `url:"user_reputation"`
	UserEmployerReputation      *bool             `url:"user_employer_reputation"`
	UserReputationExtra         *bool             `url:"user_reputation_extra"`
	UserEmployerReputationExtra *bool             `url:"user_employer_reputation_extra"`
	UserCoverImage              *bool             `url:"user_cover_image"`
	UserPastCoverImage          *bool             `url:"user_past_cover_image"`
	UserRecommendations         *bool             `url:"user_recommendations"`
	UserResponsiveness          *bool             `url:"user_responsiveness"`
	CorporateUsers              *bool             `url:"corporate_users"`
	MarketingMobileNumber       *bool             `url:"marketing_mobile_number"`
	SanctionDetails             *bool             `url:"sanction_details"`
	LimitedAccount              *bool             `url:"limited_account"`
	EquipmentGroupDetails       *bool             `url:"equipment_group_details"`
}

type GetMilestoneOptions struct {
	MilestoneID                 int   `url:"milestone_id"`
	UserAvatar                  *bool `url:"user_avatar"`
	UserCountryDetails          *bool `url:"user_country_details"`
	UserProfileDescription      *bool `url:"user_profile_description"`
	UserDisplayInfo             *bool `url:"user_display_info"`
	UserJobs                    *bool `url:"user_jobs"`
	UserBalanceDetails          *bool `url:"user_balance_details"`
	UserQualificationDetails    *bool `url:"user_qualification_details"`
	UserMembershipDetails       *bool `url:"user_membership_details"`
	UserFinancialDetails        *bool `url:"user_financial_details"`
	UserLocationDetails         *bool `url:"user_location_details"`
	UserPortfolioDetails        *bool `url:"user_portfolio_details"`
	UserPreferredDetails        *bool `url:"user_preferred_details"`
	UserBadgeDetails            *bool `url:"user_badge_details"`
	UserStatus                  *bool `url:"user_status"`
	UserReputation              *bool `url:"user_reputation"`
	UserEmployerReputation      *bool `url:"user_employer_reputation"`
	UserReputationExtra         *bool `url:"user_reputation_extra"`
	UserEmployerReputationExtra *bool `url:"user_employer_reputation_extra"`
	UserCoverImage              *bool `url:"user_cover_image"`
	UserPastCoverImage          *bool `url:"user_past_cover_image"`
	UserRecommendations         *bool `url:"user_recommendations"`
	UserResponsiveness          *bool `url:"user_responsiveness"`
	CorporateUsers              *bool `url:"corporate_users"`
	MarketingMobileNumber       *bool `url:"marketing_mobile_number"`
	SanctionDetails             *bool `url:"sanction_details"`
	LimitedAccount              *bool `url:"limited_account"`
	EquipmentGroupDetails       *bool `url:"equipment_group_details"`
}

type ListMilestoneRequestsOptions struct {
	MilestoneRequests           []int64           `url:"milestone_requests[]"`
	Projects                    []int64           `url:"projects[]"`
	ProjectOwners               []int64           `url:"project_owners[]"`
	Bidders                     []int64           `url:"bidders[]"`
	Users                       []int64           `url:"users[]"`
	Bids                        []int64           `url:"bids[]"`
	Statuses                    []MilestoneStatus `url:"statuses[]"`
	FromTime                    *int64            `url:"from_time"`
	ToTime                      *int64            `url:"to_time"`
	SortField                   *SortField        `url:"sort_field"`
	SortDirection               *SortDirection    `url:"sort_direction"`
	ExcludedMilestones          *bool             `url:"excluded_milestones"`
	UserAvatar                  *bool             `url:"user_avatar"`
	UserCountryDetails          *bool             `url:"user_country_details"`
	UserProfileDescription      *bool             `url:"user_profile_description"`
	UserDisplayInfo             *bool             `url:"user_display_info"`
	UserJobs                    *bool             `url:"user_jobs"`
	UserBalanceDetails          *bool             `url:"user_balance_details"`
	UserQualificationDetails    *bool             `url:"user_qualification_details"`
	UserMembershipDetails       *bool             `url:"user_membership_details"`
	UserFinancialDetails        *bool             `url:"user_financial_details"`
	UserLocationDetails         *bool             `url:"user_location_details"`
	UserPortfolioDetails        *bool             `url:"user_portfolio_details"`
	UserPreferredDetails        *bool             `url:"user_preferred_details"`
	UserBadgeDetails            *bool             `url:"user_badge_details"`
	UserStatus                  *bool             `url:"user_status"`
	UserReputation              *bool             `url:"user_reputation"`
	UserEmployerReputation      *bool             `url:"user_employer_reputation"`
	UserReputationExtra         *bool             `url:"user_reputation_extra"`
	UserEmployerReputationExtra *bool             `url:"user_employer_reputation_extra"`
	UserCoverImage              *bool             `url:"user_cover_image"`
	UserPastCoverImage          *bool             `url:"user_past_cover_image"`
	UserRecommendations         *bool             `url:"user_recommendations"`
	UserResponsiveness          *bool             `url:"user_responsiveness"`
	CorporateUsers              *bool             `url:"corporate_users"`
	MarketingMobileNumber       *bool             `url:"marketing_mobile_number"`
	SanctionDetails             *bool             `url:"sanction_details"`
	LimitedAccount              *bool             `url:"limited_account"`
	EquipmentGroupDetails       *bool             `url:"equipment_group_details"`
	Limit                       *int              `url:"limit"`
	Offset                      *int              `url:"offset"`
}

type GetMilestoneRequestOptions struct {
	UserAvatar                  *bool `url:"user_avatar"`
	UserCountryDetails          *bool `url:"user_country_details"`
	UserProfileDescription      *bool `url:"user_profile_description"`
	UserDisplayInfo             *bool `url:"user_display_info"`
	UserJobs                    *bool `url:"user_jobs"`
	UserBalanceDetails          *bool `url:"user_balance_details"`
	UserQualificationDetails    *bool `url:"user_qualification_details"`
	UserMembershipDetails       *bool `url:"user_membership_details"`
	UserFinancialDetails        *bool `url:"user_financial_details"`
	UserLocationDetails         *bool `url:"user_location_details"`
	UserPortfolioDetails        *bool `url:"user_portfolio_details"`
	UserPreferredDetails        *bool `url:"user_preferred_details"`
	UserBadgeDetails            *bool `url:"user_badge_details"`
	UserStatus                  *bool `url:"user_status"`
	UserReputation              *bool `url:"user_reputation"`
	UserEmployerReputation      *bool `url:"user_employer_reputation"`
	UserReputationExtra         *bool `url:"user_reputation_extra"`
	UserEmployerReputationExtra *bool `url:"user_employer_reputation_extra"`
	UserCoverImage              *bool `url:"user_cover_image"`
	UserPastCoverImage          *bool `url:"user_past_cover_image"`
	UserRecommendations         *bool `url:"user_recommendations"`
	UserResponsiveness          *bool `url:"user_responsiveness"`
	CorporateUsers              *bool `url:"corporate_users"`
	MarketingMobileNumber       *bool `url:"marketing_mobile_number"`
	SanctionDetails             *bool `url:"sanction_details"`
	LimitedAccount              *bool `url:"limited_account"`
	EquipmentGroupDetails       *bool `url:"equipment_group_details"`
}

type ListReviewsOptions struct {
	Projects                    []int64            `url:"projects[]"`
	FromUsers                   []int64            `url:"from_users[]"`
	ToUsers                     []int64            `url:"to_users[]"`
	Contests                    []int64            `url:"contests[]"`
	ReviewTypes                 []ReviewType       `url:"review_types[]"`
	JobIds                      []int64            `url:"job_ids[]"`
	CompletionStatuses          []CompletionStatus `url:"completion_statuses[]"`
	FromTime                    *int64             `url:"from_time"`
	ToTime                      *int64             `url:"to_time"`
	ReviewStatus                []string           `url:"review_status[]"`
	ProjectDetails              *bool              `url:"project_details"`
	Ratings                     *bool              `url:"ratings"`
	ReviewCount                 *bool              `url:"review_count"`
	Role                        *RoleType          `url:"role"`
	ContestDetails              *bool              `url:"contest_details"`
	UserDetails                 *bool              `url:"user_details"`
	ProjectFullDescription      *bool              `url:"project_full_description"`
	ProjectUpgradeDetails       *bool              `url:"project_upgrade_details"`
	ProjectJobDetails           *bool              `url:"project_job_details"`
	ProjectSelectedBids         *bool              `url:"project_selected_bids"`
	ProjectQualificationDetails *bool              `url:"project_qualification_details"`
	ProjectAttachmentDetails    *bool              `url:"project_attachment_details"`
	ProjectHiremeDetails        *bool              `url:"project_hireme_details"`
	ContestJobDetails           *bool              `url:"contest_job_details"`
	UserAvatar                  *bool              `url:"user_avatar"`
	UserCountryDetails          *bool              `url:"user_country_details"`
	UserProfileDescription      *bool              `url:"user_profile_description"`
	UserDisplayInfo             *bool              `url:"user_display_info"`
	UserJobs                    *bool              `url:"user_jobs"`
	UserBalanceDetails          *bool              `url:"user_balance_details"`
	UserQualificationDetails    *bool              `url:"user_qualification_details"`
	UserMembershipDetails       *bool              `url:"user_membership_details"`
	UserFinancialDetails        *bool              `url:"user_financial_details"`
	UserLocationDetails         *bool              `url:"user_location_details"`
	UserPortfolioDetails        *bool              `url:"user_portfolio_details"`
	UserPreferredDetails        *bool              `url:"user_preferred_details"`
	UserBadgeDetails            *bool              `url:"user_badge_details"`
	UserStatus                  *bool              `url:"user_status"`
	UserReputation              *bool              `url:"user_reputation"`
	UserEmployerReputation      *bool              `url:"user_employer_reputation"`
	UserReputationExtra         *bool              `url:"user_reputation_extra"`
	UserEmployerReputationExtra *bool              `url:"user_employer_reputation_extra"`
	UserCoverImage              *bool              `url:"user_cover_image"`
	UserPastCoverImage          *bool              `url:"user_past_cover_image"`
	UserRecommendations         *bool              `url:"user_recommendations"`
	UserResponsiveness          *bool              `url:"user_responsiveness"`
	CorporateUsers              *bool              `url:"corporate_users"`
	MarketingMobileNumber       *bool              `url:"marketing_mobile_number"`
	SanctionDetails             *bool              `url:"sanction_details"`
	LimitedAccount              *bool              `url:"limited_account"`
	EquipmentGroupDetails       *bool              `url:"equipment_group_details"`
	Limit                       *int               `url:"limit"`
	Offset                      *int               `url:"offset"`
	Compact                     *bool              `url:"compact"`
}

type ListExpertGuaranteesOptions struct {
	ExpertGuarantees []int64                  `url:"expert_guarantees[]"`
	Projects         []int64                  `url:"projects[]"`
	ProjectOwners    []int64                  `url:"project_owners[]"`
	Bidders          []int64                  `url:"bidders[]"`
	Bids             []int64                  `url:"bids[]"`
	Statuses         []ExpertGuaranteesStatus `url:"statuses[]"`
	FromTime         *int64                   `url:"from_time"`
	ToTime           *int64                   `url:"to_time"`
	Offset           *int                     `url:"offset"`
	Limit            *int                     `url:"limit"`
}

type ListCurrenciesOptions struct {
	CurrencyCodes             []string `url:"currency_codes[]"`
	CurrencyIDs               []int64  `url:"currency_ids[]"`
	IncludeExternalCurrencies *bool    `url:"include_external_currencies"`
}

type ListCategoriesOptions struct {
	Categories                []int64 `url:"categories[]"`
	JobDetails                *bool   `url:"job_details"`
	Lang                      *string `url:"lang"`
	ActiveProjectCountDetails *bool   `url:"active_project_count_details"`
	SeoDetails                *bool   `url:"seo_details"`
}

type ListBudgetsOptions struct {
	CurrencyCodes   []string           `url:"currency_codes[]"`
	CurrencyIDs     []int64            `url:"currency_ids[]"`
	ProjectType     *ProjectBudgetType `url:"project_type"`
	Lang            *string            `url:"lang"`
	CurrencyDetails *bool              `url:"currency_details"`
}

type ListUsersOptions struct {
	Users                         []int64  `url:"users[]"`
	Usernames                     []string `url:"usernames[]"`
	Avatar                        *bool    `url:"avatar"`
	CountryDetails                *bool    `url:"country_details"`
	ProfileDescription            *bool    `url:"profile_description"`
	DisplayInfo                   *bool    `url:"display_info"`
	Jobs                          *bool    `url:"jobs"`
	BalanceDetails                *bool    `url:"balance_details"`
	QualificationDetails          *bool    `url:"qualification_details"`
	MembershipDetails             *bool    `url:"membership_details"`
	FinancialDetails              *bool    `url:"financial_details"`
	LocationDetails               *bool    `url:"location_details"`
	PortfolioDetails              *bool    `url:"portfolio_details"`
	PreferredDetails              *bool    `url:"preferred_details"`
	BadgeDetails                  *bool    `url:"badge_details"`
	Status                        *bool    `url:"status"`
	Reputation                    *bool    `url:"reputation"`
	EmployerReputation            *bool    `url:"employer_reputation"`
	ReputationExtra               *bool    `url:"reputation_extra"`
	EmployerReputationExtra       *bool    `url:"employer_reputation_extra"`
	CoverImage                    *bool    `url:"cover_image"`
	PastCoverImage                *bool    `url:"past_cover_image"`
	MobileTracking                *bool    `url:"mobile_tracking"`
	BidQualityDetails             *bool    `url:"bid_quality_details"`
	DepositMethods                *bool    `url:"deposit_methods"`
	UserRecommendations           *bool    `url:"user_recommendations"`
	MarketingMobileNumber         *bool    `url:"marketing_mobile_number"`
	SanctionDetails               *bool    `url:"sanction_details"`
	LimitedAccount                *bool    `url:"limited_account"`
	CompletedUserRelevantJobCount *bool    `url:"completed_user_relevant_job_count"`
	EquipmentGroupDetails         *bool    `url:"equipment_group_details"`
	JobRanks                      *bool    `url:"job_ranks"`
	JobSeoDetails                 *bool    `url:"job_seo_details"`
	RisingStar                    *bool    `url:"rising_star"`
	ShareholderDetails            *bool    `url:"shareholder_details"`
	StaffDetails                  *bool    `url:"staff_details"`
}

type SearchFreelancerOptions struct {
	Query                         *string  `url:"query"`
	JobsIDs                       []int64  `url:"jobs[]"`
	Skills                        []int64  `url:"skills[]"`
	Countries                     []string `url:"countries[]"`
	HourlyRateMin                 *int     `url:"hourly_rate_min"`
	HourlyRateMax                 *int     `url:"hourly_rate_max"`
	ReviewCountMin                *int     `url:"review_count_min"`
	ReviewCountMax                *int     `url:"review_count_max"`
	OnlineOnly                    *bool    `url:"online_only"`
	LocationLatitude              *float64 `url:"location_latitude"`
	LocationLongitude             *float64 `url:"location_longitude"`
	Insignias                     []int64  `url:"insignias[]"`
	PoolIds                       []int64  `url:"pool_ids[]"`
	Ratings                       *float64 `url:"ratings"`
	SortField                     *int     `url:"sort_field"`
	ReverseSort                   *bool    `url:"reverse_sort"`
	Avatar                        *bool    `url:"avatar"`
	CountryDetails                *bool    `url:"country_details"`
	ProfileDescription            *bool    `url:"profile_description"`
	DisplayInfo                   *bool    `url:"display_info"`
	Jobs                          *bool    `url:"jobs"`
	BalanceDetails                *bool    `url:"balance_details"`
	QualificationDetails          *bool    `url:"qualification_details"`
	MembershipDetails             *bool    `url:"membership_details"`
	FinancialDetails              *bool    `url:"financial_details"`
	LocationDetails               *bool    `url:"location_details"`
	PortfolioDetails              *bool    `url:"portfolio_details"`
	PreferredDetails              *bool    `url:"preferred_details"`
	BadgeDetails                  *bool    `url:"badge_details"`
	Status                        *bool    `url:"status"`
	Reputation                    *bool    `url:"reputation"`
	EmployerReputation            *bool    `url:"employer_reputation"`
	ReputationExtra               *bool    `url:"reputation_extra"`
	EmployerReputationExtra       *bool    `url:"employer_reputation_extra"`
	CoverImage                    *bool    `url:"cover_image"`
	PastCoverImage                *bool    `url:"past_cover_image"`
	MobileTracking                *bool    `url:"mobile_tracking"`
	BidQualityDetails             *bool    `url:"bid_quality_details"`
	DepositMethods                *bool    `url:"deposit_methods"`
	UserRecommendations           *bool    `url:"user_recommendations"`
	MarketingMobileNumber         *bool    `url:"marketing_mobile_number"`
	SanctionDetails               *bool    `url:"sanction_details"`
	LimitedAccount                *bool    `url:"limited_account"`
	CompletedUserRelevantJobCount *bool    `url:"completed_user_relevant_job_count"`
	EquipmentGroupDetails         *bool    `url:"equipment_group_details"`
	JobRanks                      *bool    `url:"job_ranks"`
	JobSeoDetails                 *bool    `url:"job_seo_details"`
	RisingStar                    *bool    `url:"rising_star"`
	ShareholderDetails            *bool    `url:"shareholder_details"`
	StaffDetails                  *bool    `url:"staff_details"`
	PoolDetails                   *bool    `url:"pool_details"`
	Limit                         *int     `url:"limit"`
	Offset                        *int     `url:"offset"`
	Compact                       *bool    `url:"compact"`
}

type GetSelfInfoOptions struct {
	Avatar                        *bool `url:"avatar"`
	CountryDetails                *bool `url:"country_details"`
	ProfileDescription            *bool `url:"profile_description"`
	DisplayInfo                   *bool `url:"display_info"`
	Jobs                          *bool `url:"jobs"`
	BalanceDetails                *bool `url:"balance_details"`
	QualificationDetails          *bool `url:"qualification_details"`
	MembershipDetails             *bool `url:"membership_details"`
	FinancialDetails              *bool `url:"financial_details"`
	LocationDetails               *bool `url:"location_details"`
	PortfolioDetails              *bool `url:"portfolio_details"`
	PreferredDetails              *bool `url:"preferred_details"`
	BadgeDetails                  *bool `url:"badge_details"`
	Status                        *bool `url:"status"`
	Reputation                    *bool `url:"reputation"`
	EmployerReputation            *bool `url:"employer_reputation"`
	ReputationExtra               *bool `url:"reputation_extra"`
	EmployerReputationExtra       *bool `url:"employer_reputation_extra"`
	CoverImage                    *bool `url:"cover_image"`
	PastCoverImages               *bool `url:"past_cover_images"`
	MobileTracking                *bool `url:"mobile_tracking"`
	BidQualityDetails             *bool `url:"bid_quality_details"`
	DepositMethods                *bool `url:"deposit_methods"`
	UserRecommendations           *bool `url:"user_recommendations"`
	MarketingMobileNumber         *bool `url:"marketing_mobile_number"`
	SanctionDetails               *bool `url:"sanction_details"`
	LimitedAccount                *bool `url:"limited_account"`
	CompletedUserRelevantJobCount *bool `url:"completed_user_relevant_job_count"`
	EquipmentGroupDetails         *bool `url:"equipment_group_details"`
	JobRanks                      *bool `url:"job_ranks"`
	JobSeoDetails                 *bool `url:"job_seo_details"`
	RisingStar                    *bool `url:"rising_star"`
	ShareholderDetails            *bool `url:"shareholder_details"`
	StaffDetails                  *bool `url:"staff_details"`
	Limit                         *int  `url:"limit"`
	Offset                        *int  `url:"offset"`
	Compact                       *bool `url:"compact"`
}

type ListReputationsOptions struct {
	Users        []int64   `url:"users"`
	Jobs         []int64   `url:"jobs"`
	Role         *RoleType `url:"role"`
	JobHistory   *bool     `url:"job_history"`
	ProjectStats *bool     `url:"project_stats"`
	RehireRates  *bool     `url:"rehire_rates"`
}

type ListEnterprisesOptions struct {
	Enterprises   []int64  `url:"enterprises[]"`
	InternalNames []string `url:"internal_names[]"`
	Names         []string `url:"names[]"`
	SeoUrls       []string `url:"seo_urls[]"`
	UserID        *int64   `url:"user_id"`
	IgnoreTest    *bool    `url:"ignore_test"`
	Limit         *int     `url:"limit"`
	Offset        *int     `url:"offset"`
}

type ListPortfoliosOptions struct {
	Users  []int64 `url:"users[]"`
	Limit  *int    `url:"limit"`
	Offset *int    `url:"offset"`
}

type ListPoolsOptions struct {
	Pools           []int64  `url:"pools[]"`
	Names           []string `url:"internal_names[]"`
	SeoUrls         []string `url:"seo_urls[]"`
	IgnoreTest      *bool    `url:"ignore_test"`
	IsTalentNetwork *bool    `url:"is_talent_network"`
	Limit           *int     `url:"limit"`
	Offset          *int     `url:"offset"`
}

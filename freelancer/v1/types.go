package freelancer

import "fmt"

type ProjectType string
type ProjectUpgradeType string
type ContestUpgradeType string
type SortFieldsType string

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
)

type ListProjectsResponse struct {
	Status string         `json:"status"`
	Result ProjectsResult `json:"result"`
}

type ProjectsResult struct {
	Projects   []Project `json:"projects"`
	TotalCount int       `json:"total_count"`
}

type ListUsersResponse struct {
	Status string      `json:"status"`
	Result UsersResult `json:"result"`
}

type UsersResult struct {
	Users map[string]User `json:"users"`
}

type Country struct {
	Name         string  `json:"name,omitempty"`          // Optional
	Code         string  `json:"code,omitempty"`          // Optional
	ISO3         string  `json:"iso3,omitempty"`          // Optional
	PhoneCode    float64 `json:"phone_code,omitempty"`    // Optional, "Decimal"
	Demonym      string  `json:"demonym,omitempty"`       // Optional
	Person       string  `json:"person,omitempty"`        // Optional
	SEOURL       string  `json:"seo_url,omitempty"`       // Optional
	Sanction     bool    `json:"sanction,omitempty"`      // Optional, Boolean
	LanguageCode string  `json:"language_code,omitempty"` // Optional
	LanguageID   float64 `json:"language_id,omitempty"`   // Optional
}

type Jobs struct {
	// Define any properties of "jobs" here if needed.
	// Leaving it empty as per the schema.
}

type Category struct {
	ID   int    `json:"id"`             // "number" is represented as float64 in Go
	Name string `json:"name,omitempty"` // Optional
}

type Currency struct {
	ID           int     `json:"id"`
	Code         string  `json:"code"`
	Sign         string  `json:"sign"`
	Name         string  `json:"name"`
	ExchangeRate float64 `json:"exchange_rate"`
	Country      string  `json:"country"`
	IsExternal   bool    `json:"is_external"`
}

type Budget struct {
	Minimum     float64     `json:"minimum"`
	Maximum     float64     `json:"maximum"`
	Name        string      `json:"name"`
	ProjectType ProjectType `json:"project_type"`
	CurrencyID  int         `json:"currency_id"`
}
type Upgrades struct {
	Featured          bool `json:"featured"`
	Sealed            bool `json:"sealed"`
	Nonpublic         bool `json:"nonpublic"`
	Fulltime          bool `json:"fulltime"`
	Urgent            bool `json:"urgent"`
	Qualified         bool `json:"qualified"`
	NDA               bool `json:"NDA"`
	IpContract        bool `json:"ip_contract"`
	NonCompete        bool `json:"non_compete"`
	ProjectManagement bool `json:"project_management"`
	PfOnly            bool `json:"pf_only"`
	Premium           bool `json:"premium"`
	Enterprise        bool `json:"enterprise"`
}
type CategoryRatings struct {
	Quality         float64 `json:"quality,omitempty"`
	Communication   float64 `json:"communication,omitempty"`
	Expertise       float64 `json:"expertise,omitempty"`
	Professionalism float64 `json:"professionalism,omitempty"`
	HireAgain       float64 `json:"hire_again,omitempty"`
}
type Reputation struct {
	Overall                   float64         `json:"overall,omitempty"`
	OnBudget                  float64         `json:"on_budget,omitempty"`
	OnTime                    float64         `json:"on_time,omitempty"`
	Positive                  float64         `json:"positive,omitempty"`
	CategoryRatings           CategoryRatings `json:"category_ratings,omitempty"`
	All                       int             `json:"all,omitempty"`
	Reviews                   int             `json:"reviews,omitempty"`
	IncompleteReviews         int             `json:"incomplete_reviews,omitempty"`
	Complete                  int             `json:"complete,omitempty"`
	Incomplete                int             `json:"incomplete,omitempty"`
	Earnings                  *float64        `json:"earnings,omitempty"`
	CompletionRate            float64         `json:"completion_rate,omitempty"`
	RehireRate                *float64        `json:"rehire_rate,omitempty"`
	UserID                    int             `json:"user_id,omitempty"`
	CompletedRelevantJobCount *int            `json:"completed_relevant_job_count,omitempty"`
	JobRanks                  *int            `json:"job_ranks,omitempty"`
	RehireRank                *int            `json:"rehire_rank,omitempty"`
}

type Status struct {
	PaymentVerified        bool `json:"payment_verified"`
	EmailVerified          bool `json:"email_verified"`
	DepositMade            bool `json:"deposit_made"`
	ProfileComplete        bool `json:"profile_complete"`
	PhoneVerified          bool `json:"phone_verified"`
	IdentityVerified       bool `json:"identity_verified"`
	FacebookConnected      bool `json:"facebook_connected"`
	FreelancerVerifiedUser bool `json:"freelancer_verified_user"`
	LinkedinConnected      bool `json:"linkedin_connected"`
	CustomChargeVerified   bool `json:"custom_charge_verified"`
}

type PrimaryCurrency struct {
	ID                   int     `json:"id,omitempty"`
	Code                 string  `json:"code,omitempty"`
	Sign                 string  `json:"sign,omitempty"`
	Name                 string  `json:"name,omitempty"`
	ExchangeRate         float64 `json:"exchange_rate,omitempty"`
	Country              string  `json:"country,omitempty"`
	IsExternal           *bool   `json:"is_external,omitempty"`
	IsEscrowcomSupported *bool   `json:"is_escrowcom_supported,omitempty"`
}

type EmployerReputation struct {
	Overall                   float64         `json:"overall,omitempty"`
	OnBudget                  *float64        `json:"on_budget,omitempty"`
	OnTime                    *float64        `json:"on_time,omitempty"`
	Positive                  float64         `json:"positive,omitempty"`
	CategoryRatings           CategoryRatings `json:"category_ratings,omitempty"`
	All                       int             `json:"all,omitempty"`
	Reviews                   int             `json:"reviews,omitempty"`
	IncompleteReviews         *int            `json:"incomplete_reviews,omitempty"`
	Complete                  int             `json:"complete,omitempty"`
	Incomplete                int             `json:"incomplete,omitempty"`
	Earnings                  *float64        `json:"earnings,omitempty"`
	CompletionRate            float64         `json:"completion_rate,omitempty"`
	RehireRate                *float64        `json:"rehire_rate,omitempty"`
	UserID                    int             `json:"user_id,omitempty"`
	CompletedRelevantJobCount *int            `json:"completed_relevant_job_count,omitempty"`
	JobRanks                  *int            `json:"job_ranks,omitempty"`
	RehireRank                *int            `json:"rehire_rank,omitempty"`
}

type Location struct {
	Country struct {
		Name              string `json:"name,omitempty"`
		FlagURL           string `json:"flag_url,omitempty"`
		Code              string `json:"code,omitempty"`
		HighresFlagURL    string `json:"highres_flag_url,omitempty"`
		FlagURLCDN        string `json:"flag_url_cdn,omitempty"`
		HighresFlagURLCDN string `json:"highres_flag_url_cdn,omitempty"`
	} `json:"country,omitempty"`
	City                   string   `json:"city,omitempty"`
	Latitude               *float64 `json:"latitude,omitempty"`
	Longitude              *float64 `json:"longitude,omitempty"`
	Vicinity               *string  `json:"vicinity,omitempty"`
	AdministrativeArea     *string  `json:"administrative_area,omitempty"`
	FullAddress            *string  `json:"full_address,omitempty"`
	AdministrativeAreaCode *string  `json:"administrative_area_code,omitempty"`
	PostalCode             *string  `json:"postal_code,omitempty"`
	ID                     *int     `json:"id,omitempty"`
}

type User struct {
	ID                 int                `json:"id,omitempty"`
	Username           string             `json:"username,omitempty"`
	Suspended          *bool              `json:"suspended,omitempty"`
	Closed             bool               `json:"closed,omitempty"`
	IsActive           *bool              `json:"is_active,omitempty"`
	ForceVerify        *bool              `json:"force_verify,omitempty"`
	Avatar             string             `json:"avatar,omitempty"`
	Email              *string            `json:"email,omitempty"`
	Reputation         Reputation         `json:"reputation,omitempty"`
	HourlyRate         *float64           `json:"hourly_rate,omitempty"`
	RegistrationDate   int64              `json:"registration_date,omitempty"`
	LimitedAccount     bool               `json:"limited_account,omitempty"`
	DisplayName        string             `json:"display_name,omitempty"`
	Tagline            *string            `json:"tagline,omitempty"`
	Location           Location           `json:"location,omitempty"`
	AvatarLarge        string             `json:"avatar_large,omitempty"`
	Role               string             `json:"role,omitempty"`
	ChosenRole         string             `json:"chosen_role,omitempty"`
	EmployerReputation EmployerReputation `json:"employer_reputation,omitempty"`
	Status             Status             `json:"status,omitempty"`
	AvatarCDN          string             `json:"avatar_cdn,omitempty"`
	AvatarLargeCDN     string             `json:"avatar_large_cdn,omitempty"`
	PrimaryCurrency    PrimaryCurrency    `json:"primary_currency,omitempty"`
	PrimaryLanguage    string             `json:"primary_language,omitempty"`
	PublicName         string             `json:"public_name,omitempty"`
}
type Project struct {
	ID             int      `json:"id"`
	OwnerID        int      `json:"owner_id"`
	Title          string   `json:"title"`
	Status         string   `json:"status"`
	SeoURL         string   `json:"seo_url"`
	Currency       Currency `json:"currency"`
	Description    string   `json:"description"`
	SubmitDate     int64    `json:"submitdate"`
	PreviewDesc    string   `json:"preview_description"`
	Deleted        bool     `json:"deleted"`
	Nonpublic      bool     `json:"nonpublic"`
	HideBids       bool     `json:"hidebids"`
	Type           string   `json:"type"`
	BidPeriod      int      `json:"bidperiod"`
	Budget         Budget   `json:"budget"`
	Featured       bool     `json:"featured"`
	Urgent         bool     `json:"urgent"`
	TimeSubmitted  int64    `json:"time_submitted"`
	TimeUpdated    int64    `json:"time_updated"`
	Upgrades       Upgrades `json:"upgrades"`
	Language       string   `json:"language"`
	FrontendStatus string   `json:"frontend_project_status"`
	Location       Location `json:"location"`
	Local          bool     `json:"local"`
}

type APIError struct {
	Code     int64  `json:"code"`
	Message  string `json:"message"`
	Response []byte `json:"-"`
}

type APIError2 struct {
	Status    string `json:"status"`
	Message   string `json:"message"`
	ErrorCode string `json:"error_code"`
	ErrorObj  error2 `json:"error"`
	RequestID string `json:"request_id"`
	Response  []byte `json:"-"`
}

type error2 struct {
	Code     string `json:"code"`
	Detail   string `json:"detail"`
	HttpCode int    `json:"http_code"`
	Source   string `json:"source"`
}

func (e APIError) Error() string {
	if e.IsValid() {
		return fmt.Sprintf("<APIError> code: %d, message: %s", e.Code, e.Message)
	}
	return fmt.Sprintf("<APIError> response: %s", string(e.Response))
}

func (e APIError2) Error() string {
	return fmt.Sprintf("<APIError2> status: %s, message: %s", e.Status, e.Message)
}

func (e APIError2) IsValid() bool {
	return e.Status != "" || e.Message != ""
}

func (e APIError) IsValid() bool {
	return e.Code != 0 || e.Message != ""
}

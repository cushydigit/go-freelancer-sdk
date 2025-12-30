package freelancer

import (
	"fmt"

	"github.com/cushydigit/go-freelancer-sdk/freelancer/internal/endpoints"
)

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

type Job struct {
	ID                 int64     `json:"id"`
	Name               *string   `json:"name,omitempty"`
	Category           *Category `json:"category,omitempty"`
	ActiveProjectCount float64   `json:"active_project_count"`
	SeoURL             *string   `json:"seo_url,omitempty"`
	Local              *bool     `json:"local,omitempty"`
}

type JobHistory struct {
	CountOther int64      `json:"count_other"`
	JobCounts  []JobCount `json:"job_counts"`
}

type JobCount struct {
	Count int `json:"count"`
	Job   Job `json:"job"`
}

type Category struct {
	ID   int    `json:"id"`             // "number" is represented as float64 in Go
	Name string `json:"name,omitempty"` // Optional
}
type CategoryRating struct {
	Communication   float64 `json:"communication"`
	Expertise       float64 `json:"expertise"`
	HireAgain       float64 `json:"hire_again"`
	Quality         float64 `json:"quality"`
	Professionalism float64 `json:"professionalism"`
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
	EntireHistory *ReputationStats `json:"entire_history,omitempty"`
	Last3Months   *ReputationStats `json:"last3months,omitempty"`
	Last12Months  *ReputationStats `json:"last12months,omitempty"`
	UserID        *int64           `json:"user_id,omitempty"`
	Role          *string          `json:"role,omitempty"` // "freelancer" or "employer"
	EarningsScore *float64         `json:"earnings_score,omitempty"`
	JobHistory    *JobHistory      `json:"job_history,omitempty"`
}

type ReputationStats struct {
	Overall                   float64          `json:"overall,omitempty"`
	OnBudget                  float64          `json:"on_budget,omitempty"`
	OnTime                    float64          `json:"on_time,omitempty"`
	Positive                  float64          `json:"positive,omitempty"`
	CategoryRatings           *CategoryRatings `json:"category_ratings,omitempty"`
	All                       int              `json:"all,omitempty"`
	Reviews                   int              `json:"reviews,omitempty"`
	IncompleteReviews         int              `json:"incomplete_reviews,omitempty"`
	Complete                  int              `json:"complete,omitempty"`
	Incomplete                int              `json:"incomplete,omitempty"`
	Earnings                  *float64         `json:"earnings,omitempty"`
	CompletionRate            float64          `json:"completion_rate,omitempty"`
	RehireRate                *float64         `json:"rehire_rate,omitempty"`
	UserID                    int64            `json:"user_id,omitempty"`
	CompletedRelevantJobCount *int             `json:"completed_relevant_job_count,omitempty"`
	JobRanks                  *int             `json:"job_ranks,omitempty"`
	RehireRank                *int             `json:"rehire_rank,omitempty"`
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
	EntireHistory *ReputationStats `json:"entire_history,omitempty"`
	Last3Months   *ReputationStats `json:"last3months,omitempty"`
	Last12Months  *ReputationStats `json:"last12months,omitempty"`
	UserID        *float64         `json:"user_id,omitempty"`
	Role          *string          `json:"role,omitempty"` // "freelancer" or "employer"
	EarningsScore *float64         `json:"earnings_score,omitempty"`
	JobHistory    *JobHistory      `json:"job_history,omitempty"`
}

type Location struct {
	Country                *CountryLocation `json:"country,omitempty"`
	City                   string           `json:"city,omitempty"`
	Latitude               *float64         `json:"latitude,omitempty"`
	Longitude              *float64         `json:"longitude,omitempty"`
	Vicinity               *string          `json:"vicinity,omitempty"`
	AdministrativeArea     *string          `json:"administrative_area,omitempty"`
	FullAddress            *string          `json:"full_address,omitempty"`
	AdministrativeAreaCode *string          `json:"administrative_area_code,omitempty"`
	PostalCode             *string          `json:"postal_code,omitempty"`
	ID                     *int             `json:"id,omitempty"`
}

type CountryLocation struct {
	Name              string `json:"name,omitempty"`
	FlagURL           string `json:"flag_url,omitempty"`
	Code              string `json:"code,omitempty"`
	HighresFlagURL    string `json:"highres_flag_url,omitempty"`
	FlagURLCDN        string `json:"flag_url_cdn,omitempty"`
	HighresFlagURLCDN string `json:"highres_flag_url_cdn,omitempty"`
}

type User struct {
	ID                 int64               `json:"id,omitempty"`
	Username           string              `json:"username,omitempty"`
	Suspended          *bool               `json:"suspended,omitempty"`
	Closed             bool                `json:"closed,omitempty"`
	IsActive           *bool               `json:"is_active,omitempty"`
	ForceVerify        *bool               `json:"force_verify,omitempty"`
	Avatar             string              `json:"avatar,omitempty"`
	Email              *string             `json:"email,omitempty"`
	Reputation         *Reputation         `json:"reputation,omitempty"`
	Jobs               []Job               `json:"jobs,omitempty"`
	ProfileDescription *string             `json:"profile_description"`
	HourlyRate         *float64            `json:"hourly_rate,omitempty"`
	RegistrationDate   int64               `json:"registration_date,omitempty"`
	LimitedAccount     bool                `json:"limited_account,omitempty"`
	DisplayName        string              `json:"display_name,omitempty"`
	Tagline            *string             `json:"tagline,omitempty"`
	Location           *Location           `json:"location,omitempty"`
	AvatarLarge        string              `json:"avatar_large,omitempty"`
	Role               string              `json:"role,omitempty"`
	ChosenRole         string              `json:"chosen_role,omitempty"`
	EmployerReputation *EmployerReputation `json:"employer_reputation,omitempty"`
	Status             *Status             `json:"status,omitempty"`
	AvatarCDN          string              `json:"avatar_cdn,omitempty"`
	AvatarLargeCDN     string              `json:"avatar_large_cdn,omitempty"`
	PrimaryCurrency    *PrimaryCurrency    `json:"primary_currency,omitempty"`
	PrimaryLanguage    string              `json:"primary_language,omitempty"`
	PublicName         string              `json:"public_name,omitempty"`
}

type Project struct {
	ID             int64    `json:"id"`
	OwnerID        int64    `json:"owner_id"`
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

func (p *Project) GetFullUrl() string {
	return fmt.Sprintf("%s/%s", endpoints.BaseAPIMainURL, p.SeoURL)
}

type Timezone struct {
	ID       int     `json:"id"`                 // "number" represented as float64
	Country  string  `json:"country,omitempty"`  // Optional
	Timezone string  `json:"timezone,omitempty"` // Optional
	Offset   float32 `json:"offset,omitempty"`
}
type Device struct {
	UserAgent string `json:"user_agent"`
	Platform  string `json:"platform"`
	City      string `json:"city"`
	Country   string `json:"country"`
	LastLogin int64  `json:"last_login"`
}

type BidEditRequest struct {
	ID            int     `json:"id"`
	BidID         int     `json:"bid_id"`
	Status        string  `json:"status"` // "pending", "accepted", "declined"
	Comment       string  `json:"comment,omitempty"`
	NewAmount     float64 `json:"new_amount"`
	NewPeriod     int     `json:"new_period"`
	OldAmount     float64 `json:"old_amount"`
	OldPeriod     int     `json:"old_period"`
	TimeRequested int64   `json:"time_requested"`           // Unix timestamp
	TimeResponded *int64  `json:"time_responded,omitempty"` // optional, Unix timestamp
}

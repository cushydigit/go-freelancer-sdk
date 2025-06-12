package freelancer

type Currency struct {
	ID float64 `json:"id"`
}

type Budget struct {
	Minimum    float64  `json:"minimum"`
	Maximum    *float64 `json:"maximum,omitempty"`
	CurrencyID *float64 `json:"currency_id,omitempty"`
}

type HourlyInfo struct {
	Commitment Commitment `json:"commitment"`
}

type Commitment struct {
	Hours    float64 `json:"hours"`
	Interval string  `json:"interval"` // "WEEK" or "MONTH"
}

type HiremeInitialBid struct {
	BidderID float64 `json:"bidder_id"`
	Amount   float64 `json:"amount"`
	Period   float64 `json:"period"`
}

type Location struct {
	City               *string  `json:"city,omitempty"`
	Country            *Country `json:"country,omitempty"`
	Latitude           *float64 `json:"latitude,omitempty"`
	Longitude          *float64 `json:"longitude,omitempty"`
	AdministrativeArea *string  `json:"administrative_area,omitempty"`
	FullAddress        *string  `json:"full_address,omitempty"`
	Vicinity           *string  `json:"vicinity,omitempty"`
}

type Country struct {
	Name              string  `json:"name"`
	FlagURL           *string `json:"flag_url"`
	Code              *string `json:"code"`
	HighresFlagURL    *string `json:"highres_flag_url"`
	FlagURLCDN        *string `json:"flag_url_cdn"`
	HighresFlagURLCDN *string `json:"highres_flag_url_cdn"`
}

type LocalDetails struct {
	ID          float64      `json:"id"`
	ProjectID   float64      `json:"project_id"`
	Date        *ProjectDate `json:"date,omitempty"`
	EndLocation *EndLocation `json:"end_location,omitempty"`
}

type ProjectDate struct {
	Year  int `json:"year"`
	Month int `json:"month"`
	Day   int `json:"day"`
}

type EndLocation struct {
	Country            *Country `json:"country,omitempty"`
	City               *string  `json:"city,omitempty"`
	Latitude           *float64 `json:"latitude"`
	Longitude          *float64 `json:"longitude"`
	Vicinity           *string  `json:"vicinity"`
	AdministrativeArea *string  `json:"administritive_area"`
	FullAddress        *string  `json:"full_address"`
}

type Equipment struct {
	ID   float64 `json:"id"`
	Name string  `json:"name"`
}

type EquipmentGroup struct {
	ID                float64       `json:"id"`
	Name              string        `json:"name"`
	HasUserInput      bool          `json:"has_user_input"`
	IsProjectRequired bool          `json:"is_project_required"`
	IsUserRequired    bool          `json:"is_user_required"`
	Items             []interface{} `json:"items,omitempty"` // You can define a custom Equipment type here if items have a structure
}

package freelancer

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// TODO: refine with typed response
type listExpertGuarantees struct {
	client           *Client
	expertGuarantees []int64
	projects         []int64
	projectOwners    []int64
	bidders          []int64
	bids             []int
	statuses         []ExpertGuaranteesStatus
	fromTime         *int64
	toTime           *int64
	offset           *int
	limit            *int
}

func (s *listExpertGuarantees) Do(ctx context.Context) (*RawResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: string(PROJECTS_EXPERT_GUARANTEES),
	}

	for _, val := range s.bids {
		r.addParam("bids[]", val)
	}
	for _, val := range s.projects {
		r.addParam("projects[]", val)
	}
	for _, val := range s.bidders {
		r.addParam("bidders[]", val)
	}
	for _, val := range s.projectOwners {
		r.addParam("project_owners[]", val)
	}
	for _, val := range s.bids {
		r.addParam("bids[]", val)
	}
	for _, val := range s.statuses {
		r.addParam("statuses[]", val)
	}
	if s.fromTime != nil {
		r.addParam("from_time", *s.fromTime)
	}
	if s.toTime != nil {
		r.addParam("to_time", *s.toTime)
	}
	if s.limit != nil {
		r.addParam("limit", *s.limit)
	}
	if s.offset != nil {
		r.addParam("offset", *s.offset)
	}
	return execute[*RawResponse](ctx, s.client, r)
}

// SetExpertGuarantees sets a list of expert guarantee IDs to filter the results.
// These IDs correspond to specific expert guarantees you want to retrieve.
func (s *listExpertGuarantees) SetExpertGuarantees(vals []int64) *listExpertGuarantees {
	s.expertGuarantees = vals
	return s
}

// SetBids sets a list of bid IDs to filter expert guarantees associated with those bids.
func (s *listExpertGuarantees) SetBids(vals []int) *listExpertGuarantees {
	s.bids = vals
	return s
}

// SetProjects sets a list of project IDs to filter expert guarantees linked to those projects.
func (s *listExpertGuarantees) SetProjects(vals []int64) *listExpertGuarantees {
	s.projects = vals
	return s
}

// SetBidders sets a list of freelancer user IDs (bidders) to filter expert guarantees.
func (s *listExpertGuarantees) SetBidders(vals []int64) *listExpertGuarantees {
	s.bidders = vals
	return s
}

// SetProjectOwners sets a list of employer user IDs (project owners) to filter expert guarantees.
func (s *listExpertGuarantees) SetProjectOwners(vals []int64) *listExpertGuarantees {
	s.projectOwners = vals
	return s
}

// SetStatuses sets the list of statuses to filter expert guarantees.
// Valid statuses: disputed, requested_release, locked, released, liquidated, admin_refunded
func (s *listExpertGuarantees) SetStatuses(vals []ExpertGuaranteesStatus) *listExpertGuarantees {
	s.statuses = vals
	return s
}

// SetFromTime sets the start time (inclusive) to filter expert guarantees made after this timestamp.
func (s *listExpertGuarantees) SetFromTime(val int64) *listExpertGuarantees {
	s.fromTime = &val
	return s
}

// SetToTime sets the end time (exclusive) to filter expert guarantees made before this timestamp.
func (s *listExpertGuarantees) SetToTime(val int64) *listExpertGuarantees {
	s.toTime = &val
	return s
}

// SetLimit sets the maximum number of expert guarantees to return.
func (s *listExpertGuarantees) SetLimit(val int) *listExpertGuarantees {
	s.limit = &val
	return s
}

// SetOffset sets the offset for paginated results of expert guarantees.
func (s *listExpertGuarantees) SetOffset(val int) *listExpertGuarantees {
	s.offset = &val
	return s
}

type actionExpertGuarantee struct {
	client             *Client
	expertGuaranteesID int64
}

type ExpertGuaranteesActionRequestBody struct {
	Action ExpertGuaranteesAction `json:"action"`
}

func (s *actionExpertGuarantee) Do(ctx context.Context, b ExpertGuaranteesActionRequestBody) (*RawResponse, error) {
	m, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodPut,
		endpoint: fmt.Sprintf("%s/%s", string(PROJECTS_EXPERT_GUARANTEES), strconv.FormatInt(s.expertGuaranteesID, 10)),
		body:     bytes.NewBuffer(m),
	}
	return execute[*RawResponse](ctx, s.client, r)
}

type listCurrencies struct {
	client                    *Client
	currencyCodes             []string
	currencyIDs               []int
	includeExternalCurrencies *bool
}

func (s *listCurrencies) Do(ctx context.Context) (*ListCurrenciesResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/projects/0.1/currencies",
	}
	if s.currencyCodes != nil {
		r.addParam("currency_codes", s.currencyCodes)
	}
	if s.currencyIDs != nil {
		r.addParam("currency_ids", s.currencyIDs)
	}
	if s.includeExternalCurrencies != nil {
		r.addParam("include_external_currencies", s.includeExternalCurrencies)
	}
	return execute[*ListCurrenciesResponse](ctx, s.client, r)
}

// SetCurrencyCodes sets the filter to return only currencies with the specified currency codes.
// Example: ["AUD", "USD"]. Note that currency_codes and currency_ids are incompatible and should not be used together.
func (s *listCurrencies) SetCurrencyCodes(vals []string) *listCurrencies {
	s.currencyCodes = vals
	return s
}

// SetCurrencyIDs sets the filter to return only currencies with the specified currency IDs.
// Example: [1, 2]. Note that currency_ids and currency_codes are incompatible and should not be used together.
func (s *listCurrencies) SetCurrencyIDs(vals []int) *listCurrencies {
	s.currencyIDs = vals
	return s
}

// SetIncludeExternalCurrencies sets whether to include external currencies in the returned list.
// Example: true
func (s *listCurrencies) SetIncludeExternalCurrencies(val bool) *listCurrencies {
	s.includeExternalCurrencies = &val
	return s
}

type listCategories struct {
	client                    *Client
	categories                []int
	jobDetails                bool
	lang                      string
	activeProjectCountDetails bool
	seoDetails                bool
}

func (s *listCategories) Do(ctx context.Context) (*ListCategoriesResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: string(PROJECTS_CATEGORIES),
	}
	if s.categories != nil {
		r.addParam("categories", s.categories)
	}
	if s.jobDetails {
		r.addParam("job_details", s.jobDetails)
	}
	if s.lang != "" {
		r.addParam("lang", s.lang)
	}
	if s.activeProjectCountDetails {
		r.addParam("active_project_count_details", s.activeProjectCountDetails)
	}
	if s.seoDetails {
		r.addParam("seo_details", s.seoDetails)
	}
	return execute[*ListCategoriesResponse](ctx, s.client, r)
}

// SetCategories sets the category IDs to filter the list of categories.
// Example: []int{1, 2}.
// Optional: If not set, all categories are returned.
func (s *listCategories) SetCategories(categories []int) *listCategories {
	s.categories = categories
	return s
}

// SetJobDetails determines whether to include a map of category IDs to jobs in those categories.
// Optional: Default is false.
func (s *listCategories) SetJobDetails(jobDetails bool) *listCategories {
	s.jobDetails = jobDetails
	return s
}

// SetLang sets the language code for the categories returned.
// Example: "en".
// Optional: Default is the service's default language.
func (s *listCategories) SetLang(lang string) *listCategories {
	s.lang = lang
	return s
}

// SetActiveProjectCountDetails determines whether to include counts of active projects per job type.
// Optional: Default is false. Counts may be up to 30 seconds out of date.
func (s *listCategories) SetActiveProjectCountDetails(activeProjectCountDetails bool) *listCategories {
	s.activeProjectCountDetails = activeProjectCountDetails
	return s
}

// SetSeoDetails determines whether to include SEO information for jobs in each category.
// Optional: Default is false.
func (s *listCategories) SetSeoDetails(seoDetails bool) *listCategories {
	s.seoDetails = seoDetails
	return s
}

type listBudgets struct {
	client          *Client
	currencyCodes   []string
	currencyIDs     []int
	projectType     ProjectType
	lang            string
	currencyDetails bool
}

func (s *listBudgets) Do(ctx context.Context) (*ListBudgetsResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/projects/0.1/budgets",
	}
	if s.currencyCodes != nil {
		r.addParam("currency_codes", s.currencyCodes)
	}
	if s.currencyIDs != nil {
		r.addParam("currency_ids", s.currencyIDs)
	}
	if s.projectType != "" {
		r.addParam("project_type", s.projectType)
	}
	if s.lang != "" {
		r.addParam("lang", s.lang)
	}
	if s.currencyDetails {
		r.addParam("currency_details", s.currencyDetails)
	}
	return execute[*ListBudgetsResponse](ctx, s.client, r)
}

// SetCurrencyCodes sets a list of currency codes to filter the budgets.
// Only budgets with the specified currency codes will be returned.
// Example: ["AUD", "USD"]
// Note: currency_codes and currency_ids are incompatible with each other.
func (s *listBudgets) SetCurrencyCodes(currencyCodes []string) *listBudgets {
	s.currencyCodes = currencyCodes
	return s
}

// SetCurrencyIDs sets a list of currency IDs to filter the budgets.
// Only budgets with the specified currency IDs will be returned.
// Example: [1, 2]
// Note: currency_ids and currency_codes are incompatible with each other.
func (s *listBudgets) SetCurrencyIDs(currencyIDs []int) *listBudgets {
	s.currencyIDs = currencyIDs
	return s
}

// SetProjectType sets the type of project budgets to filter.
// Possible values: "fixed" or "hourly".
func (s *listBudgets) SetProjectType(projectType ProjectType) *listBudgets {
	s.projectType = projectType
	return s
}

// SetLang sets the language code used to translate the budget projection.
// Example: "en"
func (s *listBudgets) SetLang(lang string) *listBudgets {
	s.lang = lang
	return s
}

// SetCurrencyDetails sets whether to include detailed currency information in the budget projection.
func (s *listBudgets) SetCurrencyDetails(currencyDetails bool) *listBudgets {
	s.currencyDetails = currencyDetails
	return s
}

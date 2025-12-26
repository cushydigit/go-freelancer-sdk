package freelancer

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

type listExpertGuaranteesOptions struct {
	ExpertGuarantees []int64
	Projects         []int64
	ProjectOwners    []int64
	Bidders          []int64
	Bids             []int64
	Statuses         []ExpertGuaranteesStatus
	FromTime         *int64
	ToTime           *int64
	Offset           *int
	Limit            *int
}

// TODO: refine with typed response

// Returns a list of expert guarantees.
// It maps to the `GET` `/projects/0.1/expert_guarantees` endpoint
func (s *ExpertGuaranteesService) List(ctx context.Context, opts *listExpertGuaranteesOptions) (*RawResponse, error) {
	query := url.Values{}
	if opts != nil {
		for _, val := range opts.ExpertGuarantees {
			addInt64(query, "expert_guarantees[]", &val)
		}
		for _, val := range opts.Projects {
			addInt64(query, "projects[]", &val)
		}
		for _, val := range opts.ProjectOwners {
			addInt64(query, "project_owners[]", &val)
		}
		for _, val := range opts.Bidders {
			addInt64(query, "bidders[]", &val)
		}
		for _, val := range opts.Bids {
			addInt64(query, "bids[]", &val)
		}
		for _, val := range opts.Statuses {
			addEnum(query, "statuses[]", &val)
		}
		addInt64(query, "from_time", opts.FromTime)
		addInt64(query, "to_time", opts.ToTime)
		addInt(query, "offset", opts.Offset)
		addInt(query, "limit", opts.Limit)
	}
	return execute[*RawResponse](ctx, s.client, http.MethodGet, string(ProjectsExpertGuarantees), query, nil)
}

type ExpertGuaranteesActionRequestBody struct {
	Action ExpertGuaranteesAction `json:"action"`
}

// Perform an action on a expert guarantee.
// It maps to the `PUT` `/projects/0.1/expert_guarantees/{expert_guarantee_id}` endpoint
func (s *ExpertGuaranteesService) Action(ctx context.Context, expertGuaranteesID int64, b ExpertGuaranteesActionRequestBody) (*RawResponse, error) {
	path := fmt.Sprintf("%s/%d", ProjectsExpertGuarantees, expertGuaranteesID)
	return execute[*RawResponse](ctx, s.client, http.MethodPut, path, nil, b)
}

type ListCurrenciesOptions struct {
	CurrencyCodes             []string
	CurrencyIDs               []int
	IncludeExternalCurrencies *bool
}

// Returns a list of currencies.currency_codes and currency_ids are incompatible with each other.
// It maps to the `GET` `/projects/0.1/currencies` endpoint
func (s *CurrenciesService) List(ctx context.Context, opts *ListCurrenciesOptions) (*ListCurrenciesResponse, error) {
	query := url.Values{}
	if opts != nil {
		for _, val := range opts.CurrencyCodes {
			addString(query, "currency_codes[]", &val)
		}
		for _, val := range opts.CurrencyIDs {
			addInt(query, "currency_ids[]", &val)
		}
		addBool(query, "include_external_currencies", opts.IncludeExternalCurrencies)
	}
	return execute[*ListCurrenciesResponse](ctx, s.client, http.MethodGet, string(ProjectsCurrencies), query, nil)
}

type ListCategoriesOptions struct {
	Categories                []int
	JobDetails                *bool
	Lang                      *string
	ActiveProjectCountDetails *bool
	SeoDetails                *bool
}

// Returns a list of categories. If job_details is set, a map of category IDs to jobs in those categories.
// it maps to the `GET` `/projects/0.1/categories` endpoint
func (s *CategoriesService) List(ctx context.Context, opts *ListCategoriesOptions) (*ListCategoriesResponse, error) {
	query := url.Values{}
	if opts != nil {
		for _, val := range opts.Categories {
			addInt(query, "categories[]", &val)
		}
		addBool(query, "job_details", opts.JobDetails)
		addString(query, "lang", opts.Lang)
		addBool(query, "active_project_count_details", opts.ActiveProjectCountDetails)
		addBool(query, "seo_details", opts.SeoDetails)
	}
	return execute[*ListCategoriesResponse](ctx, s.client, http.MethodGet, string(ProjectsCategories), query, nil)
}

type ListBudgetsOptions struct {
	CurrencyCodes   []string
	CurrencyIDs     []int
	ProjectType     *ProjectType
	Lang            *string
	CurrencyDetails *bool
}

// Returns a list of budgets with the specified currencies.currency_codesandcurrency_idsare incompatible with each other.
// It maps to the `GET` `/projects/0.1/budgets` endpoint
func (s *BudgetsService) List(ctx context.Context, opts *ListBudgetsOptions) (*ListBudgetsResponse, error) {
	query := url.Values{}
	if opts != nil {
		for _, val := range opts.CurrencyCodes {
			addString(query, "currency_codes[]", &val)
		}
		for _, val := range opts.CurrencyIDs {
			addInt(query, "currency_ids[]", &val)
		}
		addEnum(query, "project_type", opts.ProjectType)
		addString(query, "lang", opts.Lang)
		addBool(query, "currency_details", opts.CurrencyDetails)
	}

	return execute[*ListBudgetsResponse](ctx, s.client, http.MethodGet, string(ProjectsBudgets), nil, opts)
}

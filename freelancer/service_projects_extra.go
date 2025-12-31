package freelancer

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cushydigit/go-freelancer-sdk/freelancer/internal/endpoints"
	"github.com/cushydigit/go-freelancer-sdk/freelancer/internal/query"
)

type listExpertGuaranteesOptions struct {
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

// TODO: refine with typed response

// Returns a list of expert guarantees.
// It maps to the `GET` `/projects/0.1/expert_guarantees` endpoint
func (s *ExpertGuaranteesService) List(ctx context.Context, opts *listExpertGuaranteesOptions) (*RawResponse, error) {
	p := endpoints.ProjectsExpertGuarantees
	q := query.Values(opts)
	return execute[*RawResponse](ctx, s.client, http.MethodGet, p, q, nil)
}

type ExpertGuaranteesActionRequestBody struct {
	Action ExpertGuaranteesAction `json:"action"`
}

// Perform an action on a expert guarantee.
// It maps to the `PUT` `/projects/0.1/expert_guarantees/{expert_guarantee_id}` endpoint
func (s *ExpertGuaranteesService) Action(ctx context.Context, expertGuaranteesID int64, b ExpertGuaranteesActionRequestBody) (*RawResponse, error) {
	p := fmt.Sprintf("%s/%d", endpoints.ProjectsExpertGuarantees, expertGuaranteesID)
	return execute[*RawResponse](ctx, s.client, http.MethodPut, p, nil, b)
}

type ListCurrenciesOptions struct {
	CurrencyCodes             []string `url:"currency_codes[]"`
	CurrencyIDs               []int64  `url:"currency_ids[]"`
	IncludeExternalCurrencies *bool    `url:"include_external_currencies"`
}

// Returns a list of currencies.currency_codes and currency_ids are incompatible with each other.
// It maps to the `GET` `/projects/0.1/currencies` endpoint
func (s *CurrenciesService) List(ctx context.Context, opts *ListCurrenciesOptions) (*ListCurrenciesResponse, error) {
	p := endpoints.ProjectsCurrencies
	q := query.Values(opts)
	return execute[*ListCurrenciesResponse](ctx, s.client, http.MethodGet, p, q, nil)
}

type ListCategoriesOptions struct {
	Categories                []int64 `url:"categories[]"`
	JobDetails                *bool   `url:"job_details"`
	Lang                      *string `url:"lang"`
	ActiveProjectCountDetails *bool   `url:"active_project_count_details"`
	SeoDetails                *bool   `url:"seo_details"`
}

// Returns a list of categories. If job_details is set, a map of category IDs to jobs in those categories.
// it maps to the `GET` `/projects/0.1/categories` endpoint
func (s *CategoriesService) List(ctx context.Context, opts *ListCategoriesOptions) (*ListCategoriesResponse, error) {
	p := endpoints.ProjectsCategories
	q := query.Values(opts)
	return execute[*ListCategoriesResponse](ctx, s.client, http.MethodGet, p, q, nil)
}

type ListBudgetsOptions struct {
	CurrencyCodes   []string     `url:"currency_codes[]"`
	CurrencyIDs     []int64      `url:"currency_ids[]"`
	ProjectType     *ProjectType `url:"project_type"`
	Lang            *string      `url:"lang"`
	CurrencyDetails *bool        `url:"currency_details"`
}

// Returns a list of budgets with the specified currencies.currency_codes and currency_ids are incompatible with each other.
// It maps to the `GET` `/projects/0.1/budgets` endpoint
func (s *BudgetsService) List(ctx context.Context, opts *ListBudgetsOptions) (*ListBudgetsResponse, error) {
	p := endpoints.ProjectsBudgets
	q := query.Values(opts)
	return execute[*ListBudgetsResponse](ctx, s.client, http.MethodGet, p, q, nil)
}

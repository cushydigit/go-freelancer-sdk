package usecases

import (
	"context"
	"fmt"

	"github.com/cushydigit/go-freelancer-sdk/freelancer"

	rr "github.com/cushydigit/go-freelancer-sdk/freelancer/reqres"
)

func FetchAndDisplayBudget(ctx context.Context, c *freelancer.Client) {
	res, _, err := c.Resources.Common.ListBudgets(
		ctx,
		&rr.ListBudgetsOptions{
			CurrencyIDs: []int64{
				*rr.Int64(1), // USD
				*rr.Int64(2), // AUD
			},
			ProjectType: rr.Enum(rr.ProjectBudgetFixed), // fixed budget type
		},
	)
	if len(res.Result.Budgets) > 0 && err == nil {
		for _, b := range res.Result.Budgets {
			fmt.Printf("Name: %s, Max: %0.1f Min: %0.1f\n", b.Name, b.Maximum, b.Minimum)
		}
		fmt.Printf("Fetched %d currencies\n", len(res.Result.Budgets))
	}
}

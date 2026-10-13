package usecases

import (
	"context"
	"fmt"

	"github.com/cushydigit/go-freelancer-sdk/freelancer"
	rr "github.com/cushydigit/go-freelancer-sdk/freelancer/reqres"
)

func ListSpecificProjects(ctx context.Context, c *freelancer.Client) {

	res, _, err := c.Resources.Projects.List(
		ctx,
		&rr.ListProjectsOptions{
			Projects: []int64{
				*rr.Int64(40681584),
				*rr.Int64(40681432),
			},
			UserDetails: rr.Bool(true), // Include user info

		},
	)

	if err == nil && len(res.Result.Projects) > 0 {
		for _, p := range res.Result.Projects {
			budgetString := fmt.Sprintf(
				"[%s%1.f - %s%1.f]",
				p.Currency.Sign,
				p.Budget.Minimum,
				p.Currency.Sign,
				p.Budget.Minimum,
			)
			fmt.Printf("\n-%d: %s %s\n", p.ID, budgetString, p.Title)
		}
		fmt.Printf("Showing %d of %d total projects\n", len(res.Result.Projects), res.Result.TotalCount)
	}
}

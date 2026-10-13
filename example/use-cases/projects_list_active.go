package usecases

import (
	"context"
	"fmt"
	"time"

	"github.com/cushydigit/go-freelancer-sdk/freelancer"
	rr "github.com/cushydigit/go-freelancer-sdk/freelancer/reqres"
)

func ListActiveProjects(ctx context.Context, c *freelancer.Client) {

	from := time.Now().Add(-24 * time.Hour)
	res, _, err := c.Resources.Projects.SearchActive(
		ctx,
		&rr.SearchActiveProjectsOptions{
			Query:           rr.String("Go developer"),
			Limit:           rr.Int(20), // Results per page
			Offset:          rr.Int(0),
			FullDescription: rr.Bool(true), // include

			// Filtering options
			FromTime:    &from,         // Filter projects within 24 hours
			UserDetails: rr.Bool(true), // Include user info
			// Sort
			SortField:   rr.Enum(rr.SortFieldsTimeUpdated),
			ReverseSort: rr.Bool(false), // Newest first

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
			// access unix time easily
			time1, time2, time3 := p.SubmitDateAt(), p.UpdatedAt(), p.SubmittedAt()
			if time1 != nil && time2 != nil && time3 != nil {
				fmt.Printf(
					"SubmitDate: %s\tTimeUpdated: %s\tTimeSubmitted: %s\n",
					time1.Format("Jan 2, 2006 at 15:04"),
					time2.Format("Jan 2, 2006 at 15:04"),
					time3.Format("Jan 2, 2006 at 15:04"),
				)
			}
		}
		fmt.Printf("Showing %d of %d total projects\n", len(res.Result.Projects), res.Result.TotalCount)
	}
}

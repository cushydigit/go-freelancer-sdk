package usecases

import (
	"context"
	"fmt"

	"github.com/cushydigit/go-freelancer-sdk/freelancer"
	rr "github.com/cushydigit/go-freelancer-sdk/freelancer/reqres"
)

func CreateHireMeProject(ctx context.Context, c *freelancer.Client) {

	res, _, err := c.Resources.Projects.Create(
		ctx,
		rr.CreateProjectBody{
			Title:       "Go Developer Needed",
			Description: "Need an experienced Go developer...",
			Budget: rr.Budget{
				Minimum:    10,
				CurrencyID: 1,
			},
			Jobs:  []int64{6, 9},
			HirMe: rr.Bool(true),
			HiremeInitialBid: &rr.HiremeInitialBid{
				BidderID:    100, // Freelancer we want to hire
				Amount:      200,
				Period:      5, // number of days
				Description: "Hello, Are you interested in my project",
			},
		},
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Project created with ID: %d\n", res.Result.ID)
}

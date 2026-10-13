package usecases

import (
	"context"
	"fmt"

	"github.com/cushydigit/go-freelancer-sdk/freelancer"
	rr "github.com/cushydigit/go-freelancer-sdk/freelancer/reqres"
)

func CreateHourlyProject(ctx context.Context, c *freelancer.Client) {

	res, _, err := c.Resources.Projects.Create(
		ctx,
		rr.CreateProjectBody{
			Title:       "Go Developer Needed",
			Description: "Need an experienced Go developer...",
			Budget: rr.Budget{
				Minimum:    10,
				CurrencyID: 1,
			},
			Jobs: []int64{6, 9},
			HourlyProjectInfo: &rr.HourlyProjectInfo{
				Commitment: rr.Commitment{
					Hours:    40,
					Interval: rr.IntervalWeek,
				},
			},
		},
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Project created with ID: %d\n", res.Result.ID)
}

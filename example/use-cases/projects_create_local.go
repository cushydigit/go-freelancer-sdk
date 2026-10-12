package usecases

import (
	"context"
	"fmt"

	"github.com/cushydigit/go-freelancer-sdk/freelancer"
	rr "github.com/cushydigit/go-freelancer-sdk/freelancer/reqres"
)

func CreateLocalProject(ctx context.Context, c *freelancer.Client) {

	res, _, err := c.Resources.Projects.Create(
		ctx,
		rr.CreateProjectBody{
			Title:       "Need assistant on local project",
			Description: "Description",
			Budget: rr.Budget{
				Minimum:    100,
				CurrencyID: 1,
			},
			Jobs: []int64{649}, // must be a local job
			Location: &rr.Location{
				Country: &rr.Country{
					Name: "Australia",
				},
				City:      "Sydney",
				Latitude:  rr.Float64(-33.875461),
				Longitude: rr.Float64(151.201678),
			},
		},
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Local project created with ID: %d\n", res.Result.ID)
}

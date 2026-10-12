package usecases

import (
	"context"
	"fmt"
	"log"

	"github.com/cushydigit/go-freelancer-sdk/freelancer"
	rr "github.com/cushydigit/go-freelancer-sdk/freelancer/reqres"
)

func ListMilestones(ctx context.Context, c *freelancer.Client) {
	res, _, err := c.Resources.Milestones.List(
		ctx,
		&rr.ListMilestonesOptions{
			Projects: []int64{101, 105}, // return milestone requests with list of project ids or
			// Bidders:  []int64{1, 2},     // with specified freelancer IDs
			// Users:    []int64{1, 2},     // with specified user IDs
			// Bids:     []int64{1, 2},     // with the specified bid IDs
			// cannot provide both Bidders Users or project_owners etc
			Statuses: []rr.MilestoneStatus{
				rr.MilestoneStatusCanceled,
				rr.MilestoneStatusFrozen,
				rr.MilestoneStatusRequestedRelease,
			},
		},
	)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("List of milestone: %v", res.Result)
}

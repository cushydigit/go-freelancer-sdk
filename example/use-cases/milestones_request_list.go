package usecases

import (
	"context"
	"fmt"
	"log"

	"github.com/cushydigit/go-freelancer-sdk/freelancer"
	rr "github.com/cushydigit/go-freelancer-sdk/freelancer/reqres"
)

func ListMilestoneRequests(ctx context.Context, c *freelancer.Client) {
	res, _, err := c.Resources.Milestones.ListRequests(
		ctx,
		&rr.ListMilestoneRequestsOptions{
			Projects: []int64{101, 105}, // return milestone requests with list of project ids or
			// Bidders:  []int64{1, 2},     // with specified freelancer IDs
			// Users:    []int64{1, 2},     // with specified user IDs
			// Bids:     []int64{1, 2},     // with the specified bid IDs
			// cannot provide both Bidders Users or project_owners etc
			Statuses: []rr.MilestoneRequestStatus{
				rr.MilestoneRequestStatusCreated,
				rr.MilestoneRequestStatusRejected,
			},
		},
	)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("List of milestone requests: %v", res.Result)
}

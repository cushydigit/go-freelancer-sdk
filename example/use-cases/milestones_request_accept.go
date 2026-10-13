package usecases

import (
	"context"
	"fmt"
	"log"

	"github.com/cushydigit/go-freelancer-sdk/freelancer"
)

func AcceptMilestoneRequest(ctx context.Context, c *freelancer.Client) {

	res, _, err := c.Resources.Milestones.AcceptRequest(ctx, 1)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Milestone request accepted: %v", res.Result)
}

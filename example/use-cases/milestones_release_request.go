package usecases

import (
	"context"
	"fmt"
	"log"

	"github.com/cushydigit/go-freelancer-sdk/freelancer"
)

func RequestReleaseMilestone(ctx context.Context, c *freelancer.Client) {
	res, _, err := c.Resources.Milestones.RequestRelease(ctx, 1)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Milestone request release: %v", res.Result)
}

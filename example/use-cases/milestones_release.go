package usecases

import (
	"context"
	"fmt"
	"log"

	"github.com/cushydigit/go-freelancer-sdk/freelancer"
)

func ReleaseMilestone(ctx context.Context, c *freelancer.Client) {
	res, _, err := c.Resources.Milestones.Release(ctx, 1, 25)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Milestone released: %v", res.Result)
}

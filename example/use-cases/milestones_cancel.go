package usecases

import (
	"context"
	"fmt"
	"log"

	"github.com/cushydigit/go-freelancer-sdk/freelancer"
)

func CancelMilestone(ctx context.Context, c *freelancer.Client) {
	res, _, err := c.Resources.Milestones.Cancel(ctx, 1)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Create milestone: %v", res.Result)
}

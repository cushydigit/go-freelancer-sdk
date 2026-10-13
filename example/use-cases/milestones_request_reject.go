package usecases

import (
	"context"
	"fmt"
	"log"

	"github.com/cushydigit/go-freelancer-sdk/freelancer"
)

func RejectMilestoneRequest(ctx context.Context, c *freelancer.Client) {
	res, _, err := c.Resources.Milestones.RejectRequest(ctx, 1)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Milestone request rejected: %v", res.Result)
}

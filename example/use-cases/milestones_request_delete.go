package usecases

import (
	"context"
	"fmt"
	"log"

	"github.com/cushydigit/go-freelancer-sdk/freelancer"
)

func DeleteMilestoneRequest(ctx context.Context, c *freelancer.Client) {
	res, _, err := c.Resources.Milestones.DeleteRequest(ctx, 1)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Milestone request deleted: %v", res.Result)
}

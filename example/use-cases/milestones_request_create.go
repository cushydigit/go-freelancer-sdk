package usecases

import (
	"context"
	"fmt"
	"log"

	"github.com/cushydigit/go-freelancer-sdk/freelancer"

	rr "github.com/cushydigit/go-freelancer-sdk/freelancer/reqres"
)

func CreateMilestoneRequest(ctx context.Context, c *freelancer.Client) {
	res, _, err := c.Resources.Milestones.CreateRequest(
		ctx,
		rr.CreateMilestoneRequestBody{
			ProjectID:   1,
			BidID:       1,
			Amount:      100,
			Description: "This is a milestone request",
		},
	)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Milestone request created: %v", res.Result)
}

package usecases

import (
	"context"
	"fmt"
	"log"

	"github.com/cushydigit/go-freelancer-sdk/freelancer"
	rr "github.com/cushydigit/go-freelancer-sdk/freelancer/reqres"
)

func CreateMilestone(ctx context.Context, c *freelancer.Client) {
	res, _, err := c.Resources.Milestones.Create(
		ctx,
		rr.CreateMilestoneBody{
			ProjectID:   1,
			BidderID:    1,
			Amount:      50,
			Reason:      rr.MilestoneCreateReasonFullPayment,
			Description: "This is full payment milestone",
		},
	)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Create milestone: %v", res.Result)
}

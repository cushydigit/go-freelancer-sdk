package usecases

import (
	"context"
	"fmt"
	"log"

	"github.com/cushydigit/go-freelancer-sdk/freelancer"
	rr "github.com/cushydigit/go-freelancer-sdk/freelancer/reqres"
)

func GetSpecificMilestone(ctx context.Context, c *freelancer.Client) {
	res, _, err := c.Resources.Milestones.Get(
		ctx,
		1, // milestone id
		&rr.GetMilestoneOptions{
			UserDisplayInfo:        rr.Bool(true),
			UserCountryDetails:     rr.Bool(true),
			UserProfileDescription: rr.Bool(true),
		},
	)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Specific milestone: %v", res.Result)
}

package projects

import (
	"context"
	"fmt"
	"log"

	"github.com/cushydigit/go-freelancer-sdk/freelancer"
	rr "github.com/cushydigit/go-freelancer-sdk/freelancer/reqres"
)

func ListMilestonesExample(ctx context.Context, c *freelancer.Client) {
	res, _, err := c.Resources.Milestones.List(
		ctx,
		&rr.ListMilestonesOptions{
			Projects: []int64{101, 102}, // list of projects ids
			Statuses: []rr.MilestoneStatus{
				rr.MilestoneStatusCanceled,
				rr.MilestoneStatusCreated,
			},
			UserDisplayInfo:        rr.Bool(true),
			UserCountryDetails:     rr.Bool(true),
			UserProfileDescription: rr.Bool(true),
		},
	)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("List of milestones: %v", res.Result)
}

func GetSpecificMilestoneExample(ctx context.Context, c *freelancer.Client) {
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

func CreateMilestoneRequestExample(ctx context.Context, c *freelancer.Client) {
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

func DeleteMilestoneRequestExample(ctx context.Context, c *freelancer.Client) {
	res, _, err := c.Resources.Milestones.DeleteRequest(ctx, 1)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Milestone request deleted: %v", res.Result)
}

func RejectMilestoneRequestExample(ctx context.Context, c *freelancer.Client) {
	res, _, err := c.Resources.Milestones.RejectRequest(ctx, 1)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Milestone request rejected: %v", res.Result)
}

func AcceptMilestoneRequestExample(ctx context.Context, c *freelancer.Client) {

	res, _, err := c.Resources.Milestones.AcceptRequest(ctx, 1)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Milestone request accepted: %v", res.Result)
}

func CreateMilestoneExample(ctx context.Context, c *freelancer.Client) {
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

func RequestReleaseMilestoneExample(ctx context.Context, c *freelancer.Client) {
	res, _, err := c.Resources.Milestones.RequestRelease(ctx, 1)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Milestone request release: %v", res.Result)
}

func ReleaseMilestoneExample(ctx context.Context, c *freelancer.Client) {
	res, _, err := c.Resources.Milestones.Release(ctx, 1, 25)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Milestone released: %v", res.Result)
}

func CancelMilestoneExample(ctx context.Context, c *freelancer.Client) {
	res, _, err := c.Resources.Milestones.Cancel(ctx, 1)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Create milestone: %v", res.Result)
}

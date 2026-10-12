package usecases

import (
	"context"
	"fmt"
	"log"

	"github.com/cushydigit/go-freelancer-sdk/freelancer"
	rr "github.com/cushydigit/go-freelancer-sdk/freelancer/reqres"
)

func PlaceBidOnProject(ctx context.Context, c *freelancer.Client) {
	res, _, err := c.Resources.Bids.Create(
		ctx,
		rr.CreateBidBody{
			ProjectID:           1, //
			BidderID:            2, // user id (your use id)
			Amount:              100,
			Period:              2, // number of days
			MilestonePercentage: 100,
			Description:         "This is my bid",
			ProfileID:           1, // profile id for bid
		},
	)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Bid placed: %v", res.Result)
}

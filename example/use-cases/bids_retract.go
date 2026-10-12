package usecases

import (
	"context"
	"fmt"

	"github.com/cushydigit/go-freelancer-sdk/freelancer"
)

func RetractBidFromProject(ctx context.Context, c *freelancer.Client) {
	res, _, err := c.Resources.Bids.Retract(ctx, 1)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Bid retracted: %v", res.Result)
}

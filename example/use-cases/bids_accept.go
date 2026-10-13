package usecases

import (
	"context"
	"fmt"

	"github.com/cushydigit/go-freelancer-sdk/freelancer"
)

func AcceptBidOnProject(ctx context.Context, c *freelancer.Client) {
	res, _, err := c.Resources.Bids.Accept(ctx, 1)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Bid accepted: %v", res.Result)
}

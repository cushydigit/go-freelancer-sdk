package usecases

import (
	"context"
	"fmt"

	"github.com/cushydigit/go-freelancer-sdk/freelancer"
)

func RevokeBidFromProject(ctx context.Context, c *freelancer.Client) {
	res, _, err := c.Resources.Bids.Revoke(ctx, 1)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Bid revoked: %v", res.Result)
}

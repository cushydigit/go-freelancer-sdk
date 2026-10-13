package usecases

import (
	"context"
	"fmt"

	"github.com/cushydigit/go-freelancer-sdk/freelancer"
)

func AwardBidProject(ctx context.Context, c *freelancer.Client) {
	res, _, err := c.Resources.Bids.Award(ctx, 1)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Bid awarded: %s", res.Result)
}

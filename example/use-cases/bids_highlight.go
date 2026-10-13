package usecases

import (
	"context"
	"fmt"

	"github.com/cushydigit/go-freelancer-sdk/freelancer"
)

func HighlightBidOnProject(ctx context.Context, c *freelancer.Client) {
	res, _, err := c.Resources.Bids.Highlight(ctx, 1)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Bid highlighted: %v", res.Result)
}

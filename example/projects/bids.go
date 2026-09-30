package projects

import (
	"context"
	"fmt"

	"github.com/cushydigit/go-freelancer-sdk/freelancer"
	rr "github.com/cushydigit/go-freelancer-sdk/freelancer/reqres"
)

func GetBids(ctx context.Context, c *freelancer.Client) {
	res, _, err := c.Resources.Bids.List(
		ctx,
		&rr.ListBidsOptions{
			Projects: []int64{101, 102}, // list of project ids
			Limit:    rr.Int(20),        // narrow down results
		},
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("List of bids: %s", res.Result)
}

func AwardBid(ctx context.Context, c *freelancer.Client) {
	res, _, err := c.Resources.Bids.Award(ctx, 1)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Bid awarded: %s", res.Result)
}

func AcceptBid(ctx context.Context, c *freelancer.Client) {
	res, _, err := c.Resources.Bids.Accept(ctx, 1)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Bid accepted: %s", res.Result)
}

func RevokeBid(ctx context.Context, c *freelancer.Client) {
	res, _, err := c.Resources.Bids.Revoke(ctx, 1)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Bid revoked: %s", res.Result)
}

func RetractBid(ctx context.Context, c *freelancer.Client) {
	res, _, err := c.Resources.Bids.Revoke(ctx, 1)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Bid retracted: %s", res.Result)
}

func HighlightBid(ctx context.Context, c *freelancer.Client) {
	res, _, err := c.Resources.Bids.Highlight(ctx, 1)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Bid highlighted: %s", res.Result)
}

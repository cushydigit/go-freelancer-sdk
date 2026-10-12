package usecases

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

	fmt.Printf("List of bids: %v", res.Result)
}

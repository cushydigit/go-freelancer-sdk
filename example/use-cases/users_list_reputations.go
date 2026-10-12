package usecases

import (
	"context"
	"fmt"
	"log"

	"github.com/cushydigit/go-freelancer-sdk/freelancer"
	rr "github.com/cushydigit/go-freelancer-sdk/freelancer/reqres"
)

func ListReputations(ctx context.Context, c *freelancer.Client) {
	role := rr.RoleFreelancer
	res, _, err := c.Resources.Users.ListReputations(
		ctx,
		&rr.ListReputationsOptions{
			Users: []int64{1, 2, 3},
			Jobs:  []int64{1, 2},
			Role:  &role,
		},
	)

	if err != nil {
		log.Println(err)
	}

	fmt.Printf("Reputations: %v", res.Result)
}

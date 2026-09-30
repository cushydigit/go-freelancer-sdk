package users

import (
	"context"
	"fmt"

	"github.com/cushydigit/go-freelancer-sdk/freelancer"
	rr "github.com/cushydigit/go-freelancer-sdk/freelancer/reqres"
)

func FetchAndDisplayFreelancers(ctx context.Context, c *freelancer.Client) {
	searchOpts := rr.SearchFreelancerOptions{

		Limit:  rr.Int(20),
		Offset: rr.Int(0),

		// Include detailed data
		Reputation:         rr.Bool(true),
		ProfileDescription: rr.Bool(true),
	}

	res, _, err := c.Resources.Users.SearchFreelancer(ctx, &searchOpts)
	fmt.Println(res.Result.TotalCount)
	if err == nil && len(res.Result.Users) > 0 {
		for _, u := range res.Result.Users {
			fmt.Printf("-%d: %s\n", u.ID, u.DisplayName)
		}
	}

}

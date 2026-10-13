package usecases

import (
	"context"
	"fmt"
	"log"

	"github.com/cushydigit/go-freelancer-sdk/freelancer"
	rr "github.com/cushydigit/go-freelancer-sdk/freelancer/reqres"
)

func ListPortfolios(ctx context.Context, c *freelancer.Client) {
	res, _, err := c.Resources.Users.ListPortfolios(
		ctx,
		&rr.ListPortfoliosOptions{
			Users: []int64{1, 2, 3},
		},
	)

	if err != nil {
		log.Println(err)
	}

	fmt.Printf("Portfolios: %v", res.Result)
}

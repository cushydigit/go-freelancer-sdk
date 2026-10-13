package usecases

import (
	"context"
	"fmt"
	"log"

	"github.com/cushydigit/go-freelancer-sdk/freelancer"
)

func GetCurrentUser(ctx context.Context, c *freelancer.Client) {
	res, _, err := c.Resources.Self.Get(ctx, nil)
	if err != nil {
		log.Println(err)
	}

	fmt.Printf("Current user: %v", res.Result)
}

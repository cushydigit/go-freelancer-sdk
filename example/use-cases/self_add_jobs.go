package usecases

import (
	"context"
	"fmt"
	"log"

	"github.com/cushydigit/go-freelancer-sdk/freelancer"
	rr "github.com/cushydigit/go-freelancer-sdk/freelancer/reqres"
)

func AddUserJobs(ctx context.Context, c *freelancer.Client) {
	res, _, err := c.Resources.Self.AddJobs(
		ctx,
		rr.JobsBody{
			Jobs: []int64{
				20, 32,
			},
		},
	)

	if err != nil {
		log.Println(err)
	}

	fmt.Printf("User jobs added: %v", res.Result)
}

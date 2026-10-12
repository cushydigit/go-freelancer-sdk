package usecases

import (
	"context"
	"fmt"
	"log"

	"github.com/cushydigit/go-freelancer-sdk/freelancer"
	rr "github.com/cushydigit/go-freelancer-sdk/freelancer/reqres"
)

func PostReviewForEmployer(ctx context.Context, c *freelancer.Client) {
	res, _, err := c.Resources.Reviews.CreateForEmployer(
		ctx,
		rr.CreateReviewForEmployerBody{
			ReviewBody: rr.ReviewBody{
				ProjectID:  1,
				ToUserID:   101, // employer user id
				FromUserID: 102, // freelancer user id
				ReviewType: rr.ReviewTypeProject,
				Comment:    "Thanks",
			},
			ReputationData: rr.EmployerReputationData{
				Category: rr.EmployerCategoryRatings{
					Communication:   5,
					Professionalism: 5,
					Clarity:         4,
					Payment:         3,
					WorkForAgain:    4,
				},
			},
		},
	)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Review created: %v", res.Result)
}

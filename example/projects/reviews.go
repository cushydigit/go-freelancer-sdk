package projects

import (
	"context"
	"fmt"
	"log"

	"github.com/cushydigit/go-freelancer-sdk/freelancer"
	rr "github.com/cushydigit/go-freelancer-sdk/freelancer/reqres"
)

func PostReviewForEmployerExample(ctx context.Context, c *freelancer.Client) {
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

func PostReviewForFreelancerExample(ctx context.Context, c *freelancer.Client) {
	res, _, err := c.Resources.Reviews.CreateForFreelancer(
		ctx,
		rr.CreateReviewForFreelancerBody{
			ReviewBody: rr.ReviewBody{
				ProjectID:  1,
				ToUserID:   101, // freelancer user id
				FromUserID: 102, // employer user id
				ReviewType: rr.ReviewTypeProject,
				Comment:    "Thanks",
			},
			ReputationData: rr.FreelancerReputationData{
				OnBudget: 5,
				OnTime:   5,
				Category: rr.FreelancerCategoryRatings{
					Communication:   5,
					Professionalism: 5,
					Expertise:       5,
					Quality:         5,
					HireAgain:       5,
				},
			},
		},
	)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Review created: %v", res.Result)
}

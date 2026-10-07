package common

import (
	"context"
	"fmt"

	"github.com/cushydigit/go-freelancer-sdk/freelancer"
	rr "github.com/cushydigit/go-freelancer-sdk/freelancer/reqres"
)

func FetchAndDisplayCategories(ctx context.Context, c *freelancer.Client) {
	opt := rr.ListCategoriesOptions{
		Categories: []int64{
			*rr.Int64(1),   // "Websites, IT & Software"
			*rr.Int64(104), // "Artificial Intelligence"
		},
		JobDetails: rr.Bool(true), // return map of category IDs to jobs in those categories
	}
	res, _, err := c.Resources.Common.ListCategories(ctx, &opt)
	if err == nil && len(res.Result.Categories) > 0 {
		for _, cat := range res.Result.Categories {
			fmt.Printf("%d: %s\n", cat.ID, cat.Name)
		}
		fmt.Printf("Fetched %d categories\n", len(res.Result.Categories))
		// jobs of Artificial Intelligence
		if jobs, exists := res.Result.Jobs["104"]; exists {
			for _, job := range jobs {
				fmt.Printf("%d: Name %s\n", job.ID, job.Name)
			}
			fmt.Printf("Fetch %d for Artificial Intelligence\n", len(jobs))
		}
	}
}

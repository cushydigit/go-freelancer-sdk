package usecases

import (
	"context"
	"fmt"

	"github.com/cushydigit/go-freelancer-sdk/freelancer"
	rr "github.com/cushydigit/go-freelancer-sdk/freelancer/reqres"
)

func CreateFixedProject(ctx context.Context, c *freelancer.Client) {

	res, _, err := c.Resources.Projects.Create(
		ctx,
		rr.CreateProjectBody{
			Title:       "Go Developer Needed",
			Description: "Need an experienced Go developer...",
			Budget: rr.Budget{
				Minimum:    100.0,
				Maximum:    250.0,
				CurrencyID: 1,
			},
			Jobs: []int64{7},
			Type: rr.Enum(rr.ProjectBudgetFixed),
		},
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Project created with ID: %d\n", res.Result.ID)
}

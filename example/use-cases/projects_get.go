package usecases

import (
	"context"
	"fmt"

	"github.com/cushydigit/go-freelancer-sdk/freelancer"
	rr "github.com/cushydigit/go-freelancer-sdk/freelancer/reqres"
)

func GetSpecificProject(ctx context.Context, c *freelancer.Client) {
	res, _, err := c.Resources.Projects.Get(
		ctx,
		40681584, // project id
		&rr.GetProjectOptions{
			FullDescription: rr.Bool(true),
			UserDetails:     rr.Bool(true),
		},
	)
	if err == nil {
		p := res.Result
		budgetString := fmt.Sprintf(
			"[%s%1.f - %s%1.f]",
			p.Currency.Sign,
			p.Budget.Minimum,
			p.Currency.Sign,
			p.Budget.Minimum,
		)
		fmt.Printf("\n-%d: %s %s\n\n%s\n", p.ID, budgetString, p.Title, p.Description)
	}
}

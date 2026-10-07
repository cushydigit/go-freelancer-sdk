package projects

import (
	"context"
	"fmt"
	"time"

	"github.com/cushydigit/go-freelancer-sdk/freelancer"
	rr "github.com/cushydigit/go-freelancer-sdk/freelancer/reqres"
)

func SearchActiveProjectsExample(ctx context.Context, c *freelancer.Client) {

	from := time.Now().Add(-24 * time.Hour)
	res, _, err := c.Resources.Projects.SearchActive(
		ctx,
		&rr.SearchActiveProjectsOptions{
			Query:           rr.String("Go developer"),
			Limit:           rr.Int(20), // Results per page
			Offset:          rr.Int(0),
			FullDescription: rr.Bool(true), // include

			// Filtering options
			FromTime:    &from,         // Filter projects within 24 hours
			UserDetails: rr.Bool(true), // Include user info
			// Sort
			SortField:   rr.Enum(rr.SortFieldsTimeUpdated),
			ReverseSort: rr.Bool(false), // Newest first

		},
	)

	if err == nil && len(res.Result.Projects) > 0 {
		for _, p := range res.Result.Projects {
			budgetString := fmt.Sprintf(
				"[%s%1.f - %s%1.f]",
				p.Currency.Sign,
				p.Budget.Minimum,
				p.Currency.Sign,
				p.Budget.Minimum,
			)
			fmt.Printf("\n-%d: %s %s\n", p.ID, budgetString, p.Title)
			// access unix time easily
			time1, time2, time3 := p.SubmitDateAt(), p.UpdatedAt(), p.SubmittedAt()
			if time1 != nil && time2 != nil && time3 != nil {
				fmt.Printf(
					"SubmitDate: %s\tTimeUpdated: %s\tTimeSubmitted: %s\n",
					time1.Format("Jan 2, 2006 at 15:04"),
					time2.Format("Jan 2, 2006 at 15:04"),
					time3.Format("Jan 2, 2006 at 15:04"),
				)
			}
		}
		fmt.Printf("Showing %d of %d total projects\n", len(res.Result.Projects), res.Result.TotalCount)
	}
}

func ListSpecificProjectsExample(ctx context.Context, c *freelancer.Client) {

	res, _, err := c.Resources.Projects.List(
		ctx,
		&rr.ListProjectsOptions{
			Projects: []int64{
				*rr.Int64(40681584),
				*rr.Int64(40681432),
			},
			UserDetails: rr.Bool(true), // Include user info

		},
	)

	if err == nil && len(res.Result.Projects) > 0 {
		for _, p := range res.Result.Projects {
			budgetString := fmt.Sprintf(
				"[%s%1.f - %s%1.f]",
				p.Currency.Sign,
				p.Budget.Minimum,
				p.Currency.Sign,
				p.Budget.Minimum,
			)
			fmt.Printf("\n-%d: %s %s\n", p.ID, budgetString, p.Title)
		}
		fmt.Printf("Showing %d of %d total projects\n", len(res.Result.Projects), res.Result.TotalCount)
	}
}

func GetSpecificProjectExample(ctx context.Context, c *freelancer.Client) {
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

func CreateProjectExample(ctx context.Context, c *freelancer.Client) {

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

func CreateHourlyProjectExample(ctx context.Context, c *freelancer.Client) {

	res, _, err := c.Resources.Projects.Create(
		ctx,
		rr.CreateProjectBody{
			Title:       "Go Developer Needed",
			Description: "Need an experienced Go developer...",
			Budget: rr.Budget{
				Minimum:    10,
				CurrencyID: 1,
			},
			Jobs: []int64{6, 9},
			HourlyProjectInfo: &rr.HourlyProjectInfo{
				Commitment: rr.Commitment{
					Hours:    40,
					Interval: rr.IntervalWeek,
				},
			},
		},
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Project created with ID: %d\n", res.Result.ID)
}

func CreateHireMeProjectExample(ctx context.Context, c *freelancer.Client) {

	res, _, err := c.Resources.Projects.Create(
		ctx,
		rr.CreateProjectBody{
			Title:       "Go Developer Needed",
			Description: "Need an experienced Go developer...",
			Budget: rr.Budget{
				Minimum:    10,
				CurrencyID: 1,
			},
			Jobs:  []int64{6, 9},
			HirMe: rr.Bool(true),
			HiremeInitialBid: &rr.HiremeInitialBid{
				BidderID:    100, // Freelancer we want to hire
				Amount:      200,
				Period:      5, // number of days
				Description: "Hello, Are you interested in my project",
			},
		},
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Project created with ID: %d\n", res.Result.ID)
}

func CreateLocalProjectExample(ctx context.Context, c *freelancer.Client) {

	res, _, err := c.Resources.Projects.Create(
		ctx,
		rr.CreateProjectBody{
			Title:       "Need assistant on local project",
			Description: "Description",
			Budget: rr.Budget{
				Minimum:    100,
				CurrencyID: 1,
			},
			Jobs: []int64{649}, // must be a local job
			Location: &rr.Location{
				Country: &rr.Country{
					Name: "Australia",
				},
				City:      "Sydney",
				Latitude:  rr.Float64(-33.875461),
				Longitude: rr.Float64(151.201678),
			},
		},
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Local project created with ID: %d\n", res.Result.ID)
}

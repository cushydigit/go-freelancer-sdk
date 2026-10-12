package usecases

import (
	"context"
	"fmt"

	"github.com/cushydigit/go-freelancer-sdk/freelancer"

	rr "github.com/cushydigit/go-freelancer-sdk/freelancer/reqres"
)

func FetchAndDisplayCountries(ctx context.Context, c *freelancer.Client) {
	res, _, err := c.Resources.Common.ListCountries(ctx, &rr.ListCountriesOptions{
		ExtraDetails: rr.Bool(true), // include extra details
	})
	// nil slices has length of zero
	if err == nil && len(res.Result.Countries) > 0 {
		for _, c := range res.Result.Countries {
			fmt.Printf("Name: %s, Code: %s PhoneCode: %.0f\n", c.Name, c.Code, c.PhoneCode)
		}
		fmt.Printf("Fetched %d countries\n", len(res.Result.Countries))
	}
}

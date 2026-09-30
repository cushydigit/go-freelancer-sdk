package common

import (
	"context"
	"fmt"

	"github.com/cushydigit/go-freelancer-sdk/freelancer"
	rr "github.com/cushydigit/go-freelancer-sdk/freelancer/reqres"
)

func FetchAndDisplayCountries(ctx context.Context, c *freelancer.Client) {
	opts := rr.ListCountriesOptions{
		ExtraDetails: rr.Bool(true), // Include extra details
	}

	res, _, err := c.Resources.Common.ListCountries(ctx, &opts)
	// nil slices has length of zero
	if err == nil && len(res.Result.Countries) > 0 {
		for _, c := range res.Result.Countries {
			fmt.Printf("Name: %s, Code: %s PhoneCode: %.0f\n", c.Name, c.Code, c.PhoneCode)
		}
		fmt.Printf("Fetched %d countries\n", len(res.Result.Countries))
	}
}

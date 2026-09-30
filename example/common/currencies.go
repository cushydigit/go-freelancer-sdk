package common

import (
	"context"
	"fmt"

	"github.com/cushydigit/go-freelancer-sdk/freelancer"
)

func FetchAndDisplayCurrencies(ctx context.Context, c *freelancer.Client) {
	res, _, err := c.Resources.Common.ListCurrencies(ctx, nil)
	// nil slices has length of zero
	if len(res.Result.Currencies) > 0 && err == nil {
		for _, cur := range res.Result.Currencies {
			// Each currency record supports exchange rates for calculations
			fmt.Printf("(%s)%d: %.3f USD rate, Country=%s\n", cur.Code, cur.ID, cur.ExchangeRate, cur.Country)
		}
		fmt.Printf("Fetched %d countries\n", len(res.Result.Currencies))
	}
}

package common

import (
	"context"
	"fmt"

	"github.com/cushydigit/go-freelancer-sdk/freelancer"
	rr "github.com/cushydigit/go-freelancer-sdk/freelancer/reqres"
)

func FetchAndDisplayTimezones(ctx context.Context, c *freelancer.Client) {
	opts := rr.ListTimezonesOptions{
		TimezoneNames: []string{
			"America/New_York",
			"Europe/London",
			"Asia/Tokyo",
		},
	}

	res, _, err := c.Resources.Common.ListTimezones(ctx, &opts)
	// nil slices has length of zero
	if err == nil && len(res.Result.Timezones) > 0 {
		for _, tz := range res.Result.Timezones {
			fmt.Printf("%d: country=%s (UTC+%.1f)\n", tz.ID, tz.Country, tz.Offset)
		}
		fmt.Printf("Fetched %d timezone(s)\n", len(res.Result.Timezones))
	}
}

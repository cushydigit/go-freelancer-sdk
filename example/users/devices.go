package users

import (
	"context"
	"fmt"

	"github.com/cushydigit/go-freelancer-sdk/freelancer"
)

func FetchAndDisplayDevices(ctx context.Context, c *freelancer.Client) {
	res, _, err := c.Resources.Self.ListDevices(ctx)
	if err == nil && len(res.Result.Devices) > 0 {
		for _, d := range res.Result.Devices {
			fmt.Printf("- %s in %s %s\n", d.City, d.Country, d.Platform)
		}
	}
}

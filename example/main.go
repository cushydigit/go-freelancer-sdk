package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/cushydigit/go-freelancer-sdk/example/client"
	"github.com/cushydigit/go-freelancer-sdk/example/common"
)

func main() {
	ctx := context.Background()
	l := slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}),
	)
	useSandBox := false
	c := client.Init(l, useSandBox)
	// common.FetchAndDisplayTimezones(ctx, c)
	common.FetchAndDisplayCountries(ctx, c)
	// projects.FetchAndDisplayCurrencies(ctx, c)
	// projects.FetchAndDisplayBudget(ctx, c)
	// projects.FetchAndDisplayCategories(ctx, c)
	// projects.FetchAndDisplayProjects(ctx, c)
	// projects.CreateProject(ctx, c)
	// users.FetchAndDisplayFreelancers(ctx, c)
	// users.FetchAndDisplayDevices(ctx, c)
}

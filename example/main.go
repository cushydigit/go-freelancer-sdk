package examples

import (
	"context"
	"log/slog"
	"os"

	"github.com/cushydigit/go-freelancer-sdk/example/client"
	usecases "github.com/cushydigit/go-freelancer-sdk/example/use-cases"
)

func main() {
	ctx := context.Background()
	l := slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelError,
		}),
	)
	useSandBox := false
	c := client.Init(l, useSandBox)
	usecases.ListActiveProjects(ctx, c)
	usecases.FetchAndDisplayBudget(ctx, c)
	usecases.FetchAndDisplayTimezones(ctx, c)
}

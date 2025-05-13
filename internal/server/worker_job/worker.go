package worker_job

import (
	"context"
	"go.temporal.io/sdk/client"
	"go.uber.org/fx"
	"pointfive/internal/config"
)

func NewWorkerClient(lc fx.Lifecycle, cfg *config.Config) (client.Client, error) {
	c, err := client.Dial(client.Options{})

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			c.Close()
			return nil
		},
	})
	return c, err
}

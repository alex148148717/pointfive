package server

import (
	"go.uber.org/fx"
	"net/http"
	"pointfive/internal/config"
	"pointfive/internal/import_job"
	"pointfive/internal/import_worker"
	"pointfive/internal/server/db"
	"pointfive/internal/server/worker_job"
)

func NewApp() *fx.App {
	return fx.New(
		fx.Provide(
			config.LoadConfig,
			db.NewDB,
			worker_job.NewWorkerClient,
			NewRouter,
			NewHTTPServer,
		),
		fx.Options(
			import_job.Module,
			import_worker.ImportWorkerModule,
		),
		fx.Invoke(
			func(*http.Server) {},
		),
	)
}

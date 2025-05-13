package import_job

import (
	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
	"pointfive/internal/import_job/domain"
	"pointfive/internal/import_job/infrastructure"
	"pointfive/internal/import_job/interfaces"
)

var Module = fx.Options(
	fx.Provide(
		domain.NewService,
		infrastructure.NewRepository,
		infrastructure.NewWorkerRepository,
		interfaces.NewHandler,
		infrastructure.NewLocalFileRepository,
	),
	fx.Invoke(func(router *chi.Mux, handler *interfaces.ImportJobHandler) {
		router.Post("/import_job/v1/job", handler.Handler)

	}),
)

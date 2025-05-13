package server

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
	"net/http"
	"pointfive/internal/config"
	"time"
)

func NewRouter() *chi.Mux {
	return chi.NewRouter()
}
func NewHTTPServer(lc fx.Lifecycle, config *config.Config, mux *chi.Mux) *http.Server {

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", config.ServerPort),
		Handler:      mux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					panic(err)
				}
			}()
			return nil

		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})

	return srv
}

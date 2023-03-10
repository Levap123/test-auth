package app

import (
	"context"
	"net/http"
	"time"

	"github.com/Levap123/salt-gen/internal/service"
	"github.com/Levap123/salt-gen/internal/transport"
)

type App struct {
	srv *http.Server
}

func NewApp() (*App, error) {
	service := service.NewService()
	transport := transport.NewTransport(service)
	routes := transport.InitRoutes()
	srv := &http.Server{
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
		Handler:      routes,
		Addr:         ":5000",
	}
	return &App{
		srv: srv,
	}, nil
}

func (a *App) Run() error {
	return a.srv.ListenAndServe()
}

func (a *App) Shutdown(ctx context.Context) error {
	return a.srv.Shutdown(ctx)
}

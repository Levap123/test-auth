package app

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/Levap123/auth-service/internal/repository"
	"github.com/Levap123/auth-service/internal/repository/mongor"
	"github.com/Levap123/auth-service/internal/service"
	"github.com/Levap123/auth-service/internal/transport"
	"github.com/Levap123/auth-service/internal/validator"
)

type App struct {
	srv *http.Server
}

func NewApp() (*App, error) {
	db, err := mongor.InitDb()
	if err != nil {
		return nil, fmt.Errorf("new app - %w", err)
	}
	repos := repository.NewRepoMongo(db)
	service := service.NewService(repos)
	validator := validator.NewValidator()
	transport := transport.NewTransport(service, validator)
	routes := transport.InitRoutes()
	srv := &http.Server{
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
		Handler:      routes,
		Addr:         ":8080",
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

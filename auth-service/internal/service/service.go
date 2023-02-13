package service

import (
	"context"

	"github.com/Levap123/auth-service/internal/models"
	"github.com/Levap123/auth-service/internal/repository"
)

type Service struct {
	UserService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		UserService: NewUser(repo.UserRepo),
	}
}

type UserService interface {
	Create(ctx context.Context, email, password string) error
	Get(ctx context.Context, email string) (*models.User, error)
}

package repository

import (
	"context"

	"github.com/Levap123/auth-service/internal/models"
	"github.com/Levap123/auth-service/internal/repository/mongor"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	UserRepo
}

type UserRepo interface {
	Create(ctx context.Context, user *models.User) error
	Get(ctx context.Context, email string) (*models.User, error)
}

func NewRepoMongo(client *mongo.Client) *Repository {
	return &Repository{
		UserRepo: mongor.NewUser(client),
	}
}

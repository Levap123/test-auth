package mongor

import (
	"context"
	"fmt"

	"github.com/Levap123/auth-service/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	coll *mongo.Collection
}

func NewUser(client *mongo.Client) *User {
	return &User{
		coll: client.Database("auth_service").Collection("users"),
	}
}

func (u *User) Create(ctx context.Context, user *models.User) error {
	_, err := u.coll.InsertOne(ctx, user)
	if err != nil {
		return fmt.Errorf("mongo repo - create user - %w", err)
	}
	return nil
}

func (u *User) Get(ctx context.Context, email string) (*models.User, error) {
	filter := bson.M{"email": email}
	result := u.coll.FindOne(ctx, filter)
	var user models.User
	if err := result.Decode(&user); err != nil {
		return nil, fmt.Errorf("mongo repo - get user - %w", err)
	}
	return &user, nil
}

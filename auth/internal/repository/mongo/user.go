package mongo

import "go.mongodb.org/mongo-driver/mongo"

type User struct {
	coll *mongo.Collection
}

func NewUser(client *mongo.Client) *User {
	return &User{
		coll: client.Database("auth_service").Collection("users"),
	}
}

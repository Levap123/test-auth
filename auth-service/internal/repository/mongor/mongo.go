package mongor

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func InitDb() (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	db, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://root:root@mongodb:27017/"))
	if err != nil {
		return nil, err
	}
	if err := db.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}

	return db, nil
}

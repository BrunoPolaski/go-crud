package mongodb

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoConnection(ctx context.Context) (*mongo.Database, error) {
	uri := os.Getenv("MONGO_URL")
	mongoDatabase := os.Getenv("MONGO_USERS_DATABASE")
	username := os.Getenv("MONGO_USERNAME")
	password := os.Getenv("MONGO_PASSWORD")

	// Check if environment variables are set
	if uri == "" || mongoDatabase == "" || username == "" || password == "" {
		return nil, fmt.Errorf("required environment variables are not set")
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri).SetAuth(options.Credential{
		Username: username,
		Password: password,
	}))
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return client.Database(mongoDatabase), nil
}

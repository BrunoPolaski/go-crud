package mongodb

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MONGO_URL            = "MONGO_URL"
	MONGO_USERS_DATABASE = "MONGO_USERS_DATABASE"
)

func NewMongoConnection(
	ctx context.Context,
) (*mongo.Database, error) {
	uri := os.Getenv(MONGO_URL)
	mongoDatabase := os.Getenv(MONGO_USERS_DATABASE)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return client.Database(mongoDatabase), nil
}

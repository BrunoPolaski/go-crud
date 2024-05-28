package mongodb

import (
	"context"

	"github.com/BrunoPolaski/go-crud/src/configuration/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitDatabase() {
	ctx := context.Background()

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	if client, err := mongo.Connect(ctx, clientOptions); err != nil {
		panic(err)
	} else {
		if err = client.Ping(ctx, nil); err != nil {
			panic(err)
		}
	}

	logger.Info("Connected to MongoDB")
}

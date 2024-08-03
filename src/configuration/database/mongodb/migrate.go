package mongodb

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

func Migrate() {
	db, err := NewMongoConnection(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	if err := db.CreateCollection(context.Background(), "apiKey"); err != nil {
		log.Fatal(err)
	}

	apiKeyCollection := db.Collection("apiKey")

	uuid := uuid.New().String()

	apiKey := bson.D{
		{Key: "name", Value: "usersApi"},
		{Key: "key", Value: uuid},
		{Key: "createdAt", Value: time.Now()},
	}

	if key := apiKeyCollection.FindOne(context.Background(), bson.D{{Key: "name", Value: "usersApi"}}); key != nil {
		log.Println("API key already exists")
		return
	}

	insertResult, err := apiKeyCollection.InsertOne(context.TODO(), apiKey)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Migration successful", insertResult.InsertedID)
}

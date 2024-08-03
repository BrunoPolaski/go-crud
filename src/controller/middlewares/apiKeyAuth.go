package middlewares

import (
	"context"

	"github.com/BrunoPolaski/go-crud/src/configuration/database/mongodb"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func ApiKeyAuth(c *gin.Context) {
	db, _ := mongodb.NewMongoConnection(context.Background())
	collection := db.Collection("apiKey")

	apiKey := c.Request.Header.Get("api-key")
	if apiKey == "" {
		c.JSON(401, "Unauthorized")
		c.Abort()
		return
	}

	filter := bson.D{{Key: "key", Value: apiKey}}

	if key := collection.FindOne(context.Background(), filter); key == nil {
		c.JSON(401, "Unauthorized")
		c.Abort()
		return
	}
}

package database

import (
	"context"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	Client     *mongo.Client
	Collection *mongo.Collection
	Ctx        context.Context
}

func OpenMongoDBConnection() (*MongoDB, error) {
	godotenv.Load(".env")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(os.Getenv("DB_MONGO_URI"))
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	collection := client.Database("Auth").Collection("users")

	return &MongoDB{Client: client, Collection: collection, Ctx: ctx}, nil
}

package database

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	Client     *mongo.Client
	Collection *mongo.Collection
}

var cred options.Credential

func OpenMongoDBConnection() (*MongoDB, error) {
	godotenv.Load(".env")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cred.Username = os.Getenv("MONGO_USER")
	cred.Password = os.Getenv("MONGO_PASSWORD")

	clientOptions := options.Client().ApplyURI(os.Getenv("DB_MONGO_URI")).SetAuth(cred)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}
	log.Println("Successfully connected and pinged.")

	collection := client.Database("auth").Collection("users")

	return &MongoDB{Client: client, Collection: collection}, nil
}

package connection

import (
	"context"
	"fmt"
	"log"
	"time"

	configs "github.com/siddhantprateek/opendesk/configs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoDBdatabase() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(configs.GetEnv("MONGO_URI")))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to MongoDB Atlas")
	return client
}

// @desp: client instance
var DB *mongo.Client = MongoDBdatabase()

// @desp: database collection
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("opendeskdb").Collection(collectionName)
	return collection
}

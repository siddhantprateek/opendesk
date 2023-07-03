package connection

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/siddhantprateek/opendesk/configs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoDbConnection() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(configs.GetEnv("MONGO_URI")))
	if err != nil {
		log.Fatal("Unable to create mongo Client.", err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connect to Mongo Atlas.")
	return client
}

var DB *mongo.Client = MongoDbConnection()

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("opendeskdb").Collection(collectionName)
	return collection
}

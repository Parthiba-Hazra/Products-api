package configs

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client {

	//checking env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error while loding env file")
	}
	mongo_URI := os.Getenv("MONGO_URI")

	client, err := mongo.NewClient(options.Client().ApplyURI(mongo_URI))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	//ping the MongoDB database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")
	return client
}

// client
var DB *mongo.Client = ConnectDB()

// Database Collection
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("productsAPI").Collection(collectionName)
	return collection
}

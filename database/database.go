package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client
var DATABASE = "settle_up"

func InitDB() {
	if err := godotenv.Load(); err != nil {
		log.Println("NO .env found")
	}

	uri := os.Getenv("MONGODB_URI")

	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	// defer func() {
	// 	if err := client.Disconnect(context.TODO()); err != nil {
	// 		panic(err)
	// 	}
	// }()

	mongoClient = client
	fmt.Println("MongoDB connected")
}

func UserCollection() *mongo.Collection {
	return getCollection("users")
}

func getCollection(collectionName string) *mongo.Collection {
	return mongoClient.Database(DATABASE).Collection(collectionName)
}

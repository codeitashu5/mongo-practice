package mongoClient

import (
	"context"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

// MongoClient making MongoClient singleTon to be used to create database using it
var MongoClient *mongo.Client

func SetUpMongo() {
	var err error
	if err = godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. " +
			"See https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	MongoClient, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
}

func ShutdownMongo() error{
	 err := MongoClient.Disconnect(context.TODO())
	 return err
}

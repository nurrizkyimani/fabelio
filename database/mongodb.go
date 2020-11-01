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


func Mongoinit() *mongo.Client {

	e := godotenv.Load()
	if e != nil {
			if e != nil {
			panic("no env file")
		}
	}

	if e != nil {
		panic("no env file")
	}

	MongoURI := os.Getenv("Mongo")
	client, err := mongo.NewClient(options.Client().ApplyURI(MongoURI)) 
	if err != nil {
      log.Fatal(err)
  }

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
      log.Fatal(err)
		}
		
	defer client.Disconnect(ctx)

	return client

}
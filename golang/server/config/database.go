package config

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	database string = "crud"
	collect  string = "col"
	client   *mongo.Client

	Collection *mongo.Collection
	Mongoctx   context.Context
)

func Init() {
	fmt.Println("+ Initializing database connection")
	Mongoctx = context.Background()
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(Mongoctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check whether the connection was succesful by pinging MongoDB server
	if err := client.Ping(Mongoctx, nil); err != nil {
		log.Fatalf("- Database connection error: %s", err)
	} else {
		fmt.Println("+ Database connection established")
	}

	Collection = client.Database(database).Collection(collect)
}

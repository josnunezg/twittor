package bd

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoCN is the connection object to BD
var MongoCN = ToConnectBD()
var clientOptions = options.Client().ApplyURI(fmt.Sprintf("mongodb+srv://%s:%s@twittor-dsdpy.gcp.mongodb.net/test?retryWrites=true&w=majority", os.Getenv("MONGO_DB_USER"), os.Getenv("MONGO_DB_PASSWORD")))

// ToConnectBD is a function that handle connection to bd
func ToConnectBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("Successfully connection to the BD")
	return client
}

// CheckingConnection is a function to check connection with the bd
func CheckingConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)

	if err != nil {
		return 0
	}

	return 1
}

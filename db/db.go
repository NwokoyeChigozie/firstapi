package db

import (
	"fmt"
)

func ConnectMONGODB() {
	fmt.Println("Starting connection")
	// Set client options
	// clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// // Connect to MongoDB
	// client, err := mongo.Connect(context.TODO(), clientOptions)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // Check the connection
	// err = client.Ping(context.TODO(), nil)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	fmt.Println("Connected to MongoDB!")
}

// package database

// import (
// 	"context"
// 	"time"

// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// 	"go.mongodb.org/mongo-driver/mongo/readpref"
// )

// // ConnectMongo establishes a connection to MongoDB using the provided URI.
// func ConnectMongo(uri, dbName string) (*mongo.Client, error) {
// 	// Set a timeout for the connection attempt.
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Ping the primary node to verify that the connection is alive.
// 	err = client.Ping(ctx, readpref.Primary())
// 	if err != nil {
// 		return nil, err
// 	}

// 	return client, nil
// }
package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongo(mongoURI string, dbName string) (*mongo.Client, error) {
	if mongoURI == "" {
		return nil, fmt.Errorf("mongoURI is empty")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	fmt.Printf(">>> USING MONGO URI: %#v\n", mongoURI)


	client, err := mongo.Connect(
		ctx,
		options.Client().ApplyURI(mongoURI),
	)
	if err != nil {
		return nil, err
	}

	// Ping to force connection
	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return client, nil
}

package sql

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestConnection() {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(DefaultMongoURI))
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(context.Background())

	// ping
	if err := client.Ping(context.Background(), nil); err != nil {
		panic(err)
	}
}

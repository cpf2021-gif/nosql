package sql

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getClient() *mongo.Client {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(DefaultMongoURI))
	if err != nil {
		panic(err)
	}

	return client
}

func InsertMany(dbName, collectionName string, data []interface{}) error {
	client := getClient()
	defer client.Disconnect(context.Background())

	// 插入数据
	coll := client.Database(dbName).Collection(collectionName)
	_, err := coll.InsertMany(context.Background(), data)
	return err
}

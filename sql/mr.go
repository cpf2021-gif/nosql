package sql

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Mr() {
	clientOptions := options.Client().ApplyURI(DefaultMongoURI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())

	db := client.Database("test")

	mapFunction := `
    function() {
        if (this.raw_keywords == null) {
            emit("其他", 1);
            return;
        }
        for (var i = 0; i < this.raw_keywords.length; i++) {
            var keyword = this.raw_keywords[i];
            for (var key in provinceMapping) {
                if (provinceMapping[key].indexOf(keyword) !== -1) {
                    emit(key, 1);
                    return;
                }
            }
        }
        emit("其他", 1);
    }`

	reduceFunction := `
    function(key, values) {
        return Array.sum(values);
    }`

	scope := bson.M{
		"provinceMapping": Province,
	}

	res := db.RunCommand(context.Background(), bson.D{
		{Key: "mapReduce", Value: "news"},
		{Key: "map", Value: mapFunction},
		{Key: "reduce", Value: reduceFunction},
		{Key: "out", Value: "news_out"},
		{Key: "scope", Value: scope},
	})

	if res.Err() != nil {
		log.Fatal(res.Err())
	}
}

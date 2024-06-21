package sql

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// 按省份统计新闻数量
func MrByProvince() {
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
		{Key: "out", Value: "news_by_province"},
		{Key: "scope", Value: scope},
	})

	if res.Err() != nil {
		log.Fatal(res.Err())
	}
}

// 按时间统计疫情新闻数量
func MrByCtime() {
	clientOptions := options.Client().ApplyURI(DefaultMongoURI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())

	db := client.Database("test")

	mapFunction := `
	function() {
		let date = new Date(+this.ctime * 1000);
		let key = date.getFullYear() + '-' + (date.getMonth() + 1);
		emit(key, 1);
	}
	`

	reduceFunction := `
	function(key, values) {
		return Array.sum(values);
	}
	`

	res := db.RunCommand(context.Background(), bson.D{
		{Key: "mapReduce", Value: "news"},
		{Key: "map", Value: mapFunction},
		{Key: "reduce", Value: reduceFunction},
		{Key: "query", Value: bson.M{"raw_keywords": bson.M{"$in": []string{"疫情", "新冠", "新冠肺炎", "新型冠状病毒", "防疫"}}}},
		{Key: "out", Value: "news_by_time"},
	})

	if res.Err() != nil {
		log.Fatal(res.Err())
	}
}

// 按类别统计新闻数量
func MrByCate() {
	clientOptions := options.Client().ApplyURI(DefaultMongoURI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())

	db := client.Database("test")

	mapFuction := `
	function() {
		emit(this.cate, 1);
	}
	`

	reduceFunction := `
	function(key, values) {
		return Array.sum(values);
	}
	`

	res := db.RunCommand(context.Background(), bson.D{
		{Key: "mapReduce", Value: "news"},
		{Key: "map", Value: mapFuction},
		{Key: "reduce", Value: reduceFunction},
		{Key: "out", Value: "news_by_cate"},
	})

	if res.Err() != nil {
		log.Fatal(res.Err())
	}
}

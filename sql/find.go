package sql

import (
	"context"
	"strconv"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
)

type ProvinceDocument struct {
	Id    string `bson:"_id"`
	Value int    `bson:"value"`
}

// 查询省份
func FindProvince() ([]ProvinceDocument, error) {
	var result []ProvinceDocument
	client := getClient()
	defer client.Disconnect(context.Background())

	collection := client.Database("test").Collection("news_by_province")

	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var document ProvinceDocument
		err := cursor.Decode(&document)
		if err != nil {
			return nil, err
		}
		result = append(result, document)
	}
	return result, nil
}

// 查询疫情时间数据
type TimeSeries struct {
	Id    string `bson:"_id"`
	Value int    `bson:"value"`
}

func FindTimeSeries() ([]TimeSeries, error) {
	var result []TimeSeries
	client := getClient()
	defer client.Disconnect(context.Background())

	collection := client.Database("test").Collection("news_by_time")

	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var document TimeSeries
		err := cursor.Decode(&document)
		if err != nil {
			return nil, err
		}
		result = append(result, document)
	}
	return result, nil
}

// 查询不同新闻类型的数量
type CateDocument struct {
	Id    string `bson:"_id"`
	Value int    `bson:"value"`
}

func FindCate() ([]CateDocument, error) {
	var result []CateDocument
	client := getClient()
	defer client.Disconnect(context.Background())

	collection := client.Database("test").Collection("news_by_cate")

	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var document CateDocument
		err := cursor.Decode(&document)
		if err != nil {
			return nil, err
		}
		result = append(result, document)
	}
	return result, nil
}

func CmpTime(a, b string) int {

	// time: XXXX-XX

	/*
		if a > b, return 1
		if a == b, return 0
		if a < b, return -1
	*/

	// split a and b
	astrs := strings.Split(a, "-")
	aYear, aMonth := astrs[0], astrs[1]
	bstrs := strings.Split(b, "-")
	bYear, bMonth := bstrs[0], bstrs[1]
	// 比较年份
	if aYear > bYear {
		return 1
	}
	if aYear < bYear {
		return -1
	}

	// 比较月份
	aMonthInt, _ := strconv.Atoi(aMonth)
	bMonthInt, _ := strconv.Atoi(bMonth)
	if aMonthInt > bMonthInt {
		return 1
	}
	if aMonthInt < bMonthInt {
		return -1
	}

	return 0
}

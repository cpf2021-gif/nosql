package main

import (
	"log"
	"nosql/extract"
	"nosql/sql"
	"time"
)

// go run cmd/load/load.go
func main() {
	// 提取数据
	extractStart := time.Now()
	data, err := extract.GetDocuments("data.txt")
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("extract data cost: %v\n", time.Since(extractStart))

	// 导入数据到 MongoDB
	loadStart := time.Now()
	err = sql.InsertMany("test", "news", data)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("load data cost: %v\n", time.Since(loadStart))
}

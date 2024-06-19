package main

import (
	// "nosql/mr"
	// "time"
	// "log"
	// "nosql/extract"
	"nosql/sql"
)


// go run main.go
func main() {
	// // 测试连接
	// sql.TestConnection()

	// // 提取数据
	// data, err := extract.GetDocuments("data.txt")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// // 导入数据到 MongoDB
	// err = sql.InsertMany("test", "news", data)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// 测试mr
	// _ = mr.MakeCoordinator(nil, 2)
	// for i := 0; i < 10; i++ {
	// 	go mr.Worker()
	// }

	// time.Sleep(3 * time.Second)
	sql.Mr()
}

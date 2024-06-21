package main

import (
	"log"
	"nosql/sql"
	"time"
)

// go run cmd/mapreduce/mapreduce.go
func main() {
	// 按省份统计新闻数量
	mrbyProvinceStart := time.Now()
	sql.MrByProvince()
	log.Printf("mrbyProvince cost: %v\n", time.Since(mrbyProvinceStart))

	// 按时间统计疫情新闻数量
	mrbyTimeStart := time.Now()
	sql.MrByCtime()
	log.Printf("mrbyTime cost: %v\n", time.Since(mrbyTimeStart))

	// 按类别统计疫情新闻数量
	mrbyCateStart := time.Now()
	sql.MrByCate()
	log.Printf("mrbyCate cost: %v\n", time.Since(mrbyCateStart))
}

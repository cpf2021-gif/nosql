package main

import (
	"log"
	"nosql/sql"
	"slices"
	"time"
)

// go run cmd/mapreduce/mapreduce.go
func main() {
	// 按省份统计新闻数量
	mrbyProvinceStart := time.Now()
	sql.MrByProvince()
	log.Printf("mrbyProvince cost: %v\n", time.Since(mrbyProvinceStart))
	// 获取省份列表
	provinceData, err := sql.FindProvince()
	if err != nil {
		log.Println(err)
	}
	log.Printf("province: %v\n", provinceData)

	// 按时间统计疫情新闻数量
	mrbyTimeStart := time.Now()
	sql.MrByCtime()
	log.Printf("mrbyTime cost: %v\n", time.Since(mrbyTimeStart))
	// 获取时间列表
	timeSeriesData, err := sql.FindTimeSeries()
	if err != nil {
		log.Println(err)
	}
	slices.SortFunc(timeSeriesData, func(a, b sql.TimeSeries) int {
		return sql.CmpTime(a.Id, b.Id)
	})
	log.Printf("time: %v\n", timeSeriesData)

	// 按类别统计新闻数量
	mrbyCateStart := time.Now()
	sql.MrByCate()
	log.Printf("mrbyCate cost: %v\n", time.Since(mrbyCateStart))
	// 获取类别列表
	categoryData, err := sql.FindCate()
	if err != nil {
		log.Println(err)
	}
	log.Printf("cate: %v\n", categoryData) 
}

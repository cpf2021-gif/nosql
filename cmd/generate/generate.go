package main

import "nosql/plot/example"

// go run cmd/generate/generate.go
func main() {
	// 生成数据图表
	examples := []example.Exampler{
		example.LineExamples{}, // 折线图
		example.BarExamples{},  // 柱状图
		example.PieExamples{},  // 饼图
		example.MapExamples{},  // 地图
	}

	for _, e := range examples {
		e.Examples()
	}
}

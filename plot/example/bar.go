package example

import (
	"cmp"
	"fmt"
	"io"
	"log"
	"nosql/sql"
	"os"
	"slices"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
)

// 生成新闻按省份统计图表 - 柱状图
func barTooltip() *charts.Bar {
	// 获取省份数据
	res, err := sql.FindProvince()
	if err != nil {
		log.Fatalln(err)
	}

	// 降序排序
	slices.SortFunc(res, func(i, j sql.ProvinceDocument) int {
		return cmp.Compare(j.Value, i.Value)
	})

	var province []string
	var count []opts.BarData

	length := len(res)
	for i, item := range res {
		if item.Id == "其他" {
			continue
		}
		province = append(province, item.Id)

		r := 255 * (length - i) / length
		color := fmt.Sprintf("rgb(%d, 50, 50)", r)

		count = append(count, opts.BarData{
			Value: item.Value,
			ItemStyle: &opts.ItemStyle{
				Color: color,
			},
			Label: &opts.Label{
				Show:     true,
				Position: "top",
			},
		})
	}

	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{
			Width:  "1800px",
			Height: "500px",
		}),
		charts.WithXAxisOpts(opts.XAxis{
			AxisLabel: &opts.AxisLabel{
				Show:         true,
				Rotate:       45,
				ShowMinLabel: true,
				ShowMaxLabel: true,
			},
		}),
		charts.WithTitleOpts(opts.Title{Title: "新闻按省份统计"}),
		charts.WithTooltipOpts(opts.Tooltip{
			Show:    true,
			Trigger: "item",
		}),
		charts.WithLegendOpts(opts.Legend{
			Show: true,
		}),
	)

	bar.SetXAxis(province).
		AddSeries("省份", count)
	return bar
}

type BarExamples struct{}

func (BarExamples) Examples() {
	page := components.NewPage()
	page.AddCharts(
		barTooltip(),
	)
	f, err := os.Create("static/bar.html")
	if err != nil {
		panic(err)
	}
	page.Render(io.MultiWriter(f))
}

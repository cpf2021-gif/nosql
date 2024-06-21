package example

import (
	"io"
	"log"
	"nosql/sql"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func pieRadius() *charts.Pie {
	res, err := sql.FindCate()
	if err != nil {
		log.Fatalln(err)
	}

	items := make([]opts.PieData, 0)
	for _, v := range res {
		items = append(items, opts.PieData{Name: v.Id, Value: v.Value})
	}

	pie := charts.NewPie()

	// 自定义颜色
	colors := []string{
		"#fc8251", "#5470c6", "#9A60B4", "#ef6567",
		"#f9c956", "#3BA272", "#E164B8", "#FF9F7F",
		"#37A2DA", "#FFDB5C",
	}

	pie.SetGlobalOptions(
		charts.WithColorsOpts(colors),
		charts.WithInitializationOpts(opts.Initialization{
			Width:  "1200px",
			Height: "500px",
		}),
		charts.WithTitleOpts(opts.Title{Title: "按类别统计新闻数量"}),
	)

	pie.AddSeries("pie", items).
		SetSeriesOptions(
			charts.WithLabelOpts(opts.Label{
				Show:      true,
				Formatter: "{b}: {c}",
			}),
			charts.WithPieChartOpts(opts.PieChart{
				Radius: []string{"40%", "75%"},
			}),
		)
	return pie
}

type PieExamples struct{}

func (PieExamples) Examples() {
	page := components.NewPage()
	page.AddCharts(
		pieRadius(),
	)
	f, err := os.Create("static/pie.html")
	if err != nil {
		panic(err)
	}
	page.Render(io.MultiWriter(f))
}

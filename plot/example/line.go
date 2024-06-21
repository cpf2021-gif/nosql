package example

import (
	"io"
	"log"
	"nosql/sql"
	"os"
	"slices"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func lineMarkPoint() *charts.Line {
	res, err := sql.FindTimeSeries()
	if err != nil {
		log.Fatalln(err)
	}
	slices.SortFunc(res, func(i, j sql.TimeSeries) int {
		return sql.CmpTime(i.Id, j.Id)
	})

	var time []string
	var values []opts.LineData

	for _, item := range res {
		time = append(time, item.Id)
		values = append(values, opts.LineData{Value: item.Value})
	}

	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{
			Width:  "1500px",
			Height: "700px",
		}),
		charts.WithXAxisOpts(opts.XAxis{
			AxisLabel: &opts.AxisLabel{
				Show:         true,
				Rotate:       45,
				ShowMinLabel: true,
				ShowMaxLabel: true,
			},
		}),
		charts.WithTitleOpts(opts.Title{
			Title: "不同时间的疫情数据",
		}),
		charts.WithTooltipOpts(opts.Tooltip{
			Show:    true,
			Trigger: "axis",
		}),
	)

	line.SetXAxis(time).AddSeries("时间", values).
		SetSeriesOptions(
			charts.WithLineChartOpts(opts.LineChart{
				ShowSymbol: true,
			}),
			charts.WithLabelOpts(opts.Label{
				Show: true,
			}),
		)
	return line
}

type LineExamples struct{}

func (LineExamples) Examples() {
	page := components.NewPage()
	page.AddCharts(
		lineMarkPoint(),
	)
	f, err := os.Create("static/line.html")
	if err != nil {
		panic(err)
	}
	page.Render(io.MultiWriter(f))
}

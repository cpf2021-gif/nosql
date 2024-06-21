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

func mapVisualMap(data []opts.MapData) *charts.Map {
	mc := charts.NewMap()
	mc.RegisterMapType("china")
	mc.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{
			Width:  "1500px",
			Height: "700px",
		},
		),
		charts.WithTitleOpts(opts.Title{
			Title: "VisualMap",
		}),
		charts.WithVisualMapOpts(opts.VisualMap{
			Calculable: true,
			Show:       true,
			Text:       []string{"High", "Low"},
			Min:        0,
			Max:        500,
		}),
		charts.WithTooltipOpts(opts.Tooltip{
			Show:    true,
			Trigger: "item",
		},
		),
	)

	mc.AddSeries("map", data).
		SetSeriesOptions(
			charts.WithLabelOpts(opts.Label{
				Show: true,
			}),
		)
	return mc
}

type MapExamples struct{}

func (MapExamples) Examples() {
	res, err := sql.FindProvince()
	if err != nil {
		log.Fatalln(err)
	}

	var data []opts.MapData
	for _, item := range res {
		if item.Id == "其他" {
			continue
		}

		data = append(data, opts.MapData{Name: item.Id, Value: float64(item.Value)})
	}

	page := components.NewPage()
	page.AddCharts(
		mapVisualMap(data),
	)

	f, err := os.Create("static/map.html")
	if err != nil {
		panic(err)
	}
	page.Render(io.MultiWriter(f))
}

package trend

import (
	"net/http"

	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/types"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	log "github.com/sirupsen/logrus"
)

func (t *Trend) WcBase() *charts.WordCloud {
	log.Debug()
	wc := charts.NewWordCloud()
	wc.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "실시간 검색어를 이용한 Trend 분석",
		}))

	wc.AddSeries("wordcloud", t.generateWCData(t.KeywordData))
	wc.SetSeriesOptions(
		charts.WithWorldCloudChartOpts(
			opts.WordCloudChart{
				SizeRange: []float32{10, 80},
			}),
	)
	return wc
}

func (t *Trend) generateWCData(data map[string]int) (items []opts.WordCloudData) {
	log.Debug()
	items = make([]opts.WordCloudData, 0)
	for k, v := range data {
		items = append(items, opts.WordCloudData{Name: k, Value: v})
	}
	return
}

func (t *Trend) graphBase() *charts.Graph {
	log.Debug()
	g := NewGraph("ternd graph build")
	for k, v := range t.KeywordData {
		node := opts.GraphNode{
			Name:  k,
			Value: float32(v),
		}
		g.AddNode(node)
	}
	for k1, v := range t.RelKeywordData {
		for k2, attr := range v {
			link := opts.GraphLink{
				Source: k1,
				Target: k2,
				Value:  float32(attr.Count),
			}
			g.AddLink(link)
		}
	}
	graph := charts.NewGraph()
	graph.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "basic graph example"}),
	)
	globalOptionInit := charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros, Width: "1800px", Height: "1000px"})
	globalOptionTielt := charts.WithTitleOpts(opts.Title{Title: "실시간 검색어를 이용한 Trend 분석"})
	graph.SetGlobalOptions(globalOptionInit, globalOptionTielt)
	graph.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "실시간 검색어를 이용한 Trend 분석",
		}))

	graph.AddSeries("graph", g.Node, g.Link,
		charts.WithGraphChartOpts(
			opts.GraphChart{
				Force:              &opts.GraphForce{Repulsion: 8000},
				FocusNodeAdjacency: true,
				Roam:               true,
			},
		),
		charts.WithLabelOpts(opts.Label{Show: true, Position: "right"}),
	)
	return graph
}

func (t *Trend) WcardHandler(w http.ResponseWriter, _ *http.Request) {
	wc := charts.NewWordCloud()
	globalOptionInit := charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros, Width: "1800px", Height: "1000px"})
	globalOptionTielt := charts.WithTitleOpts(opts.Title{Title: "실시간 검색어를 이용한 Trend 분석"})
	wc.SetGlobalOptions(globalOptionInit, globalOptionTielt)
	wc.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "실시간 검색어를 이용한 Trend 분석",
		}))

	wc.AddSeries("wordcloud", t.generateWCData(t.KeywordData))
	wc.SetSeriesOptions(
		charts.WithWorldCloudChartOpts(
			opts.WordCloudChart{
				SizeRange: []float32{10, 100},
			}),
	)
	wc.Render(w)
}

func (t *Trend) Handler(w http.ResponseWriter, _ *http.Request) {
	page := components.NewPage()
	page.AddCharts(
		t.graphBase(),
		//graphBar(),
		t.WcBase(),
	)
	page.Render(w)
}

package main

import (
	"math/rand"
	"net/http"
	"os"
	"strconv"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
	log.SetReportCaller(true)
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

// generate random data for line chart
func generateLineItems() []opts.LineData {
	items := make([]opts.LineData, 0)
	for i := 0; i < 7; i++ {
		items = append(items, opts.LineData{Value: rand.Intn(300)})
	}
	return items
}

func httpserver(w http.ResponseWriter, _ *http.Request) {
	// create a new line instance
	line := charts.NewLine()
	// set some global options like Title/Legend/ToolTip or anything else
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
		charts.WithTitleOpts(opts.Title{
			Title:    "Line example in Westeros theme",
			Subtitle: "Line chart rendered by the http server this time",
		}))

	// Put data into instance
	line.SetXAxis([]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}).
		AddSeries("Category A", generateLineItems()).
		AddSeries("Category B", generateLineItems()).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))
	line.Render(w)
}

func graphBar() *charts.Line {
	line := charts.NewLine()
	// set some global options like Title/Legend/ToolTip or anything else
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
		charts.WithTitleOpts(opts.Title{
			Title:    "Line example in Westeros theme",
			Subtitle: "Line chart rendered by the http server this time",
		}))

	// Put data into instance
	line.SetXAxis([]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}).
		AddSeries("Category A", generateLineItems()).
		AddSeries("Category B", generateLineItems()).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))
	return line
}

var graphNodes = []opts.GraphNode{}

func genNodes() []opts.GraphNode {
	graphNodes = []opts.GraphNode{}
	for i := 0; i < 100; i++ {
		node := opts.GraphNode{}
		node.Name = "node" + strconv.Itoa(i)
		graphNodes = append(graphNodes, node)
	}
	node := opts.GraphNode{}
	node.Name = "node100"
	graphNodes = append(graphNodes, node)
	return graphNodes
}

func genLink() []opts.GraphLink {
	links := make([]opts.GraphLink, 0)
	for i := 0; i < 100; i += 10 {
		l := opts.GraphLink{}
		l.Source = "node100"
		l.Target = graphNodes[i].Name
		links = append(links, l)
	}

	for i := 0; i < 50; i += 10 {
		for j := i + 1; j < i+10; j++ {
			l := opts.GraphLink{}
			l.Source = graphNodes[i].Name
			l.Target = graphNodes[j].Name
			links = append(links, l)
		}
	}
	return links
}

func genLinks() []opts.GraphLink {
	links := make([]opts.GraphLink, 0)
	for i := 0; i < len(graphNodes); i++ {
		for j := 0; j < len(graphNodes); j++ {
			links = append(links, opts.GraphLink{Source: graphNodes[i].Name, Target: graphNodes[j].Name})
		}
	}
	return links
}

func graphBase() *charts.Graph {
	graph := charts.NewGraph()
	graph.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "basic graph example"}),
	)
	graph.AddSeries("graph", genNodes(), genLink(),
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

func generateWCData(data map[string]interface{}) (items []opts.WordCloudData) {
	items = make([]opts.WordCloudData, 0)
	for k, v := range data {
		items = append(items, opts.WordCloudData{Name: k, Value: v})
	}
	return
}

var wcData = map[string]interface{}{
	"대한민국":  10000,
	"미국":    6181,
	"이탈리아":  4386,
	"그리스":   4055,
	"영국":    2467,
	"프랑스":   2244,
	"호주":    1898,
	"싱가포르":  1484,
	"인도네시아": 1689,
	"노르웨이":  1112,
	"덴마크":   985,
	"브라질":   847,
}

func wcBase() *charts.WordCloud {
	wc := charts.NewWordCloud()
	globalOptionInit := charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros, Width: "1800px", Height: "1000px"})
	globalOptionTitle := charts.WithTitleOpts(opts.Title{Title: "basic WordCloud example"})
	wc.SetGlobalOptions(globalOptionInit, globalOptionTitle)
	wc.AddSeries("wordcloud", generateWCData(wcData)).
		SetSeriesOptions(
			charts.WithWorldCloudChartOpts(
				opts.WordCloudChart{
					SizeRange: []float32{14, 80},
				}),
		)
	return wc
}

func GraphHandler(w http.ResponseWriter, _ *http.Request) {
	page := components.NewPage()
	page.AddCharts(
		graphBase(),
		//graphCircle(),
		graphBar(),
		wcBase(),
	)

	page.Render(w)
}

func main() {
	http.HandleFunc("/", GraphHandler)
	http.HandleFunc("/graph", GraphHandler)
	http.ListenAndServe(":8081", nil)
	log.Debug("running server at http://localhost:8081")
}

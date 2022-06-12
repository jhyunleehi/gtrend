package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

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

// 	{Name: "Node1"},
// 	{Name: "Node2"},
// 	{Name: "Node3"},
// 	{Name: "Node4"},
// 	{Name: "Node5"},
// 	{Name: "Node6"},
// 	{Name: "Node7"},
// 	{Name: "Node8"},
// }

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
func graphCircle() *charts.Graph {
	graph := charts.NewGraph()
	graph.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "Circular layout"}),
	)

	graph.AddSeries("graph", graphNodes, genLinks()).
		SetSeriesOptions(
			charts.WithGraphChartOpts(
				opts.GraphChart{
					Force:  &opts.GraphForce{Repulsion: 8000},
					Layout: "circular",
				}),
			charts.WithLabelOpts(opts.Label{Show: true, Position: "right"}),
		)
	return graph
}

func graphNpmDep() *charts.Graph {
	graph := charts.NewGraph()
	graph.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "npm dependencies demo",
		}))

	f, err := ioutil.ReadFile("fixtures/npmdepgraph.json")
	if err != nil {
		panic(err)
	}

	type Data struct {
		Nodes []opts.GraphNode
		Links []opts.GraphLink
	}

	var data Data
	if err := json.Unmarshal(f, &data); err != nil {
		fmt.Println(err)
	}

	graph.AddSeries("graph", data.Nodes, data.Links).
		SetSeriesOptions(
			charts.WithGraphChartOpts(opts.GraphChart{
				Layout:             "none",
				Roam:               true,
				FocusNodeAdjacency: true,
			}),
			charts.WithEmphasisOpts(opts.Emphasis{
				Label: &opts.Label{
					Show:     true,
					Color:    "black",
					Position: "left",
				},
			}),
			charts.WithLineStyleOpts(opts.LineStyle{
				Curveness: 0.3,
			}),
		)
	return graph
}

func GraphHandler(w http.ResponseWriter, _ *http.Request) {
	page := components.NewPage()
	page.AddCharts(
		graphBase(),
		//graphCircle(),
		//graphNpmDep(),
		graphBar(),
	)

	page.Render(w)
}

func main() {
	http.HandleFunc("/", GraphHandler)
	http.HandleFunc("/graph", GraphHandler)
	http.ListenAndServe(":8081", nil)
}

package main

import (
	"gtrend/trend"
	"gtrend/vendor/github.com/go-echarts/go-echarts/v2/components"
	"net/http"
)
var mytrend *trend.Trend

func GraphHandler(w http.ResponseWriter, _ *http.Request) {
	page := components.NewPage()
	page.AddCharts(
		//graphBase(),		
		//graphBar(),
		mytrend.WcBase(),
	)

	page.Render(w)
}

func main() {
	mytrend = trend.NewTrend("trend")
	go mytrend.Run()


}

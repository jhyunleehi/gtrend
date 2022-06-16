package trend

import (
	"net/http"

	"gtrend/vendor/github.com/go-echarts/go-echarts/v2/components"
	"gtrend/vendor/github.com/go-echarts/go-echarts/v2/types"

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

func (t *Trend) GraphHandler(w http.ResponseWriter, _ *http.Request) {
	page := components.NewPage()
	page.AddCharts(
		//graphBase(),
		//graphBar(),
		t.WcBase(),
	)
	page.Render(w)
}

func (t *Trend) GraphHandler1(w http.ResponseWriter, _ *http.Request) {
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
				SizeRange: []float32{10, 80},
			}),
	)
	wc.Render(w)
}

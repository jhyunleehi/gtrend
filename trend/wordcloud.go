package trend

import (
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
				SizeRange: []float32{14, 80},
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
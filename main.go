package main

import (
	"gtrend/trend"
	"time"
)

// func init() {
// 	log.SetFormatter(&log.TextFormatter{
// 		DisableColors: true,
// 		FullTimestamp: true,
// 	})
// 	log.SetReportCaller(true)
// 	log.SetOutput(os.Stdout)
// 	log.SetLevel(log.DebugLevel)
// }

func main() {
	t:=trend.NewTrend("trend")
	go t.Run()
	time.Sleep(120*time.Second)

}

package main

import (
	"gtrend/trend"
	"time"
)


func main() {
	t:=trend.NewTrend("trend")
	go t.Run()
	time.Sleep(120*time.Second)

}

package main

import (
	"fmt"
	"gtrend/trend"
	"net/http"
)

var mytrend *trend.Trend

func main() {
	mytrend = trend.NewTrend("trend")
	go mytrend.Run()
	http.HandleFunc("/", mytrend.GraphHandler)
	http.HandleFunc("/graph", mytrend.GraphHandler)
	http.ListenAndServe(":8081", nil)
	fmt.Printf("running server at http://localhost:8081")
}

package main

import (
	"fmt"
	"gtrend/trend"
	"net/http"
)

var mytrend *trend.Trend

func main() {
	mytrend = trend.NewTrend("trend")
	mytrend.GetInit()
	go mytrend.Run()
	http.HandleFunc("/", mytrend.Handler)
	http.HandleFunc("/wc", mytrend.WcardHandler)
	http.ListenAndServe(":8081", nil)
	fmt.Printf("running server at http://localhost:8081")
}

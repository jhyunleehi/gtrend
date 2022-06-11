package main

import (
	"fmt"
	"github.com/anaskhan96/soup"
	"os"
)

func main() {


	resp, err := soup.Get("https://awesome-ui.netlify.app/rank")
	if err != nil {
		os.Exit(1)
	}
	fmt.Printf("%s",resp)
	doc := soup.HTMLParse(resp)
	links := doc.FindAll("a","data-v-470f41c0")
	for _, link := range links {
		fmt.Println(link.Text(), "| Link :", link.Attrs()["href"])
	}
}

//<a data-v-470f41c0="" data-v-76c55539="" href="https://search.naver.com/search.naver?where=news&amp;sm=tab_jum&amp;query=그것이 알고싶다" target="_blank" class="rank-layer"><span data-v-470f41c0="" class="rank-num">1</span><span data-v-470f41c0="" class="rank-text">그것이 알고싶다</span><span data-v-470f41c0="" class="rank-icon"><i data-v-470f41c0="" class="fi-rr-minus-small"></i></span></a>




https://keyzard.org/realtimekeyword
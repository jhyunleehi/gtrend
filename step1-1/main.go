package main

import (
	//"fmt"
	"os"

	"github.com/anaskhan96/soup"
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

func main() {
	resp, err := soup.Get("https://awesome-ui.netlify.app/mzum")
	if err != nil {
		os.Exit(1)
	}
	//fmt.Printf("%s",resp)
	doc := soup.HTMLParse(resp)
	div := doc.FindAll("div", "class", "issue-keyword")
	for _, d := range div {
		links := d.FindAll("a")
		for _, link := range links {
			rankitem := link.Text()
			log.Debug(rankitem)
			//fmt.Println(link.Text(), "| Link :", link.Attrs()["href"])
		}
	}
}

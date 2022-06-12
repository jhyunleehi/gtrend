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
	resp, err := soup.Get("https://keyzard.org/realtimekeyword")
	if err != nil {
		os.Exit(1)
	}
	//fmt.Printf("%s",resp)
	doc := soup.HTMLParse(resp)
	div := doc.FindAll("div", "class", "col-sm-12")
	for _, d := range div {
		links := d.FindAll("a")
		for _, link := range links {
			rankitem := (link.Attrs()["title"])
			log.Debug(rankitem)
			//fmt.Println(link.Text(), "| Link :", link.Attrs()["href"])
		}
	}
}

type Keyword struct {
	RelKeyworkd        string `json:"relKeyword"`           //: "대한민국파라과이",
	MonthlyPcQcCn      int    `json:"monthlyPcQcCnt"`       //: 9880,
	MonthlyMobildQcCnt int    `json:"monthlyMobileQcCnt"`   //: 43900,
	Total              int    `json:"total"`                //: 15939,
	UpdateDate         int    `json:"updateDate,omitempty"` //: "2022-06-09 13:37:33",
	KeywordLevel       int    `json:"keywordLevel"`         //: 1,
}

type RelKeyword struct {
	List  []Keyword `json:"list,omitempty"`
	Goole []Keyword `json:"auto_google,omitempty"`
	Daum  []Keyword `json:"auto_daum,omitempty"`
	Naver []Keyword `json:"auto_naver,omitempy"`
}

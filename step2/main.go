package main

import (
	//"fmt"
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

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
	url := "https://m.some.co.kr/sometrend/analysis/composite/v2/association-transition"
	method := "POST"
	relkey := SearchKeyword{}
	searchkey := SearchKeyword{}
	searchkey.RelKeyworkd = "행복"
	searchkey.RequestQr = 10
	_, err := doRequest(method, url, &searchkey, &relkey)
	if err != nil {
		log.Error(err)
		return
	}

	for _, key := range relkey.Goole {
		log.Debug(key.RelKeyworkd, key.Total, key.UpdateDate)
	}
	for _, key := range relkey.Daum {
		log.Debug(key.RelKeyworkd, key.Total, key.UpdateDate)
	}
	for _, key := range relkey.Naver {
		log.Debug(key.RelKeyworkd, key.Total, key.UpdateDate)
	}
}

//doRequest ...
func doRequest(method, url string, in, out interface{}) (http.Header, error) {
	log.Debugf("[%+v] [%+v]", method, url)
	var inbody []byte
	var body *bytes.Buffer
	var req *http.Request
	if in != nil {
		inbody, _ = json.Marshal(in)
		body = bytes.NewBuffer(inbody)
		req, _ = http.NewRequest(method, url, body)
	} else {
		req, _ = http.NewRequest(method, url, nil)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("accept", "application/json")
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := http.Client{Transport: tr}
	resp, errReq := client.Do(req)

	if errReq != nil {
		if resp == nil {
			log.Error(errReq)
			return nil, errReq
		}
		buf, _ := ioutil.ReadAll(resp.Body)
		msg := fmt.Sprintf("doRequest() error: [%+v] [%+v]", errReq, string(buf))
		log.Errorf(msg)
		return nil, errReq

	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		buf, _ := ioutil.ReadAll(resp.Body)
		msg := fmt.Sprintf("[%+v]", strings.Replace(string(buf), "\n", " ", 999))
		log.Errorf(msg)
		return nil, errors.New(msg)

	}
	buf, _ := ioutil.ReadAll(resp.Body)

	json.Unmarshal([]byte(buf), out)
	return resp.Header, nil
}

type SearchKeyword struct {
	StartDate              string `json:"startDate"`             //: "20220512",
	EndDate                string `json:"endDate"`               //": "20220611",
	TopN                   int    `json:"topN"`                  //": 500,
	Period                 string `json:"period"`                //": "1",
	AnalysisMonths         int    `json:"analysisMonths"`        //": 0,
	CategorySetName        string `json:"categorySetName"`       //": "SMT",
	Sources                string `json:"sources"`               //": "blog,news,twitter",
	Keyword                string `json:"keyword"`               //": "이정현",
	Synonym                string `json:"synonym"`               //": null,
	KeywordFilterIncludes  string `json:"keywordFilterIncludes"` //": null,
	KeyworkdFilterExcludes string `json:"keywordFilterExclude"`  //s": null,
	IncludeWordOperatros   string `json:"includeWordOperators"`  //": "||",
	ExcludeWordOperators   string `json:"excludeWordOperators"`  //": "||",
	ScoringKeyWord         string `json:"scoringKeyword"`        //": "",
	ExForHash              string `json:"exForHash"`             //": "",
	CategoryList           string `json:"categoryList"`          //": "politician,celebrity,sportsman,characterEtc,government,business,agency,groupEtc,tourism,restaurant,shopping,scene,placeEtc,brandFood,cafe,brandBeverage,brandElectronics,brandFurniture,brandBeauty,brandFashion,brandEtc,productFood,productBeverage,productElectronics,productFurniture,productBeauty,productFashion,productEtc,economy,social,medicine,education,culture,sports,cultureEtc,animal,plant,naturalPhenomenon,naturalEtc"
}

type Association struct {
	Item   Item   `json:"item"`
	Code   string `json:"code"`
	Errors string `json:"errors"`
	Error  string `json:"error"`
}

type Item struct {
	DataList []DataList `json:"dataList"`
	Keyword  string     `json:"keyworkd"`
}
type DataList struct {
	Source string `json:"source"`
	Data   Data   `json:"data"`
}
type Data struct {
	Rows        []Row  `json:"rows"`
	CategoryMap string `json:"categoryMap"`
}
type Row struct {
	date            string            `json:"date"`
	AssociationData []AssociationData `json:"associationData"`
}
type AssociationData struct {
	Label     string `json:"label"`
	Frequency int    `json:"frequency"`
}

package main

import (
	//"fmt"
	"bytes"
	//"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

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
	rkey := Association{}

	ntime := time.Now()
	today := ntime.Format("20060102")
	yesterday := ntime.AddDate(0, 0, -1).Format("20060102")

	skey := SearchKeyword{}
	skey.StartDate = yesterday
	skey.EndDate = today
	skey.TopN = 500
	skey.Period = "1"
	skey.AnalysisMonths = 0
	skey.CategorySetName = "SMT"
	skey.Sources = "blog,news,twitter"
	skey.Keyword = "행복"
	//skey.Synonym = ""
	//skey.KeywordFilterIncludes = ""
	//skey.KeyworkdFilterExcludes = ""
	skey.IncludeWordOperatros = "||"
	skey.ExcludeWordOperators = "||"
	skey.ScoringKeyWord = "행복"
	skey.ExForHash = ""
	skey.CategoryList = "politician,celebrity,sportsman,characterEtc,government,business,agency,groupEtc,tourism,restaurant,shopping,scene,placeEtc,brandFood,cafe,brandBeverage,brandElectronics,brandFurniture,brandBeauty,brandFashion,brandEtc,productFood,productBeverage,productElectronics,productFurniture,productBeauty,productFashion,productEtc,economy,social,medicine,education,culture,sports,cultureEtc,animal,plant,naturalPhenomenon,naturalEtc"

	_, err := doRequest(method, url, &skey, &rkey)
	if err != nil {
		log.Error(err)
		return
	}

	for i, data := range rkey.Item.DataList {
		for j, rows := range data.Data.Rows {
			for k, ass := range rows.AssociationData {
				log.Debugf("ITEM [%d][%d][%d] [%s] [%d] ", i, j, k, ass.Label, ass.Frequency)
			}
		}

	}

}

//doRequest ...
func doRequest(method, url string, in, out interface{}) (http.Header, error) {
	log.Debugf("[%+v] [%+v] [%+v]", method, url, in)
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
	req.Header.Add("Accept", "application/json")
	// tr := &http.Transport{
	// 	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	// }
	//client := http.Client{Transport: tr}
	client := http.Client{
		//Transport: tr,
		Timeout: 30 * time.Second,
	}
	log.Debug(client)

	//client1 := http.Client{	}
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
	StartDate              string `json:"startDate"`                       //: "20220512",
	EndDate                string `json:"endDate"`                         //": "20220611",
	TopN                   int    `json:"topN"`                            //": 500,
	Period                 string `json:"period"`                          //": "1",
	AnalysisMonths         int    `json:"analysisMonths"`                  //": 0,
	CategorySetName        string `json:"categorySetName"`                 //": "SMT",
	Sources                string `json:"sources"`                         //": "blog,news,twitter",
	Keyword                string `json:"keyword"`                         //": "이정현",
	Synonym                string `json:"synonym,omitempty"`               //": null,
	KeywordFilterIncludes  string `json:"keywordFilterIncludes,omitempty"` //": null,
	KeyworkdFilterExcludes string `json:"keywordFilterExclude,omitempty"`  //s": null,
	IncludeWordOperatros   string `json:"includeWordOperators"`            //": "||",
	ExcludeWordOperators   string `json:"excludeWordOperators"`            //": "||",
	ScoringKeyWord         string `json:"scoringKeyword"`                  //": "",
	ExForHash              string `json:"exForHash"`                       //": "",
	CategoryList           string `json:"categoryList"`                    //": "politician,celebrity,sportsman,characterEtc,government,business,agency,groupEtc,tourism,restaurant,shopping,scene,placeEtc,brandFood,cafe,brandBeverage,brandElectronics,brandFurniture,brandBeauty,brandFashion,brandEtc,productFood,productBeverage,productElectronics,productFurniture,productBeauty,productFashion,productEtc,economy,social,medicine,education,culture,sports,cultureEtc,animal,plant,naturalPhenomenon,naturalEtc"
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

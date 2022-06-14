package trend

import (
	"bytes"
	"sync"

	//"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/anaskhan96/soup"
	log "github.com/sirupsen/logrus"
)

type Trend struct {
	KeywordData    map[string]Attrs            //실검
	RelKeywordData map[string]map[string]Attrs //연관검색어
	Keyword        map[string]Attrs            //실검
	RelKeyword     map[string]map[string]Attrs //연관검색어
	done           chan struct{}
	mutex          sync.Mutex
}

type Attrs struct {
	Source string
	Count  int
}

func NewTrend(name string) *Trend {
	log.Debugf("%s", name)
	trend := Trend{
		Keyword:    make(map[string]Attrs),
		RelKeyword: make(map[string]map[string]Attrs),
	}
	return &trend
}

func (t *Trend) AddKeyword(name string, attr Attrs) error {
	log.Debugf("[%s][%v]", name, attr)
	if _, exists := t.Keyword[name]; !exists {
		t.Keyword[name] = attr
	} else {
		k := t.Keyword[name]
		a := Attrs{
			Count:  k.Count + attr.Count,
			Source: attr.Source,
		}
		t.Keyword[name] = a
	}
	return nil
}

func (t *Trend) AddRelKeyword(from string, to string, A Attrs) error {
	log.Debugf("[%s] [%s] [%+v]", from, to, A)
	if _, exists := t.RelKeyword[from]; !exists {
		t.RelKeyword[from] = make(map[string]Attrs)
		t.RelKeyword[from][to] = A
	} else {
		t.RelKeyword[from][to] = A
	}
	return nil
}

func (t *Trend) PrintKeywordData() error {
	for key, val := range t.Keyword {
		log.Debugf("[%s][%s]", key, val)
	}
	for k1, v1 := range t.RelKeyword {
		for k2, v2 := range v1 {
			log.Debugf("[%s]-->[%s] [%s][%d]", k1, k2, v2.Source, v2.Count)
		}
	}
	return nil
}

// Run starts countbeat.
func (t *Trend) Run() error {
	log.Debug("running get keyword...")
	tickerGetKeyWord := time.NewTicker(conf.Collect)
	for {
		select {
		case <-t.done:
			log.Debug("get ticker get bt.done")
			return nil
		case <-tickerGetKeyWord.C:
			log.Info("ticker=> GetKeyWord")
			t.KeywordData = t.Keyword
			t.RelKeywordData = t.RelKeyword
			t.Keyword = map[string]Attrs{}
			t.RelKeyword = map[string]map[string]Attrs{}
			err := t.GetRealTimeKeyword1() //keyzard
			if err != nil {
				log.Error(err)
				continue
			}
			err = t.GetRealTimeKeyword2() //mzum
			if err != nil {
				log.Error(err)
				continue
			}
			err = t.GetRelKeyword()
			if err != nil {
				log.Error(err)
				continue
			}
			continue
		}
	}
}

// Stop stops countbeat.
func (t *Trend) Stop() {
	close(t.done)
}

func (t *Trend) GetRealTimeKeyword1() error {
	log.Debug()
	resp, err := soup.Get("https://keyzard.org/realtimekeyword")
	if err != nil {
		log.Error(err)
		return err
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
			attr := Attrs{}
			attr.Count = 0
			attr.Source = ""
			t.AddKeyword(rankitem, attr)
		}
	}
	return nil
}

func (t *Trend) GetRealTimeKeyword2() error {
	log.Debug()
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
			label := strings.ReplaceAll(rankitem, " ", "")
			log.Debug(label)
			attr := Attrs{}
			attr.Count = 0
			attr.Source = ""
			t.AddKeyword(label, attr)
		}
	}
	return nil
}

func (t *Trend) GetRelKeyword() error {
	log.Debug()
	for key, _ := range t.Keyword {
		err := t.GetRelKeywordItem(key)
		if err != nil {
			log.Error(err)
			return err
		}
	}
	return nil
}

func (t *Trend) GetRelKeywordItem(searchword string) error {
	log.Debugf("[%s]", searchword)
	url := "https://m.some.co.kr/sometrend/analysis/composite/v2/association-transition"
	method := "POST"
	rkey := Association{}
	ntime := time.Now()
	today := ntime.Format("20060102")
	yesterday := ntime.AddDate(0, 0, -1).Format("20060102")
	skey := SearchKeyword{}
	skey.StartDate = yesterday
	skey.EndDate = today
	skey.TopN = 100 //500
	skey.Period = "1"
	skey.AnalysisMonths = 0
	skey.CategorySetName = "SMT"
	skey.Sources = "blog,news,twitter"
	skey.Keyword = searchword
	skey.Synonym = ""
	skey.KeywordFilterIncludes = ""
	skey.KeyworkdFilterExcludes = ""
	skey.IncludeWordOperatros = "||"
	skey.ExcludeWordOperators = "||"
	skey.ScoringKeyWord = ""
	skey.ExForHash = ""
	skey.CategoryList = "politician,celebrity,sportsman,characterEtc,government,business,agency,groupEtc,tourism,restaurant,shopping,scene,placeEtc,brandFood,cafe,brandBeverage,brandElectronics,brandFurniture,brandBeauty,brandFashion,brandEtc,productFood,productBeverage,productElectronics,productFurniture,productBeauty,productFashion,productEtc,economy,social,medicine,education,culture,sports,cultureEtc,animal,plant,naturalPhenomenon,naturalEtc"

	_, err := t.doRequest(method, url, &skey, &rkey)
	if err != nil {
		log.Error(err)
		return err
	}

	for i, data := range rkey.Item.DataList {
		for j, rows := range data.Data.Rows {
			for k, ass := range rows.AssociationData {
				label := strings.ReplaceAll(ass.Label, " ", "")
				log.Debugf("ITEM [%d][%d][%d] [%s] [%d] ", i, j, k, label, ass.Frequency)
				attr := Attrs{}
				attr.Count = ass.Frequency
				attr.Source = data.Source
				t.AddKeyword(label, attr)
				t.AddRelKeyword(searchword, label, attr)
			}
		}
	}
	return nil
}

//doRequest ...
func (t *Trend) doRequest(method, url string, in, out interface{}) (http.Header, error) {
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

	// tr := &http.Transport{
	// 	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	// }
	//client := http.Client{Transport: tr}

	client := http.Client{
		Timeout: 30 * time.Second,
	}
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

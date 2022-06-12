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
	url := "https://keyzard.org/query/searchs"
	method := "POST"
	relkey := RelKeyword{}
	searchkey := SearchKeyworkd{}
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

type RelKeyword struct {
	List  []Keyword `json:"list,omitempty"`
	Goole []Keyword `json:"auto_google,omitempty"`
	Daum  []Keyword `json:"auto_daum,omitempty"`
	Naver []Keyword `json:"auto_naver,omitempty"`
}

type Keyword struct {
	RelKeyworkd        string `json:"relKeyword"`           //: "대한민국파라과이",
	MonthlyPcQcCn      int    `json:"monthlyPcQcCnt"`       //: 9880,
	MonthlyMobildQcCnt int    `json:"monthlyMobileQcCnt"`   //: 43900,
	Total              int    `json:"total"`                //: 15939,
	UpdateDate         int    `json:"updateDate,omitempty"` //: "2022-06-09 13:37:33",
	KeywordLevel       int    `json:"keywordLevel"`         //: 1,
}

type SearchKeyworkd struct {
	RelKeyworkd string `json:"relKeyword"`
	RequestQr   int    `json:"request_rq"`
}

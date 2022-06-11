package main

import (
	"crypto/tls"
	"fmt"
	//"io/ioutil"
	"log"
	"net/http"
	//"strings"

	"github.com/PuerkitoBio/goquery"
)

func ExampleScrape() {
	// Request the HTML page.
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	res, err := client.Get("https://keyzard.org/realtimekeyword")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
	//buf, _ := ioutil.ReadAll(res.Body)
	//fmt.Printf("%+v", string(buf))
	// Load the HTML document
	//rHtml := strings.NewReader(string(buf))
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v", doc)

	// Find the review items
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		//s.Find("button").Each(func(i int, t *goquery.Selection) {
			band := s.Text()
			title := s.Text()
			fmt.Printf("Review %d: %s - %s\n", i, band, title)
		//})cd .

	})
}

func main() {
	ExampleScrape()
}

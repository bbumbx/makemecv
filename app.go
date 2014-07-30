package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"	
	"net/url"
)

func main() {
	location := "California"
	salary := "$100,000+"
	title := "java"
	format := "json"
	// https://ads.indeed.com/jobroll/xmlfeed
	u, err := url.Parse("http://api.indeed.com/ads/apisearch?publisher=5068861968542440")
	if err != nil {
		log.Fatal(err)
	}
	q := u.Query()
	q.Set("q", "title:" + title + ", " + salary + " jobs")
	q.Set("l", location)
	q.Set("format", format)
	q.Set("userip", "1.2.3.4")
	q.Set("useragent", "Mozilla")
	q.Set("v", "2")
	q.Set("highlight", "0")
	u.RawQuery = q.Encode()
	fmt.Println(u.String())
	res, err := http.Get(u.String())
	if err != nil {
		log.Fatal(err)
	}
	robots, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)
}
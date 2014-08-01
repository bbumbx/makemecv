package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"	
	"net/url"
	"encoding/json"
	// "reflect"
)



func main() {
	location := "California"
	salary := "$100,000+"
	title := "java"
	format := "json"
	limit := "1"
	// https://ads.indeed.com/jobroll/xmlfeed
	u, err := url.Parse("http://api.indeed.com/ads/apisearch?publisher=5068861968542440")
	if err != nil {
		log.Fatal(err)
	}
	q := u.Query()
	q.Set("q", "title:" + title + ", " + salary + " jobs")
	q.Set("l", location)
	q.Set("format", format)
	q.Set("limit", limit)
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
	// fmt.Printf("%s", robots)
	urls := make([]string, 1)
	// var urls []string
	var r map[string] interface{}

	err = json.Unmarshal(robots, &r)
	switch res_arr := r["results"].(type) {
       case []interface{}:
       	for i,element := range res_arr {
       		switch obj := element.(type) {
       			case map[string] interface{}:
       				urls[i] = obj["url"].(string)
       				// fmt.Println(obj["url"])
       		}
  			
		}
    }
    // for _, u := range urls{
    // 	fmt.Println(u)
    // }
    loadBodies(urls)
	//t := reflect.TypeOf(r)
	//fmt.Println(t)
	
	// if err != nil{
	// 	panic(err)
	// } else {
	// 	fmt.Println(res_arr)	
	// }
}

func loadBodies(urls []string) {
	for _, url := range urls {
		res, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		page, err := ioutil.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s", page)
	}	
}
package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

const sitemap string = "https://www.washingtonpost.com/news-sitemap-index.xml"

type SitemapIndex struct {
	Locations []string `xml:"sitemap>loc"`
}

type News struct {
	Titles    []string `xml:"url>news>title"`
	Keywords  []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

type NewsMap struct {
	Keyword  string
	Location string
}

// type Location struct {
// 	Loc string `xml:"loc"`
// }
// func (l Location) String() string {
// 	return fmt.Sprintf(l.Loc)
// }

func main() {
	var s SitemapIndex
	var n News
	news_map := make(map[string]NewsMap)

	resp, _ := http.Get(sitemap)
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	xml.Unmarshal(bytes, &s)
	// fmt.Println(s.Locations)
	// for i := 0; i < 10; i++ {
	// }

	// range returns two values the index and the value
	for _, Location := range s.Locations {
		// fmt.Println(Location)
		resp, _ := http.Get(Location)
		bytes, _ := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		xml.Unmarshal(bytes, &n)
		for idx, _ := range n.Titles {
			news_map[n.Titles[idx]] = NewsMap{n.Keywords[idx], n.Locations[idx]}
		}
	}

	for idx, data := range news_map {
		fmt.Println("\n\n\n", idx)
		fmt.Println("\n", data.Keyword)
		fmt.Println("\n", data.Location)
	}
}

// Using maps
// mapping := make(map[string]float32)
// mapping["Garrett"] = 44
// mapping["Test"] = 90
// delete(mapping, "Test")
// for k, v := range mapping {

// }

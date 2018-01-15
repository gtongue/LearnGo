package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

const sitemap string = "https://www.washingtonpost.com/news-sitemap-index.xml"

type SitemapIndex struct {
	Locations []Location `xml:"sitemap"`
}

type Location struct {
	Loc string `xml:"loc"`
}

func (l Location) String() string {
	return fmt.Sprintf(l.Loc)
}

func main() {
	resp, _ := http.Get(sitemap)
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	var s SitemapIndex
	xml.Unmarshal(bytes, &s)
	fmt.Println(s.Locations)
}

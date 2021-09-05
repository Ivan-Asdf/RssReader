package main

import (
	"fmt"

	"github.com/Ivan-Asdf/RssReader/reader"
)

func main() {
	urls := []string{
		"https://www.rssboard.org/files/sample-rss-091.xml",
		"https://www.rssboard.org/files/sample-rss-092.xml",
		"https://www.rssboard.org/files/sample-rss-2.xml",
	}
	_, errors := reader.Parse(urls)
	if len(errors) != 0 {
		fmt.Println(errors)
	}
	// fmt.Println(results)
}

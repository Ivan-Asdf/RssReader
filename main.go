package main

import (
	"fmt"

	"github.com/Ivan-Asdf/RssReader/reader"
)

func main() {
	// err := reader.Parse("https://www.rssboard.org/files/sample-rss-092.xml")
	err := reader.Parse("https://www.rssboard.org/files/sample-rss-2.xml")
	if err != nil {
		fmt.Println(err)
	}
}

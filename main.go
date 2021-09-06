package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/Ivan-Asdf/RssReader/reader"
)

func printJson(rssItems []reader.RssItem) {
	jsonString, err := json.MarshalIndent(rssItems, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(jsonString))

	file, err := os.OpenFile("output.json", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	bytesWritten, err := file.Write(jsonString)
	if err != nil {
		log.Fatal(err)
	} else if bytesWritten != len(jsonString) {
		log.Println("Not all bytes written to file")
	}
	err = file.Close()
	if err != nil {
		log.Println(err)
	}
}

func main() {
	urls := []string{
		"https://www.rssboard.org/files/sample-rss-091.xml",
		"https://www.rssboard.org/files/sample-rss-092.xml",
		"https://www.rssboard.org/files/sample-rss-2.xml",
	}
	results, errors := reader.Parse(urls)
	if len(errors) != 0 {
		log.Println(errors)
	}
	printJson(results)
}

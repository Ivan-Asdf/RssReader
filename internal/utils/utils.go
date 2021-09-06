package utils

import (
	"encoding/json"
	"flag"
	"log"
	"os"

	"github.com/Ivan-Asdf/RssReader/pkg/reader"
)

func PrintJson(rssItems []reader.RssItem) {
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

func GetInput() []string {
	args := os.Args[1:]
	if len(args) == 0 {
		log.Fatal("Please specify one or more urls as arguments")
	}

	runTestUrlsFlag := flag.Bool("t", false, "If flag is specified will run use predifined test urls")
	flag.Parse()
	if *runTestUrlsFlag {
		return []string{
			"https://www.rssboard.org/files/sample-rss-091.xml",
			"https://www.rssboard.org/files/sample-rss-092.xml",
			"https://www.rssboard.org/files/sample-rss-2.xml",
		}
	} else {
		return args
	}
}

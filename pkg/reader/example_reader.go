package reader

import "fmt"

func parseExample() {
	var rssItems []RssItem
	var errors []error

	rssItems, errors = Parse([]string{
		"https://www.rssboard.org/files/sample-rss-091.xml",
		"https://www.rssboard.org/files/sample-rss-092.xml",
	})

	if len(errors) != 0 {
		fmt.Println(errors)
	}
	for _, rssItem := range rssItems {
		fmt.Printf("%v\n\n", rssItem)
	}
}

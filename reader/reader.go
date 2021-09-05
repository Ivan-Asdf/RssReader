package reader

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

type RssItem struct {
	Title       string `xml:"title"`
	Source      string `xml:"source"`
	SourceURL   string
	Link        string `xml:"link"`
	PublishDate string `xml:"pubDate"`
	Description string `xml:"description"`
}

func getRssItems(xmlReader io.ReadCloser, itemTag string) ([]RssItem, error) {
	rssItems := make([]RssItem, 0)

	decoder := xml.NewDecoder(xmlReader)
	for {
		t, err := decoder.Token()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, err
			}
		}
		if t, ok := t.(xml.StartElement); ok {
			if t.Name.Local == itemTag {
				var item RssItem
				err = decoder.DecodeElement(&item, &t)
				if err != nil {
					return nil, err
				}
				rssItems = append(rssItems, item)
			}
		}
	}

	return rssItems, nil
}

func printJson(object interface{}) {
	jsonString, _ := json.MarshalIndent(object, "", "  ")
	fmt.Println(string(jsonString))
}

func Parse(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	rssItems, err := getRssItems(resp.Body, "item")
	if err != nil {
		fmt.Println(err)
	}
	printJson(rssItems)

	return nil
}

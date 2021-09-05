package reader

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"time"
)

type rawRssItem struct {
	Title       string     `xml:"title"`
	Source      itemSource `xml:"source"`
	Link        string     `xml:"link"`
	PublishDate string     `xml:"pubDate"`
	Description string     `xml:"description"`
}

type itemSource struct {
	XMLName   xml.Name
	SourceURL string `xml:"url,attr"`
}

type RssItem struct {
	Title       string     `json:"title"`
	Source      string     `json:"source"`
	SourceUrl   string     `json:"sourceUrl"`
	Link        string     `json:"link"`
	PublishDate *time.Time `json:"pubDate"`
	Description string     `json:"description"`
}

const ITEM_TAG = "item"

func getRawRssItems(xmlReader io.ReadCloser) ([]rawRssItem, error) {
	rssItems := make([]rawRssItem, 0)

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
			if t.Name.Local == ITEM_TAG {
				var item rawRssItem
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

func getDateFormats() []string {
	return []string{
		time.RFC822,
		"Mon, 02 Jan 2006 15:04:05 GMT",
	}
}

func getRssItems(rawItems []rawRssItem) []RssItem {
	rssItems := make([]RssItem, 0)
	for _, v := range rawItems {
		var rssItem RssItem
		rssItem.Title = v.Title
		rssItem.Source = v.Source.XMLName.Local
		rssItem.SourceUrl = v.Source.SourceURL
		rssItem.Link = v.Link

		for _, format := range getDateFormats() {
			time, err := time.Parse(format, v.PublishDate)
			if err == nil {
				rssItem.PublishDate = &time
				break
			}
		}

		rssItem.Description = v.Description
		rssItems = append(rssItems, rssItem)
	}
	return rssItems
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

	rssItems, err := getRawRssItems(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	printJson(getRssItems(rssItems[:2]))

	return nil
}

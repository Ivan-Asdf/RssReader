package reader

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"golang.org/x/net/html/charset"
)

// Necessary for parsing xml
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

// The final product of this package
type RssItem struct {
	Title       string     `json:"title"`
	Source      string     `json:"source"`
	SourceUrl   string     `json:"sourceUrl"`
	Link        string     `json:"link"`
	PublishDate *time.Time `json:"pubDate"`
	Description string     `json:"description"`
}

const ITEM_TAG = "item"

// Parse xml and return rss items
func getRawRssItems(xmlReader io.ReadCloser) ([]rawRssItem, error) {
	rssItems := make([]rawRssItem, 0)

	decoder := xml.NewDecoder(xmlReader)
	// Add ISO-8859-1 encoding support
	decoder.CharsetReader = charset.NewReaderLabel
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

// Supported date formats. If more are needed they can be added
func getDateFormats() []string {
	return []string{
		time.RFC822,
		"Mon, 02 Jan 2006 15:04:05 GMT",
	}
}

// Convert rawRssItems to RssItems
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

// Get and parse a single rss url
func processUrl(url string, wg *sync.WaitGroup, resultChan chan<- []RssItem, errorChan chan<- error) {
	resp, err := http.Get(url)
	if err != nil {
		errorChan <- err
		wg.Done()
		return
	}
	rssItems, err := getRawRssItems(resp.Body)
	if err != nil {
		errorChan <- err
		wg.Done()
		return
	}
	resultChan <- getRssItems(rssItems[:2])
	wg.Done()
}

// Given a slice of urls will return all RssItems extracted from them
func Parse(urls []string) ([]RssItem, []error) {
	resultsChan := make(chan []RssItem)
	errorsChan := make(chan error)

	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go processUrl(url, &wg, resultsChan, errorsChan)
	}

	done := make(chan interface{})
	go func() {
		wg.Wait()
		close(done)
	}()

	results := make([]RssItem, 0)
	errors := make([]error, 0)
loop:
	for {
		select {
		case rssItems := <-resultsChan:
			results = append(results, rssItems...)
			printJson(rssItems[:2])
		case err := <-errorsChan:
			errors = append(errors, err)
		case <-done:
			break loop
		}
	}

	return results, errors
}

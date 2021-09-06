package reader

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRawRssItems(t *testing.T) {
	xmlFile, err := os.Open("testdata/test.xml")
	if err != nil {
		t.Error(err)
	}

	rawRssItems, err := getRawRssItems(xmlFile)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 3, len(rawRssItems))

	assert.Equal(t, "Article1", rawRssItems[0].Title)
	assert.Equal(t, "", rawRssItems[0].Source.Source)
	assert.Equal(t, "", rawRssItems[0].Source.SourceURL)
	assert.Equal(t, "", rawRssItems[0].Link)
	assert.Equal(t, "", rawRssItems[0].PublishDate)
	assert.Equal(t, "", rawRssItems[0].Description)

	assert.Equal(t, "Article2", rawRssItems[1].Title)
	assert.Equal(t, "ABC blog", rawRssItems[1].Source.Source)
	assert.Equal(t, "www.abc.com", rawRssItems[1].Source.SourceURL)
	assert.Equal(t, "www.xyz.com", rawRssItems[1].Link)
	assert.Equal(t, "09 Aug 20 19:22 EET", rawRssItems[1].PublishDate)
	assert.Equal(
		t,
		`Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.`,
		rawRssItems[1].Description)

	assert.Equal(t, "Article3", rawRssItems[2].Title)
	assert.Equal(t, "QWE blog", rawRssItems[2].Source.Source)
	assert.Equal(t, "www.qwe.com", rawRssItems[2].Source.SourceURL)
	assert.Equal(t, "", rawRssItems[2].Link)
	assert.Equal(t, "", rawRssItems[2].PublishDate)
	assert.Equal(t, "", rawRssItems[2].Description)
}

func TestGetRssItems(t *testing.T) {
	xmlFile, err := os.Open("testdata/test.xml")
	if err != nil {
		t.Error(err)
	}

	rawRssItems, err := getRawRssItems(xmlFile)
	if err != nil {
		t.Error(err)
	}
	rssItems := getRssItems(rawRssItems)

	assert.Equal(t, 3, len(rssItems))

	assert.Equal(t, "Article1", rssItems[0].Title)
	assert.Equal(t, "", rssItems[0].Source)
	assert.Equal(t, "", rssItems[0].SourceURL)
	assert.Equal(t, "", rssItems[0].Link)
	assert.Nil(t, rssItems[0].PublishDate)
	assert.Equal(t, "", rssItems[0].Description)

	assert.Equal(t, "Article2", rssItems[1].Title)
	assert.Equal(t, "ABC blog", rssItems[1].Source)
	assert.Equal(t, "www.abc.com", rssItems[1].SourceURL)
	assert.Equal(t, "www.xyz.com", rssItems[1].Link)
	assert.NotNil(t, rssItems[1].PublishDate)
	assert.Equal(
		t,
		`Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.`,
		rssItems[1].Description)

	assert.Equal(t, "Article3", rssItems[2].Title)
	assert.Equal(t, "QWE blog", rssItems[2].Source)
	assert.Equal(t, "www.qwe.com", rssItems[2].SourceURL)
	assert.Equal(t, "", rssItems[2].Link)
	assert.Nil(t, rssItems[2].PublishDate)
	assert.Equal(t, "", rssItems[2].Description)
}

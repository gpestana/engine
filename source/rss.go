package source

import (
	"github.com/gpestana/engine/data"
	"github.com/mmcdole/gofeed"
	"log"
	"time"
)

type Rss struct {
	Type string
	Url  string
}

type Parent struct {
	Title       string
	Url         string
	Description string
}

// Fetches data from RSS feed and creates a data.DataUnit based on the RSS items
// returned.
func (r Rss) Fetch() []data.DataUnit {
	fp := gofeed.NewParser()
	f, err := fp.ParseURL(r.Url)
	if err != nil {
		log.Println("source.Fetch: " + err.Error())
		return []data.DataUnit{}
	}

	units := genDataUnits(f)
	return units
}

func genDataUnits(feed *gofeed.Feed) []data.DataUnit {
	units := []data.DataUnit{}
	parent := Parent{
		Title:       feed.Title,
		Url:         feed.Link,
		Description: feed.Description,
	}

	for _, i := range feed.Items {
		u := data.DataUnit{
			Url:            i.Link,
			Timestamp:      time.Now(),
			Parent:         &parent,
			FetchedContent: i,
			Description:    i.Description,
		}
		units = append(units, u)
	}

	return units
}

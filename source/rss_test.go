package source

import (
	"fmt"
	"github.com/mmcdole/gofeed"
	"testing"
)

func TestParseDataUnits(t *testing.T) {
	mockItems := []*gofeed.Item{
		&gofeed.Item{
			Link:        "https://link1.test",
			Description: "Test item 1",
		},
		&gofeed.Item{
			Link:        "https://link2.test",
			Description: "Test item 2",
		},
	}

	mockFeed := gofeed.Feed{
		Title:       "test feed",
		Link:        "https://feed.test",
		Description: "Test feed",
		Items:       mockItems,
	}

	units := genDataUnits(&mockFeed)
	if l := len(units); l != 2 {
		t.Fatal(
			fmt.Sprintf("Lenght of data units is %v but expected 2", l))
	}

	if u := units[0].Url; u != "https://link1.test" {
		t.Error(
			fmt.Sprintf("Data unit Url is %v but expected https://link1.test", u))
	}

	if d := units[1].Description; d != "Test item 2" {
		t.Error(
			fmt.Sprintf("Data unit Description is %v but expected 'Test item 2'", d))
	}

	if p := units[0].Parent; p == nil {
		t.Error("Data unit Parent is nil")
	}
}

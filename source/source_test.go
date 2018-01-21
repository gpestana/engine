package source

import (
	"testing"
)

func TestNew(t *testing.T) {
	tp := "rss"
	u := "UrlA"
	c := make(chan struct{})

	src := New(tp, u, c)
	rss := src.(Rss)

	if rss.Type != tp {
		t.Error("rss.Type should be 'rss', instead:" + rss.Type)
	}

	if rss.Url != u {
		t.Error("rss.Url should be 'UrlA', instead:" + rss.Url)
	}
}

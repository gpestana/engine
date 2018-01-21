package source

import (
	"github.com/gpestana/engine/data"
)

type Rss struct {
	Type   string
	Url    string
	signal chan struct{}
}

func (r Rss) Fetch() []data.DataUnit {
	results := []data.DataUnit{}
	return results
}

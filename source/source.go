package source

import (
	"github.com/gpestana/engine/data"
)

type Source interface {
	Fetch() []data.DataUnit
}

func New(t string, u string) Source {
	switch t {
	case "rss":
		return Rss{
			Type: t,
			Url:  u,
		}
	default:
		panic("Source type is invalid: " + t)
	}
}

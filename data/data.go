package data

import (
	"encoding/json"
	"time"
)

type DataUnit struct {
	Url         string
	Timestamp   time.Time
	Description string
	// Pointer from where DataUnit was fetched
	Parent interface{}
	// Original content fetched from Parent
	FetchedContent interface{}
	Payload        []byte
}

func (d DataUnit) Handle() DataUnit {
	return d
}

func (d DataUnit) String() string {
	s, _ := json.Marshal(d)
	return string(s)
}

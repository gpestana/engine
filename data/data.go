package data

import (
	"time"
)

type DataUnit struct {
	Url       string
	Timestamp time.Time
	Payload   []byte
}

func (d DataUnit) Handle() DataUnit {
	return d
}

func (d DataUnit) String() string {
	return d.Url
}

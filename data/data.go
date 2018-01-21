package data

import (
	"time"
)

type DataUnit struct {
	Url       string
	Timestamp time.Time
	Payload   []byte
}

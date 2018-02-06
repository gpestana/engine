package data

import (
	"encoding/json"
	"github.com/gpestana/htmlizer"
	"io/ioutil"
	"net/http"
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
	Errors         []error
}

func (d *DataUnit) Handle() DataUnit {
	err := d.getContent()
	if err != nil {
		d.addError(err)
		// breaks data unit handling
		return *d
	}

	return *d
}

func (d *DataUnit) getContent() error {
	r, err := http.Get(d.Url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	h := htmlizer.New()
	h.Load(string(body))
	d.Payload = []byte(h.HumanReadable())
	return nil
}

func (d *DataUnit) addError(err error) {
	d.Errors = append(d.Errors, err)
}

func (d DataUnit) String() string {
	type DataUnitPrint struct {
		Url            string
		Timestamp      time.Time
		Description    string
		FetchedContent interface{}
		Payload        string
		Errors         []error
	}

	unitPrint := DataUnitPrint{
		Url:            d.Url,
		Timestamp:      d.Timestamp,
		Description:    d.Description,
		FetchedContent: d.FetchedContent,
		Payload:        string(d.Payload),
		Errors:         d.Errors,
	}

	s, err := json.Marshal(unitPrint)
	if err != nil {
		return "DataUnit: Error printing -  " + err.Error()
	}
	return string(s)
}

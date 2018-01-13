package main

import (
	"log"
	"os"
	"testing"
)

func TestEmptyConfig(t *testing.T) {
	p := "./test_c.json"
	c := `{
	"sources": [],
	"filters": [],
	"outputs": []
}
`
	defer os.Remove(p)
	createConfigFile(p, c)
	conf, err := config(p)
	if err != nil {
		t.Error(err)
	}

	if len(conf.Sources) != 0 {
		t.Error("empty conf: sources should have been empty but it has number of entries " + string(len(conf.Sources)))
	}

	if len(conf.Filters) != 0 {
		t.Error("empty conf: filters should have been empty but it has number of entries " + string(len(conf.Filters)))
	}

	if len(conf.Outputs) != 0 {
		t.Error("empty conf: outputs should have been empty but it has number of entries " + string(len(conf.Outputs)))
	}

}

func TestConfig(t *testing.T) {
	p := "./test_c.json"
	c := `{
	"sources": [{
		"type": "rss",
		"url": "url_rss"
	},
	{
		"type": "twitter",
		"url": "url_twitter"
	}
	],
	"filters": [{
		"type": "news",
		"filter": "a filter defined by user"
	}],
	"outputs": []
}
`
	defer os.Remove(p)
	createConfigFile(p, c)
	conf, err := config(p)
	if err != nil {
		t.Error(err)
	}

	if len(conf.Sources) != 2 {
		t.Error("conf: number of sources should be 2, instead it has number of entries " + string(len(conf.Sources)))
	}

	if len(conf.Filters) != 1 {
		t.Error("conf: filters should be 2, instead we got " + string(len(conf.Filters)))
		return
	}

	url := conf.Sources[1].Url
	if url != "url_twitter" {
		t.Error("conf: url for 2nd filter conf should be 'url_twitter', instead it is " + url)
	}

	filter := conf.Filters[0].Filter
	if filter != "a filter defined by user" {
		t.Error("conf: filter should be 'a filter defined by user', instead " + filter)
	}

}

func createConfigFile(p string, conf string) {
	f, err := os.Create(p)
	if err != nil {
		log.Fatal("test setup failed" + err.Error())
	}
	_, err = f.Write([]byte(conf))
}

package main

import (
	"encoding/json"
	"flag"
	// "github.com/gpestana/engine/data"
	"github.com/gpestana/engine/source"
	"io/ioutil"
	"log"
)

type Config struct {
	Sources []ConfigEntry
	Filters []ConfigEntry
	Outputs []ConfigEntry
}

type ConfigEntry struct {
	Url    string `json:"url,omitempty"`
	Type   string `json:"type,omitempty"`
	Filter string `json:"filter,omitempty"`
}

func main() {
	c := flag.String("config", "", "path for configuration file")
	flag.Parse()

	if *c == "" {
		log.Fatal("Error: Provide a configuration file path (-conf=/path)")
	}

	conf, err := config(*c)
	if err != nil {
		log.Fatal("Configuration parsing: " + err.Error())
	}

	dataCount := 0
	dataCh := make(chan data.DataUnit)

	sources := []source.Source{}
	for _, src := range conf.Sources {
		s := source.New(src.Type, src.Url)
		sources = append(sources, s)

		go func(s source.Source) {
			data := s.Fetch()
			//mutex!
			dataCount += len(data)
			dataCh <- data
		}(s)
	}

}

func config(p string) (Config, error) {
	c, err := ioutil.ReadFile(p)
	if err != nil {
		return Config{}, err
	}
	conf := Config{}
	err = json.Unmarshal([]byte(c), &conf)
	if err != nil {
		return Config{}, err
	}
	return conf, nil
}

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gpestana/engine/data"
	"github.com/gpestana/engine/source"
	"io/ioutil"
	"log"
	"sync"
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

	sources := []source.Source{}
	srcCh := make(chan []data.DataUnit, len(sources))
	unitsCh := make(chan data.DataUnit)
	var wg sync.WaitGroup

	for _, src := range conf.Sources {
		s := source.New(src.Type, src.Url)
		sources = append(sources, s)

		go func() {
			srcCh <- s.Fetch()
		}()
	}

	for i := 0; i < len(sources); i++ {
		units := []data.DataUnit{}
		units = <-srcCh
		wg.Add(len(units))
		for _, u := range units {
			go func(u data.DataUnit) {
				unitsCh <- u.Handle()
				wg.Done()
			}(u)
		}
	}

	// data channel closer
	go func() {
		wg.Wait()
		close(unitsCh)
	}()

	// receive units
	for u := range unitsCh {
		fmt.Println(u)
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

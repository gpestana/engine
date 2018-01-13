package main

import (
	"encoding/json"
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
	// parse config
	// get from source (RSS, etc..)
	// -> Content
	// add to storage
	// query storoge
	// add to hit storage
	// output
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

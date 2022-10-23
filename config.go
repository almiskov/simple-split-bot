package main

import (
	"encoding/json"
	"io/ioutil"
)

type config struct {
	Token string `json:"token"`
	Debug bool   `json:"debug"`
}

func (cfg *config) Parse(file string) error {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, cfg)
}

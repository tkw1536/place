package config

import (
	"encoding/json"
	"io/ioutil"
)

// Load loads the configuration from a file
func (cfg *Config) Load(path string) error {
	file, e := ioutil.ReadFile(path)
	if e != nil {
		return e
	}
	return json.Unmarshal(file, cfg)
}

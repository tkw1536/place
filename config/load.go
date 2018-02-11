package config

import (
	"encoding/json"
	"io"
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

// Save Configuration to a file
func (cfg *Config) Save(path string) error {
	s, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(path, s, 0644)
}

// Read the configuration from anything that can read
func (cfg *Config) Read(r io.Reader) error {
	dec := json.NewDecoder(r)
	return dec.Decode(cfg)
}

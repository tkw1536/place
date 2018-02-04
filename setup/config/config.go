package config

import (
	"encoding/json"
	"github.com/tkw1536/place/config"
	"os"
)

const (
	// DefaultPath indicates the default location the config is loaded from and saved to
	DefaultPath = "config.json"
)

var theConfig config.Config

// ReadFromPath loads the config file from specified path
func ReadFromPath(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	dec := json.NewDecoder(file)
	err = dec.Decode(&theConfig)
	return err
}

// WriteToPath saves the config file to the specified path
func WriteToPath(path string) error {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}

	enc := json.NewEncoder(file)
	err = enc.Encode(&theConfig)
	return err
}

// Get returns the currently loaded configuration
func Get() *config.Config {
	return &theConfig
}

func Set(conf *config.Config) {
	theConfig = *conf
}

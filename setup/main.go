package main

import (
	"os"

	"github.com/tkw1536/place/config"
	"github.com/tkw1536/place/utils"
)

/* minimal configuration */
var (
	listenAddr   = "0.0.0.0:80"  // address to server under
	configPath   = "config.json" // path to read / write configuration from
	childCommand = os.Args[1:]   // command to start
)

func init() {
	if envPath := os.Getenv("CONFIG_PATH"); envPath != "" {
		configPath = envPath
	}

	if envAddr := os.Getenv("BIND_ADDRESS"); envAddr != "" {
		listenAddr = envAddr
	}
}

// isConfigured checks if the configuration file exists and is valid
func isConfigured() bool {
	var cfg config.Config

	err := cfg.Load(configPath)
	if err != nil {
		utils.Logger.Printf("Failed to load config file from %v: %v", configPath, err)
	}

	return err == nil
}

// The Main program
func main() {
	for true {
		// if the configuration file is valid, start the program and exit with the right code
		if isConfigured() {
			utils.Logger.Printf("Configuration file at %s valid, starting program", configPath)
			os.Exit(runChild())

			// else run the setup
		} else {
			utils.Logger.Printf("Configuration file at %s invalid, starting setup", configPath)
			setupServer()
			utils.Logger.Printf("Setup terminated, re-starting")
		}
	}
}

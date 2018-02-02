// implements an entry point for the place docker container
package main

import "../../place"
import "../../place/config"

var cfg config.Config

func main() {
	// read the configuration
	cfg.ParseOrPanic()
	cfg.Inspect()

	// and we are going to place stuff now
	place.StartPlace(&cfg)
}

// place-git-update utility
// update static deployment using git
package main

import (
	"os"

	"../../updater"
	"../../updater/config"
)

var cfg config.Config

func main() {
	// read the configuration
	config.ParseConfigOrPanic(&cfg, os.Args[1:])
	config.InspectConfig(&cfg)

	// and start the server
	if err := updater.RunUpdate(&cfg); err != nil {
		panic(err)
	}
}

// implements an entry point for the place docker container
package main

import (
	"fmt"
	"os"

	"github.com/tkw1536/place/config"
	"github.com/tkw1536/place/server"
)

var cfg config.Config

func main() {
	if len(os.Args) != 2 {
		panic(fmt.Errorf("Usage: %s configfile", os.Args[0]))
	}
	// read the configuration
	if err := cfg.Load(os.Args[1]); err != nil {
		panic(fmt.Errorf("Can't load configfile %s: %s", os.Args[1], err.Error()))
	}
	cfg.Inspect()

	// and we are going to place stuff now
	server.StartServer(&cfg)
}

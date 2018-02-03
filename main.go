// implements an entry point for the place docker container
package main

import (
	"os"

	"github.com/tkw1536/place/config"
	"github.com/tkw1536/place/server"
)

var cfg config.Config

func main() {
	// read the configuration
	cfg.Load(os.Args[1])
	cfg.Inspect()

	// and we are going to place stuff now
	server.StartServer(&cfg)
}

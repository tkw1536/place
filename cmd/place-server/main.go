// place-server utility
// runs a webserver that serves static directories along with listening for webhooks
package main

import (
	"os"

	"../../server"
	"../../server/config"
)

var cfg config.Config

func main() {
	// read the configuration
	cfg.ParseOrPanic(os.Args[1:])
	cfg.Inspect()

	// and start the server
	server.StartServer(&cfg)
}
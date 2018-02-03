package config

import (
	"bytes"
	"net/url"
	"time"

	"../../utils"
	"../checkers"
)

// All configuration settings used by Place
type Config struct {
	BindAddress string
	HookPath    string

	Checkers []checkers.Checker

	ScriptCommand []string
	ScriptTimeout time.Duration

	StaticPath string
	ProxyURL   *url.URL
}

// Inspect prints some configuration information to the logger
func (cfg Config) Inspect() {
	utils.Logger.Printf("BindAddress: %s\n", cfg.BindAddress)
	utils.Logger.Printf("HookPath:    %s\n", cfg.HookPath)

	buffer := bytes.NewBufferString("Checkers:  ")
	for _, checker := range cfg.Checkers {
		buffer.WriteString(" ")
		buffer.WriteString(checker.String())
	}
	utils.Logger.Print(buffer.String())

	utils.Logger.Printf("ScriptCommand: %s\n", cfg.ScriptCommand)
	utils.Logger.Printf("ScriptTimeout: %d\n", cfg.ScriptTimeout/time.Second)
}

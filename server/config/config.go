package config

import (
	"bytes"
	"log"
	"net/url"
	"time"

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

	Logger *log.Logger
}

// Inspect prints some configuration information to the logger
func (cfg *Config) Inspect() {
	cfg.Logger.Printf("BindAddress: %s\n", cfg.BindAddress)
	cfg.Logger.Printf("HookPath:    %s\n", cfg.HookPath)

	buffer := bytes.NewBufferString("Checkers:  ")
	for _, checker := range cfg.Checkers {
		buffer.WriteString(" ")
		buffer.WriteString(checker.String())
	}
	cfg.Logger.Print(buffer.String())

	cfg.Logger.Printf("ScriptCommand: %s\n", cfg.ScriptCommand)
	cfg.Logger.Printf("ScriptTimeout: %d\n", cfg.ScriptTimeout/time.Second)

	if cfg.StaticPath != "" {

	}

}

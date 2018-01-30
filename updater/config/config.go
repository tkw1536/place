package config

import (
	"log"
)

// Config represents the Configuration used by the updater script
type Config struct {
	RepositoryURL string
	SSHKeyPath    string
	OutDirectory  string
	Ref           string
	BuildScript   string

	Logger *log.Logger
}

// InspectConfig prints information about the configuration
func InspectConfig(cfg *Config) {
	cfg.Logger.Printf("RepositoryURL: %s\n", cfg.RepositoryURL)
	cfg.Logger.Printf("SSHKeyPath:    %s\n", cfg.SSHKeyPath)
	cfg.Logger.Printf("OutDirectory:  %s\n", cfg.OutDirectory)
	cfg.Logger.Printf("Ref:           %s\n", cfg.Ref)
	cfg.Logger.Printf("BuildScript:   %s\n", cfg.BuildScript)
}

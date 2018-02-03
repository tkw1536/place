package config

import (
	"../../utils"
	"gopkg.in/src-d/go-git.v4/plumbing/transport"
)

// Config represents the Configuration used by the updater script
type Config struct {
	RepositoryURL *transport.Endpoint
	SSHKeyPath    string
	OutDirectory  string
	Ref           string
	BuildScript   string
}

// Inspect prints information about the configuration
func (cfg Config) Inspect() {
	utils.Logger.Printf("RepositoryURL: %s\n", cfg.RepositoryURL.String())
	utils.Logger.Printf("SSHKeyPath:    %s\n", cfg.SSHKeyPath)
	utils.Logger.Printf("OutDirectory:  %s\n", cfg.OutDirectory)
	utils.Logger.Printf("Ref:           %s\n", cfg.Ref)
	utils.Logger.Printf("BuildScript:   %s\n", cfg.BuildScript)
}

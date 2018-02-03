package config

import (
	"net/url"
	"time"

	"../../utils"

	"gopkg.in/src-d/go-git.v4/plumbing/transport"
)

// Config specifies the configuration used by place
type Config struct {
	SSHKeyPath      string
	BindAddress     string
	WebhookPath     string
	GitURL          *transport.Endpoint
	GitBranch       string
	GitCloneTimeout time.Duration
	GitHubSecret    string
	GitLabSecret    string
	Debug           bool
	StaticPath      string
	BuildScript     string
	ProxyURL        *url.URL
}

// Inspect inspects the configuration
func (cfg Config) Inspect() {
	utils.Logger.Printf("SSHKeyPath:      %s\n", cfg.SSHKeyPath)
	utils.Logger.Printf("BindAddress:     %s\n", cfg.BindAddress)
	utils.Logger.Printf("WebhookPath:     %s\n", cfg.WebhookPath)
	utils.Logger.Printf("GitBranch:       %s\n", cfg.GitBranch)
	utils.Logger.Printf("GitURL:          %s\n", cfg.GitURL)
	utils.Logger.Printf("GitCloneTimeout: %d\n", cfg.GitCloneTimeout/1000)
	utils.Logger.Printf("GitHubSecret:    %s\n", cfg.GitHubSecret)
	utils.Logger.Printf("GitLabSecret:    %s\n", cfg.GitLabSecret)
	utils.Logger.Printf("Debug:           %t\n", cfg.Debug)
	utils.Logger.Printf("StaticPath:      %s\n", cfg.StaticPath)
	utils.Logger.Printf("BuildScript:     %s\n", cfg.BuildScript)
	utils.Logger.Printf("ProxyURL:        %s\n", cfg.ProxyURL)
}

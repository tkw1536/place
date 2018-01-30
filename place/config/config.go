package config

import "log"

// Config specifies the configuration used by place
type Config struct {
	SSHKeyPath      string
	BindAddress     string
	WebhookPath     string
	GitURL          string
	GitBranch       string
	GitCloneTimeout string
	GitHubSecret    string
	GitLabSecret    string
	Debug           bool
	StaticPath      string
	BuildScript     string
	ProxyURL        string

	Logger *log.Logger
}

// InspectConfig prints information about the configuration
func InspectConfig(cfg *Config) {
	cfg.Logger.Printf("SSHKeyPath:      %s\n", cfg.SSHKeyPath)
	cfg.Logger.Printf("BindAddress:     %s\n", cfg.BindAddress)
	cfg.Logger.Printf("WebhookPath:     %s\n", cfg.WebhookPath)
	cfg.Logger.Printf("GitBranch:       %s\n", cfg.GitBranch)
	cfg.Logger.Printf("GitURL:          %s\n", cfg.GitURL)
	cfg.Logger.Printf("GitCloneTimeout: %s\n", cfg.GitCloneTimeout)
	cfg.Logger.Printf("GitHubSecret:    %s\n", cfg.GitHubSecret)
	cfg.Logger.Printf("GitLabSecret:    %s\n", cfg.GitLabSecret)
	cfg.Logger.Printf("Debug:           %t\n", cfg.Debug)
	cfg.Logger.Printf("StaticPath:      %s\n", cfg.StaticPath)
	cfg.Logger.Printf("BuildScript:     %s\n", cfg.BuildScript)
	cfg.Logger.Printf("ProxyURL:        %s\n", cfg.ProxyURL)
}

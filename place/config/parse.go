package config

import (
	"log"
	"os"
)

// ParseConfig parses configuration from command line arguments
func ParseConfig(cfg *Config) error {
	cfg.Logger = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)

	cfg.SSHKeyPath = os.Getenv("SSH_KEY_PATH")
	cfg.BindAddress = os.Getenv("BIND_ADDRESS")
	cfg.WebhookPath = os.Getenv("WEBHOOK_PATH")
	cfg.GitURL = os.Getenv("GIT_URL")
	cfg.GitBranch = os.Getenv("GIT_BRANCH")
	cfg.GitCloneTimeout = os.Getenv("GIT_CLONE_TIMEOUT")
	cfg.GitHubSecret = os.Getenv("GITHUB_SECRET")
	cfg.GitLabSecret = os.Getenv("GITLAB_SECRET")
	cfg.Debug = os.Getenv("DEBUG") == "1"
	cfg.StaticPath = os.Getenv("STATIC_PATH")
	cfg.BuildScript = os.Getenv("BUILD_SCRIPT")
	cfg.ProxyURL = os.Getenv("PROXY_URL")

	return nil
}

// ParseConfigOrPanic parses and returns configuration or panics
func ParseConfigOrPanic(cfg *Config) {
	err := ParseConfig(cfg)
	if err != nil {
		panic(err.Error())
	}
}

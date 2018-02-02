package config

import (
	"log"
	"net/url"
	"os"
	"strconv"
	"time"
)

// Parse parses the configuration from the environment
func (cfg Config) Parse() error {
	cfg.Logger = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)

	cfg.SSHKeyPath = os.Getenv("SSH_KEY_PATH")
	cfg.BindAddress = os.Getenv("BIND_ADDRESS")
	cfg.WebhookPath = os.Getenv("WEBHOOK_PATH")
	cfg.GitURL = os.Getenv("GIT_URL")
	cfg.GitBranch = os.Getenv("GIT_BRANCH")
	timeout, err := strconv.Atoi(os.Getenv("GIT_CLONE_TIMEOUT"))
	if err != nil || timeout <= 0 {
		if err != nil {
			return err
		}
		cfg.GitCloneTimeout = time.Duration(600) * time.Second
	} else {
		cfg.GitCloneTimeout = time.Duration(timeout) * time.Second
	}
	cfg.GitHubSecret = os.Getenv("GITHUB_SECRET")
	cfg.GitLabSecret = os.Getenv("GITLAB_SECRET")
	cfg.Debug = os.Getenv("DEBUG") == "1"
	cfg.StaticPath = os.Getenv("STATIC_PATH")
	cfg.BuildScript = os.Getenv("BUILD_SCRIPT")

	pus := os.Getenv("PROXY_URL")
	if pus != "" {
		pu, err := url.Parse(pus)
		if err != nil {
			return err
		}
		cfg.ProxyURL = pu
	}

	return nil
}

// ParseOrPanic parses the configuration or panics
func (cfg Config) ParseOrPanic() {
	if err := cfg.Parse(); err != nil {
		panic(err.Error())
	}
}

package config

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"time"

	"gopkg.in/src-d/go-git.v4/plumbing/transport"
)

func getEnvWithDefault(env, def string) string {
	if val := os.Getenv(env); val != "" {
		return val
	}
	return def
}

// Parse parses the configuration from the environment
func (cfg *Config) Parse() error {
	// Bind Address and webhook path
	cfg.BindAddress = getEnvWithDefault("BIND_ADDRESS", "0.0.0.0:80")
	cfg.WebhookPath = getEnvWithDefault("WEBHOOK_PATH", "/webhook/")

	// git url and optional ssh key path
	g := os.Getenv("GIT_URL")
	gu, err := transport.NewEndpoint(g)
	if err != nil {
		return fmt.Errorf("%s is not a valid GIT_URL", g)
	}
	cfg.GitURL = gu
	cfg.SSHKeyPath = os.Getenv("SSH_KEY_PATH")

	// git branch -- defaults to master
	cfg.GitBranch = getEnvWithDefault("GIT_BRANCH", "master")

	// timeout defaults to 600
	timeout, err := strconv.Atoi(getEnvWithDefault("GIT_CLONE_TIMEOUT", "600"))
	if err != nil {
		return err
	} else if timeout <= 0 {
		return fmt.Errorf("Timeout must be > 0, not %d", timeout)
	} else {
		cfg.GitCloneTimeout = time.Duration(timeout) * time.Second
	}

	// github and gitlab secrets (if any)
	// TODO: Use checker environments for this and allow auto-generation of secrets
	cfg.GitHubSecret = os.Getenv("GITHUB_SECRET")
	cfg.GitLabSecret = os.Getenv("GITLAB_SECRET")
	cfg.Debug = os.Getenv("DEBUG") == "1"

	cfg.StaticPath = getEnvWithDefault("STATIC_PATH", "/var/www/html")

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
func (cfg *Config) ParseOrPanic() {
	if err := cfg.Parse(); err != nil {
		panic(err.Error())
	}
}

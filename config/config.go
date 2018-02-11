package config

import (
	"time"

	"github.com/tkw1536/place/utils"
)

// Config specifies the configuration used by place
type Config struct {
	BindAddress string // address to bind to, defaults to 0.0.0.0:80
	WebhookPath string // path to server webhook under, defaults to /webhook/

	GitURL    *MarshalableEndpoint // url to clone repository from
	GitBranch string               // branch to clone and set of events to listen to, defaults to "master"

	GitSSHKeyPath string // ssh key for git clone (if any)
	GitUsername   string // username for git clone (if any)
	GitPassword   string // password for git clone (if any)

	GitCloneTimeout time.Duration // timeout for git clone

	GitHubSecret string // secret to use for GitHub Checker, or nil to disable
	GitLabSecret string // secret to use for GitLab Checker, or nil to disable
	Debug        bool   // enable the Debug checker, if any

	StaticPath string          // if set, serve static directories from this path
	ProxyURL   *MarshalableURL // if set, proxy urls to this path

	BuildScript string // use this script to build the website dynamically
}

// Inspect inspects the configuration
func (cfg Config) Inspect() {
	utils.Logger.Printf("BindAddress:     %s\n", cfg.BindAddress)
	utils.Logger.Printf("WebhookPath:     %s\n", cfg.WebhookPath)

	utils.Logger.Printf("GitURL:          %s\n", cfg.GitURL.Endpoint().String())
	utils.Logger.Printf("GitBranch:       %s\n", cfg.GitBranch)

	utils.Logger.Printf("SSHKeyPath:      %s\n", cfg.GitSSHKeyPath)
	utils.Logger.Printf("GitUsername:     %s\n", cfg.GitUsername)

	if cfg.GitPassword != "" {
		utils.Logger.Printf("GitPassword:     ******\n")
	} else {
		utils.Logger.Printf("GitPassword:     \n")
	}

	utils.Logger.Printf("GitCloneTimeout: %d\n", cfg.GitCloneTimeout/1000)

	utils.Logger.Printf("GitHubSecret:    %s\n", cfg.GitHubSecret)
	utils.Logger.Printf("GitLabSecret:    %s\n", cfg.GitLabSecret)
	utils.Logger.Printf("Debug:           %t\n", cfg.Debug)

	utils.Logger.Printf("StaticPath:      %s\n", cfg.StaticPath)
	utils.Logger.Printf("BuildScript:     %s\n", cfg.BuildScript)
	utils.Logger.Printf("ProxyURL:        %s\n", cfg.ProxyURL.URL().String())
}

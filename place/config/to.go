package config

// ToServerConfig turns an entry configuration into a server configuration
import (
	"../../server/checkers"
	serverconfig "../../server/config"
	updaterconfig "../../updater/config"
)

// ToServerConfig turns this configuration into a server configuration
func (cfg Config) ToServerConfig() serverconfig.Config {
	var c serverconfig.Config
	c.Logger = cfg.Logger
	c.BindAddress = cfg.BindAddress
	c.HookPath = cfg.WebhookPath

	var chks []checkers.Checker
	if cfg.GitHubSecret != "" {
		// load all the checkers
		if cfg.GitHubSecret != "" {
			chks = append(chks, checkers.CreateChecker("github", cfg.GitHubSecret))
		}

		if cfg.GitLabSecret != "" {
			chks = append(chks, checkers.CreateChecker("gitlab", cfg.GitLabSecret))
		}

		if cfg.Debug {
			chks = append(chks, checkers.CreateChecker("debug", ""))
		}
	}
	c.Checkers = chks

	c.ScriptCommand = cfg.ToUpdaterConfig().ToCommand("/place/bin/place-git-update")
	c.ScriptTimeout = cfg.GitCloneTimeout
	c.StaticPath = cfg.StaticPath
	c.ProxyURL = cfg.ProxyURL
	return c
}

// ToUpdaterConfig turns this configuration into a updater configuration
func (cfg Config) ToUpdaterConfig() updaterconfig.Config {
	var c updaterconfig.Config
	c.Logger = cfg.Logger
	c.RepositoryURL = cfg.GitURL
	c.SSHKeyPath = cfg.SSHKeyPath
	c.OutDirectory = cfg.StaticPath
	c.Ref = "refs/heads/" + cfg.GitBranch
	c.BuildScript = cfg.BuildScript

	return c
}

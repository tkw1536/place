package config

// ToServerConfig turns an entry configuration into a server configuration
import (
	"os"
	"path/filepath"

	"../../server/checkers"
	serverconfig "../../server/config"
	updaterconfig "../../updater/config"
)

// ToServerConfig turns this configuration into a server configuration
func (cfg Config) ToServerConfig() serverconfig.Config {
	var c serverconfig.Config
	c.BindAddress = cfg.BindAddress
	c.HookPath = cfg.WebhookPath

	// load all the checkers
	chks := make([]checkers.Checker, 0)
	if cfg.GitHubSecret != "" {
		chks = append(chks, checkers.CreateChecker("github", cfg.GitHubSecret+",refs/heads/"+cfg.GitBranch))
	}

	if cfg.GitLabSecret != "" {
		chks = append(chks, checkers.CreateChecker("gitlab", cfg.GitLabSecret+",refs/heads/"+cfg.GitBranch))
	}

	if cfg.Debug {
		chks = append(chks, checkers.CreateChecker("debug", ""))
	}
	c.Checkers = chks

	var placePath string
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err == nil {
		placePath = filepath.Join(dir, "place-git-update")
	} else {
		placePath = "/place/bin/place-git-update"
	}
	c.ScriptCommand = cfg.ToUpdaterConfig().ToCommand(placePath)
	c.ScriptTimeout = cfg.GitCloneTimeout
	c.StaticPath = cfg.StaticPath
	c.ProxyURL = cfg.ProxyURL
	return c
}

// ToUpdaterConfig turns this configuration into a updater configuration
func (cfg Config) ToUpdaterConfig() updaterconfig.Config {
	var c updaterconfig.Config
	c.RepositoryURL = cfg.GitURL
	c.SSHKeyPath = cfg.SSHKeyPath
	c.OutDirectory = cfg.StaticPath
	c.Ref = "refs/heads/" + cfg.GitBranch
	c.BuildScript = cfg.BuildScript

	return c
}

package config

import (
	"github.com/tkw1536/place/server/checkers"
	"github.com/tkw1536/place/utils/sshkey"
	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
)

// Checkers returns the set of checkers induced by this configuration
func (cfg Config) Checkers() []checkers.Checker {
	var chks = make([]checkers.Checker, 0)

	if cfg.GitHubSecret != "" {
		var checker checkers.GitHubChecker
		checker.Create(cfg.GitHubSecret, cfg.GitBranch)
		chks = append(chks, &checker)
	}

	if cfg.GitLabSecret != "" {
		var checker checkers.GitLabChecker
		checker.Create(cfg.GitLabSecret, cfg.GitBranch)
		chks = append(chks, &checker)
	}

	if cfg.Debug {
		var checker checkers.DebugChecker
		chks = append(chks, &checker)
	}

	return chks
}

// GitRef returns the ref pointed to by this Configuration
func (cfg Config) GitRef() string {
	return "refs/head/" + cfg.GitBranch
}

// GitCloneOptions generates options for use with git clone
func (cfg Config) GitCloneOptions() (*git.CloneOptions, error) {
	var options git.CloneOptions

	options.URL = cfg.GitURL.Endpoint().String()
	options.ReferenceName = plumbing.ReferenceName(cfg.GitRef())
	options.SingleBranch = true
	options.Depth = 1

	if cfg.GitSSHKeyPath != "" {
		auth, err := sshkey.Get(cfg.GitSSHKeyPath)
		if err != nil {
			return nil, err
		}
		options.Auth = auth
	}

	if cfg.GitUsername != "" && cfg.GitPassword != "" {
		options.Auth = &http.BasicAuth{Username: cfg.GitUsername, Password: cfg.GitPassword}
	}

	return &options, nil
}

package config

import "github.com/tkw1536/place/server/checkers"

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

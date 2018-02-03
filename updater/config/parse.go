package config

import (
	"flag"

	"gopkg.in/src-d/go-git.v4/plumbing/transport"
)

// ParseConfig parses configuration from command line arguments
func ParseConfig(cfg *Config, args []string) error {
	flagset := flag.NewFlagSet("updater", flag.ContinueOnError)

	var repoURL string
	flagset.StringVar(&repoURL, "from", "", "URL the remote repository is located at")
	flagset.StringVar(&cfg.Ref, "ref", "refs/heads/master", "Ref the repository should be checked out at")
	flagset.StringVar(&cfg.SSHKeyPath, "ssh-key", "", "If set, load a passwordless ssh key from the given path and use it to clone the repository")
	flagset.StringVar(&cfg.OutDirectory, "to", "", "Local path the static files should be placed at")
	flagset.StringVar(&cfg.BuildScript, "build", "", "Build script to replace default checkout process")

	flagset.Parse(args)

	tep, err := transport.NewEndpoint(repoURL)
	if err != nil {
		return err
	}
	cfg.RepositoryURL = tep

	return validateConfig(cfg)
}

// ParseConfigOrPanic is a utility function that parses the configuration or panics
func ParseConfigOrPanic(cfg *Config, args []string) {
	err := ParseConfig(cfg, args)
	if err != nil {
		panic(err.Error())
	}
}

package config

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"time"

	"../../utils/command"
	"../checkers"
)

// ParseConfig parses configuration from the command line
func ParseConfig(cfg *Config, args []string) error {
	cfg.Logger = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)

	flag.StringVar(&cfg.BindAddress, "bind", "", "address to bind to")
	flag.StringVar(&cfg.HookPath, "webhook", "", "path that should respond to webhooks")

	var gitHubChecker string
	flag.StringVar(&gitHubChecker, "github", "", "Run the webhook whenever a GitHub web hook request is received")

	var gitLabChecker string
	flag.StringVar(&gitLabChecker, "gitlab", "", "Run the webhook whenever a GitLab web hook request is received.")

	var debugChecker bool
	flag.BoolVar(&debugChecker, "debug", false, "Run in debug mode and run the webhook whenever any POST request is received.")

	var scriptLine string
	flag.StringVar(&scriptLine, "script", "", "script command to run")

	var timeout int
	flag.IntVar(&timeout, "timeout", 600, "timeout for hook script in seconds")

	flag.StringVar(&cfg.StaticPath, "static", "", "file system path to manage static files in")

	var proxyURL string
	flag.StringVar(&proxyURL, "proxy", "", "Instead of serving static content, proxy all requests to url")

	flag.CommandLine.Parse(args)

	// split the hookLine into a set of arguments
	var err error
	cfg.ScriptCommand, err = command.Split(scriptLine)
	if err != nil {
		return fmt.Errorf("unable to parse --script: %s", err.Error())
	}

	// load all the checkers
	if gitHubChecker != "" {
		cfg.Checkers = append(cfg.Checkers, checkers.CreateChecker("github", gitHubChecker))
	}

	if gitLabChecker != "" {
		cfg.Checkers = append(cfg.Checkers, checkers.CreateChecker("gitlab", gitLabChecker))
	}

	if debugChecker {
		cfg.Checkers = append(cfg.Checkers, checkers.CreateChecker("debug", ""))
	}

	// turn timeout into a duration
	cfg.ScriptTimeout = time.Duration(timeout) * time.Second

	// configure a list of checkers
	cfg.Checkers = []checkers.Checker{}

	// parse proxyURL
	if proxyURL != "" {
		u, err := url.Parse(proxyURL)
		if err != nil {
			return fmt.Errorf("unable to parse --proxy: %s", err.Error())
		}
		cfg.ProxyURL = u
	}

	return verifyConfig(cfg)
}

// ParseConfigOrPanic parses and returns configuration or panics
func ParseConfigOrPanic(cfg *Config, args []string) {
	err := ParseConfig(cfg, args)
	if err != nil {
		panic(err.Error())
	}
}

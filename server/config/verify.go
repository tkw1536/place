package config

import (
	"fmt"
	"regexp"

	"../../utils/verify"
	"../checkers"
)

// verifies the configuration settings
func verifyConfig(cfg *Config) error {
	if err := verifyAddress(cfg.BindAddress); err != nil {
		return err
	}

	if err := verifyHookPath(cfg.HookPath); err != nil {
		return err
	}

	if err := verifyCheckers(cfg.Checkers); err != nil {
		return err
	}

	if err := verify.Command(cfg.ScriptCommand, false); err != nil {
		return err
	}

	if cfg.ScriptTimeout <= 0 {
		return fmt.Errorf("ScriptTimout must be postivive, not %d", cfg.ScriptTimeout)
	}

	if cfg.StaticPath == "" && cfg.ProxyURL == nil {
		return fmt.Errorf("either --static or --proxy must be given")
	} else if cfg.StaticPath != "" && cfg.ProxyURL != nil {
		return fmt.Errorf("at most one of --static or --proxy may be given, not both")
	}

	if cfg.StaticPath != "" {
		if err := verify.Dir(cfg.StaticPath); err != nil {
			return err
		}
	}

	return nil
}

var addressRegex = regexp.MustCompile("^(([^:]+):)(\\d+)$")

// verify a bind Address
func verifyAddress(address string) error {
	if addressRegex.MatchString(address) {
		return nil
	}

	return fmt.Errorf("not a valid BindAddress: %q", address)
}

var hookRegex = regexp.MustCompile("^/[^\\s\\?]+/$")

// verify a hook path
func verifyHookPath(path string) error {
	if hookRegex.MatchString(path) {
		return nil
	}

	return fmt.Errorf("not a valid HookPath: %q, should start and end with %q", path, "/")
}

func verifyCheckers(checkers []checkers.Checker) error {
	if len(checkers) == 0 {
		return fmt.Errorf("at least one webhook listener must be enabled")
	}

	return nil
}

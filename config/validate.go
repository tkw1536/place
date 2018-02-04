package config

import (
	"fmt"
	"regexp"

	"github.com/tkw1536/place/utils/verify"
)

// Validate checks if the configuration is valid
// and returns nil if it is, and an error message if it is not.
func (cfg Config) Validate() error {
	if cfg.SSHKeyPath != "" {
		if err := verify.File(cfg.SSHKeyPath); err != nil {
			return err
		}
	}

	if err := validateAddress(cfg.BindAddress); err != nil {
		return err
	}

	if err := verifyHookPath(cfg.WebhookPath); err != nil {
		return err
	}

	if cfg.StaticPath != "" && cfg.ProxyURL != nil {
		return fmt.Errorf("Both StaticPath and ProxyURL are set, you must one exactly one. ")
	}

	if cfg.StaticPath == "" && cfg.ProxyURL == nil {
		return fmt.Errorf("Neither StaticPath and ProxyURL are set, you must one exactly one. ")
	}

	if cfg.StaticPath != "" {
		if err := verify.Dir(cfg.StaticPath); err != nil {
			return err
		}
	}

	// TODO: Do the validation
	return nil
}

var addressRegex = regexp.MustCompile("^(([^:]+):)(\\d+)$")

func validateAddress(address string) error {
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

	return fmt.Errorf("not a valid WebhookPath: %q, should start and end with %q", path, "/")
}

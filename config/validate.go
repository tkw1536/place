package config

import (
	"fmt"
	"regexp"

	"github.com/tkw1536/place/utils/verify"
)

// Validate checks if the configuration is valid
// and returns nil if it is, and an error message if it is not.
func (cfg Config) Validate() error {
	if cfg.GitUsername != "" && cfg.GitSSHKeyPath != "" {
		return fmt.Errorf("Both GitUsername and GitSSHKeyPath are set, you may use at most one. ")
	}

	if cfg.GitUsername != "" && cfg.GitPassword != "" {
		return fmt.Errorf("GitUsername is set, but no password was given. ")
	}

	if cfg.GitSSHKeyPath != "" {
		if err := verify.File(cfg.GitSSHKeyPath); err != nil {
			return err
		}
	}

	if err := validateAddress(cfg.BindAddress); err != nil {
		return err
	}

	if err := verifyHookPath(cfg.WebhookPath); err != nil {
		return err
	}

	if cfg.StaticPath == "" {
		return fmt.Errorf("StaticPath must be set. ")
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

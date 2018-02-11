package config

import (
	"fmt"
	"regexp"

	"github.com/tkw1536/place/utils/verify"
)

// Validate checks if the configuration is valid
// and returns nil if it is, and an error message if it is not.
func (cfg Config) Validate() verify.ValidationError {

	var errors []verify.FieldError

	if cfg.GitURL.IsEmpty() {
		errors = append(errors, verify.FmtFieldError("GitURL", "GitURL may not be empty. "))
	}

	if cfg.GitUsername != "" && cfg.GitSSHKeyPath != "" {
		errors = append(errors,
			verify.FmtFieldError("GitUsername", "Cannot set GitUsername when GitSSHKeyPath is set, you may use at most one. "),
			verify.FmtFieldError("GitSSHKeyPath", "Cannot set GitSSHKeyPath when GitUsername is set, you may use at most one. "),
		)
	}

	if cfg.GitUsername != "" && cfg.GitPassword == "" {
		errors = append(errors, verify.FmtFieldError("GitPassword", "GitUsername is set, but no password was given. "))
	}

	if cfg.GitSSHKeyPath != "" {
		if err := verify.File(cfg.GitSSHKeyPath); err != nil {
			errors = append(errors, verify.AsFieldError("GitSSHKeyPath", err))
		}
	}

	if err := validateAddress(cfg.BindAddress); err != nil {
		errors = append(errors, verify.AsFieldError("BindAddress", err))
	}

	if err := verifyHookPath(cfg.WebhookPath); err != nil {
		errors = append(errors, verify.AsFieldError("WebhookPath", err))
	}

	if cfg.StaticPath == "" {
		errors = append(errors, verify.FmtFieldError("StaticPath", "StaticPath must be set. "))
	}

	if cfg.StaticPath != "" {
		if err := verify.Dir(cfg.StaticPath); err != nil {
			errors = append(errors, verify.AsFieldError("StaticPath", err))
		}
	}

	// return either an error of all the fields
	// or nil
	if len(errors) != 0 {
		return verify.NewValidationError(errors)
	}

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

package config

import (
	"fmt"

	"../../utils/verify"
)

// verifies the configuration settings
func validateConfig(cfg *Config) error {
	if cfg.RepositoryURL == nil {
		return fmt.Errorf("--from must be specified")
	}

	if err := verify.Dir(cfg.OutDirectory); err != nil {
		return err
	}

	if cfg.Ref == "" {
		return fmt.Errorf("--ref may not be empty")
	}

	if cfg.SSHKeyPath != "" {
		if err := verify.File(cfg.SSHKeyPath); err != nil {
			return err
		}
	}

	return nil
}

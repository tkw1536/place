package updater

import "./config"

// RunUpdate runs an update with the given configuration
func RunUpdate(cfg *config.Config) error {
	var err error
	if len(cfg.BuildScript) == 0 {
		err = updateWithGit(cfg)
	} else {
		err = updateWithScript(cfg)
	}

	return err
}

package updater

import (
	"context"

	"github.com/tkw1536/place/config"
)

// RunUpdate runs an update with the given configuration
func RunUpdate(ctx context.Context, cfg *config.Config) error {
	var err error
	if len(cfg.BuildScript) == 0 {
		err = updateWithGit(cfg)
	} else {
		err = updateWithScript(cfg)
	}

	return err
}

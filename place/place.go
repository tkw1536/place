package place

import (
	"./config"

	"../server"
)

// StartPlace starts the place system
func StartPlace(cfg *config.Config) error {

	// load or create ssh key
	if cfg.SSHKeyPath != "" {
		if err := loadOrCreateSSHKey(cfg.SSHKeyPath, cfg.Logger); err != nil {
			return err
		}
	}

	// turn this configuration into a server configuration
	scfg := cfg.ToServerConfig()
	return server.StartServer(&scfg)
}

package place

import (
	"../utils"
	"./config"

	"../server"
)

// StartPlace starts the place system
func StartPlace(cfg *config.Config) error {

	// load or create ssh key
	if cfg.SSHKeyPath != "" {
		if err := loadOrCreateSSHKey(cfg.SSHKeyPath, utils.Logger); err != nil {
			return err
		}
	}

	utils.Logger.Println("Done initializing, starting up server ...")
	// turn this configuration into a server configuration
	scfg := cfg.ToServerConfig()
	scfg.Inspect()
	return server.StartServer(&scfg)
}

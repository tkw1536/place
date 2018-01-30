package place

import "./config"

import serverconfig "../server/config"
import "../server"

// StartPlace starts the place system
func StartPlace(cfg *config.Config) error {

	// load or create ssh key
	if cfg.SSHKeyPath != "" {
		if err := loadOrCreateSSHKey(cfg); err != nil {
			return err
		}
	}

	// TODO: Make proper command line arguments
	var scfg serverconfig.Config

	if err := serverconfig.ParseConfig(&scfg, []string{"-help"}); err != nil {
		return err
	}

	return server.StartServer(&scfg)
}

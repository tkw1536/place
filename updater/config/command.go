package config

// ToCommand turns this configuration into a command line
// that can be used to call this executable
func (cfg Config) ToCommand(path string) []string {
	command := []string{
		path,
		"-from", cfg.RepositoryURL.String(),
		"-ssh-key", cfg.SSHKeyPath,
		"-to", cfg.OutDirectory,
		"-ref", cfg.Ref,
	}
	if cfg.BuildScript != "" {
		command = append(command, "-build", cfg.BuildScript)
	}
	return command
}

package config

// ToCommand turns this configuration into a command line
// that can be used to call this executable
func (cfg Config) ToCommand(path string) []string {
	return []string{
		path,
		"-from", cfg.RepositoryURL,
		"-ssh-key", cfg.SSHKeyPath,
		"-to", cfg.OutDirectory,
		"-ref", cfg.Ref,
		"-build", cfg.BuildScript,
	}
}

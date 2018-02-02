package updater

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"../utils/git"

	"./config"
)

func updateWithScript(cfg *config.Config) error {
	// create a temporary folder
	tmpDir, err := ioutil.TempDir("", "update")
	if err != nil {
		return err
	}
	defer os.RemoveAll(tmpDir)

	cfg.Logger.Printf("cloning %s into %s", cfg.RepositoryURL, tmpDir)

	// clone into it
	if _, err := git.Clone(tmpDir, cfg.RepositoryURL, cfg.Ref, false, cfg.SSHKeyPath); err != nil {
		return err
	}

	// get the shell variable
	shell := os.Getenv("SHELL")
	if shell == "" {
		shell = "/bin/sh"
	}

	cfg.Logger.Printf("running build script")

	// and run the build script
	cmd := exec.Command(shell, "-c", cfg.BuildScript+" "+cfg.OutDirectory)
	cmd.Dir = tmpDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// start the command
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("Failed to run hook: %s", err.Error())
	}

	return cmd.Wait()
}
package updater

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/tkw1536/place/utils"
	git "gopkg.in/src-d/go-git.v4"

	"github.com/tkw1536/place/config"
)

func updateWithScript(ctx *context.Context, cfg *config.Config) error {
	// create a temporary folder
	tmpDir, err := ioutil.TempDir("", "update")
	if err != nil {
		return err
	}
	defer os.RemoveAll(tmpDir)

	// load options
	opts, err := cfg.GitCloneOptions()
	if err != nil {
		return err
	}

	// and clone
	utils.Logger.Printf("cloning %s into %s", cfg.GitURL, tmpDir)

	if _, err := git.PlainCloneContext(*ctx, tmpDir, false, opts); err != nil {
		return err
	}

	// get the shell variable
	shell := os.Getenv("SHELL")
	if shell == "" {
		shell = "/bin/sh"
	}

	utils.Logger.Printf("running build script")

	// and run the build script
	cmd := exec.CommandContext(*ctx, shell, "-c", cfg.BuildScript+" "+cfg.StaticPath)
	cmd.Dir = tmpDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// start the command
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("Failed to run hook: %s", err.Error())
	}

	return cmd.Wait()
}

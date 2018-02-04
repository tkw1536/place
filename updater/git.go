package updater

import (
	"context"
	"io/ioutil"
	"os"

	"github.com/tkw1536/place/utils"

	"github.com/tkw1536/place/config"
	"gopkg.in/src-d/go-billy.v4/osfs"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

// clone `source`, revision `refrev`, and force check it out at dest
func updateWithGit(ctx *context.Context, cfg *config.Config) error {
	// create a temporary script
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

	utils.Logger.Printf("cloning %s into %s", cfg.GitURL.String(), tmpDir)

	// do a bare clone into it
	r, err := git.PlainCloneContext(*ctx, tmpDir, true, opts)
	if err != nil {
		return err
	}

	// open a new repository
	cor, err := git.Open(r.Storer, osfs.New(cfg.StaticPath))
	if err != nil {
		panic(err.Error())
	}

	//
	wtr, err := cor.Worktree()
	if err != nil {
		panic(err.Error())
	}

	rev, err := cor.ResolveRevision(plumbing.Revision(cfg.GitRef()))
	if err != nil {
		return err
	}

	// and do a reset
	var reset git.ResetOptions
	reset.Mode = git.HardReset
	reset.Commit = *rev

	utils.Logger.Printf("checking out %s in %s", rev, cfg.StaticPath)

	return wtr.Reset(&reset)
}

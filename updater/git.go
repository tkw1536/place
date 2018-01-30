package updater

import (
	"io/ioutil"
	"os"

	"./config"
	"gopkg.in/src-d/go-billy.v4/osfs"
	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

// clone `source`, revision `refrev`, and force check it out at dest
func updateWithGit(cfg *config.Config) error {
	// create a temporary script
	tmpDir, err := ioutil.TempDir("", "update")
	if err != nil {
		return err
	}
	defer os.RemoveAll(tmpDir)

	cfg.Logger.Printf("cloning %s into %s", cfg.RepositoryURL, tmpDir)

	// do a bare clone into it
	r, err := cloneRepo(tmpDir, cfg.RepositoryURL, cfg.Ref, true, cfg.SSHKeyPath)
	if err != nil {
		return err
	}

	// open a new repository
	cor, err := git.Open(r.Storer, osfs.New(cfg.OutDirectory))
	if err != nil {
		panic(err.Error())
	}

	//
	wtr, err := cor.Worktree()
	if err != nil {
		panic(err.Error())
	}

	rev, err := cor.ResolveRevision(plumbing.Revision(cfg.Ref))
	if err != nil {
		return err
	}

	// and do a reset
	var reset git.ResetOptions
	reset.Mode = git.HardReset
	reset.Commit = *rev

	cfg.Logger.Printf("checking out %s in %s", rev, cfg.OutDirectory)

	return wtr.Reset(&reset)
}

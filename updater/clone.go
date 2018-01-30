package updater

import (
	"../utils/sshkey"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

// clone a repository to a local path
// using an optional ssh key for authentication
func cloneRepo(path string, repository string, ref string, bare bool, keyPath string) (*git.Repository, error) {

	var options git.CloneOptions
	options.URL = repository
	options.ReferenceName = plumbing.ReferenceName(ref)
	options.SingleBranch = true
	options.Depth = 1

	if keyPath != "" {
		auth, err := sshkey.Get(keyPath)
		if err != nil {
			return nil, err
		}
		options.Auth = auth
	}

	return git.PlainClone(path, bare, &options)
}

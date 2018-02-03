package git

import (
	"os"

	"../sshkey"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/transport"
)

// Clone clones a repository to a local path
// using a given ref and optionally a given ssh key
func Clone(path string, repository string, ref string, bare bool, keyPath string) (*git.Repository, error) {

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

	repo, err := git.PlainClone(path, bare, &options)

	// if cloning fails because of an invalid auth method (e.g. when using ssh keys on a git repository)
	// then try the same thing again
	if err == transport.ErrInvalidAuthMethod && options.Auth != nil {
		os.RemoveAll(path)
		options.Auth = nil
		return git.PlainClone(path, bare, &options)
	}

	return repo, err
}

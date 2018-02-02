package sshkey

import (
	"fmt"
	"os/user"

	"gopkg.in/src-d/go-git.v4/plumbing/transport/ssh"
)

// Get gets the ssh key for a repository from a local path.
// assuming the ssh key does not have a passphrase
func Get(path string) (k *ssh.PublicKeys, e error) {
	// read the current username
	usr, err := user.Current()
	if err != nil {
		return nil, nil
	}

	// if we paniced, we should setup an error
	defer func() {
		if r := recover(); r != nil {
			k = nil
			e = fmt.Errorf("Unable to load ssh key: %s", r)
		}
	}()

	// and read in the ssh key
	return ssh.NewPublicKeysFromFile(usr.Username, path, "")
}

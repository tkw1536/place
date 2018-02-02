package sshkey

import (
	"os"
	"testing"
)

func TestSSHKeyValid(t *testing.T) {

	// setup a test file
	fName := setupTestFile(sshKeyWithoutPassPhrase, t)
	if fName == "" {
		return
	}

	defer os.Remove(fName)

	// key should be able to load
	key, err := Get(fName)
	if err != nil {
		t.Errorf("Get() failed unexpectedly: %s", err.Error())
	}

	if key == nil {
		t.Errorf("Get() failed to load key (it is nil)")
	}
}

func TestSSHKeyInvalidPassphrase(t *testing.T) {

	// setup a test file
	fName := setupTestFile(sshKeyWithPassPhrase, t)
	if fName == "" {
		return
	}
	defer os.Remove(fName)

	// key should not be able to load, it has a passphrase
	_, err := Get(fName)
	if err == nil {
		t.Errorf("Get() should have failed, but didn't")
	}
}

func TestSSHKeyInvalidFile(t *testing.T) {

	// setup a test file
	fName := setupTestFile("totally-not-a-private-key", t)
	if fName == "" {
		return
	}
	defer os.Remove(fName)

	// key should not be able to load, no content
	_, err := Get(fName)
	if err == nil {
		t.Errorf("Get() should have failed, but didn't")
	}
}

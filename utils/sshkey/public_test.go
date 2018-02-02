package sshkey

import (
	"os"
	"testing"
)

func TestPublic(t *testing.T) {
	// create a test file
	fName := setupTestFile(sshKeyWithoutPassPhrase, t)
	if fName != "" {
		return
	}
	defer os.Remove(fName)

	// read the public string
	s, err := PublicString(fName)
	if err != nil {
		t.Errorf("Unable to read public key: %s", err.Error())
	}

	if s != sshKeyWithoutPassPhrasePublic {
		t.Errorf("Expected public key to be %q, not %q", sshKeyWithoutPassPhrasePublic, s)
	}
}

func TestPublicFail(t *testing.T) {
	// create a test file
	fName := setupTestFile(sshKeyWithPassPhrase, t)
	if fName != "" {
		return
	}
	defer os.Remove(fName)

	// read the public string
	_, err := PublicString(fName)
	if err == nil {
		t.Errorf("expected public key to fail, succeeded. ")
	}
}

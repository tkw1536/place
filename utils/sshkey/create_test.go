package sshkey

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/tkw1536/place/utils/verify"
)

func TestCreate(t *testing.T) {
	// Make a temporary name
	dName, err := ioutil.TempDir("", "id_rsa")
	if err != nil {
		t.Errorf("Failed to create temp file: %s", err.Error())
	}
	defer os.RemoveAll(dName)

	// get private and public key path
	privateKeyPath := dName + "/id_rsa"
	publicKeyPath := privateKeyPath + ".pub"

	// create a key pair
	if err := CreateKeyPair(privateKeyPath, 2048); err != nil {
		t.Errorf("Unable to create key pair: %s", err.Error())
		return
	}

	// public key written
	if verify.File(publicKeyPath) != nil {
		t.Errorf("Public key file not written")
		return
	}

	// private key written
	if verify.File(privateKeyPath) != nil {
		t.Errorf("Private key file not written")
		return
	}

	// read the public key path
	fc, err := ioutil.ReadFile(publicKeyPath)
	if err != nil {
		t.Errorf("Unable to read public key: %s", err.Error())
		return
	}

	// get the public key from the private key
	publicKey, err := PublicString(privateKeyPath)
	if err != nil {
		t.Errorf("Unable to get public key signature: %s", err.Error())
		return
	}

	// and check that the content is identical
	if string(fc) != publicKey {
		t.Errorf("Incorrect public key, got %s but expected %s", string(fc), publicKey)
	}
}

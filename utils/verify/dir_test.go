package verify

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestDirExists(t *testing.T) {
	// create a dir to check for existence
	dPath, err := ioutil.TempDir(os.TempDir(), "test")
	if err != nil {
		t.Errorf("Unable to create testing dir: %s", err.Error())
		return
	}

	defer os.RemoveAll(dPath)

	// verifying that a dir exists
	actual := Dir(dPath)
	if actual != nil {
		t.Errorf("Dir Verification should succeed, but fails: %s", actual.Error())
	}
}

func TestDirMissing(t *testing.T) {
	// create a dir and remove it to make sure it doesn't exist
	dPath, err := ioutil.TempDir(os.TempDir(), "test")
	if err != nil {
		t.Errorf("Unable to create testing dir: %s", err.Error())
		return
	}

	os.RemoveAll(dPath)

	// verifying that a file does not exist
	actual := Dir(dPath)
	if actual == nil {
		t.Errorf("Dir Verification should fail, but passes: %s", dPath)
	}
}

func TestDirFile(t *testing.T) {
	// create a file to check for existence
	f, err := ioutil.TempFile(os.TempDir(), "test")
	if err != nil {
		t.Errorf("Unable to create testing file: %s", err.Error())
		return
	}

	fPath := f.Name()
	defer os.Remove(fPath)

	// verifying that a file fails a dir test
	actual := Dir(fPath)
	if actual == nil {
		t.Errorf("Dir Verification should fail, but succeeds: %s", fPath)
	}
}

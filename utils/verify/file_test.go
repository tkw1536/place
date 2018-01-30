package verify

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestFileExists(t *testing.T) {
	// create a file to check for existence
	f, err := ioutil.TempFile(os.TempDir(), "test")
	if err != nil {
		t.Errorf("Unable to create testing file: %s", err.Error())
		return
	}

	fPath := f.Name()
	defer os.Remove(fPath)

	// verifying that a file exists
	actual := File(fPath)
	if actual != nil {
		t.Errorf("File Verification should succeed, but fails: %s", actual.Error())
	}
}

func TestFileMissing(t *testing.T) {
	// create and delete a file to make sure it doesn't exist
	f, err := ioutil.TempFile(os.TempDir(), "test")
	if err != nil {
		t.Errorf("Unable to create testing file: %s", err.Error())
		return
	}

	fPath := f.Name()
	os.Remove(fPath)

	// verifying that a file does not exist
	actual := File(fPath)
	if actual == nil {
		t.Errorf("File Verification should fail, but succeeds: %s", fPath)
	}
}

func TestFileDir(t *testing.T) {
	// create a directory
	dPath, err := ioutil.TempDir(os.TempDir(), "test")
	if err != nil {
		t.Errorf("Unable to create testing dir: %s", err.Error())
		return
	}

	defer os.RemoveAll(dPath)

	// verifying that a directory doesn't pass the file test
	actual := File(dPath)
	if actual == nil {
		t.Errorf("File Verification should fail, but succeeds: %s", dPath)
	}
}

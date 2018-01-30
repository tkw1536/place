package config

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

// a simple test that parsing command line arguments works
func TestParseShort(t *testing.T) {
	expected := setupBasicTest(t)
	defer os.RemoveAll(expected.OutDirectory)

	var actual Config
	ParseConfig(&actual, []string{
		"--from", "git@github.com/dummy/repository.git",
		"--ref", "refs/head/master",
		"--to", expected.OutDirectory,
	})

	// so that deep equal works properly
	actual.Logger = expected.Logger

	if !reflect.DeepEqual(*expected, actual) {
		t.Errorf("expected configuration to parse properly. ")
	}
}

// a longer test that parsing command line arguments works
func TestParseLong(t *testing.T) {
	expected := setupBasicTest(t)
	defer os.RemoveAll(expected.OutDirectory)

	f, err := ioutil.TempFile(os.TempDir(), "test")
	if err != nil {
		t.Errorf("Unable to create testing file: %s", err.Error())
		return
	}

	expected.SSHKeyPath = f.Name()
	defer os.Remove(expected.SSHKeyPath)

	expected.BuildScript = "dummy-build-script"

	var actual Config
	ParseConfig(&actual, []string{
		"--from", "git@github.com/dummy/repository.git",
		"--ref", "refs/head/master",
		"--ssh-key", expected.SSHKeyPath,
		"--build", expected.BuildScript,
		"--to", expected.OutDirectory,
	})

	// so that deep equal works properly
	actual.Logger = expected.Logger

	if !reflect.DeepEqual(*expected, actual) {
		t.Errorf("expected configuration to parse properly. ")
	}
}

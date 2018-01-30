package config

import (
	"io/ioutil"
	"os"
	"testing"
)

func assertValidateSuccess(cfg *Config, t *testing.T) {
	if err := validateConfig(cfg); err != nil {
		t.Errorf("Configuration validation failed: %s", err.Error())
	}
}

func assertValidateFailure(cfg *Config, t *testing.T) {
	if err := validateConfig(cfg); err == nil {
		t.Errorf("Configuration validation succeeded, but expected failure. ")
	}
}

func setupBasicTest(t *testing.T) *Config {
	// create a dir to use as out dir
	dPath, err := ioutil.TempDir(os.TempDir(), "test")
	if err != nil {
		t.Errorf("Unable to create testing dir: %s", err.Error())
		return nil
	}

	// a very simple configuration that should pass
	var cfg Config
	cfg.RepositoryURL = "git@github.com/dummy/repository.git"
	cfg.Ref = "refs/head/master"
	cfg.SSHKeyPath = ""
	cfg.OutDirectory = dPath
	cfg.BuildScript = ""

	return &cfg
}

// a simple configuration should pass
func TestValidateSuccess(t *testing.T) {
	cfg := setupBasicTest(t)
	defer os.RemoveAll(cfg.OutDirectory)

	assertValidateSuccess(cfg, t)
}

// invalid repository url -- may not be empty
func TestValidateRepo(t *testing.T) {
	cfg := setupBasicTest(t)
	defer os.RemoveAll(cfg.OutDirectory)

	cfg.RepositoryURL = ""

	assertValidateFailure(cfg, t)
}

// invalid ref -- may not be empty
func TestValidateRef(t *testing.T) {
	cfg := setupBasicTest(t)
	defer os.RemoveAll(cfg.OutDirectory)

	cfg.Ref = ""

	assertValidateFailure(cfg, t)
}

// non-existing ssh key path
func TestValidateSSHKeyPathMissing(t *testing.T) {
	cfg := setupBasicTest(t)
	defer os.RemoveAll(cfg.OutDirectory)

	cfg.SSHKeyPath = "non-existent-file"

	assertValidateFailure(cfg, t)
}

func TestValidateSSHKeyPathFound(t *testing.T) {
	cfg := setupBasicTest(t)
	defer os.RemoveAll(cfg.OutDirectory)

	f, err := ioutil.TempFile(os.TempDir(), "test")
	if err != nil {
		t.Errorf("Unable to create testing file: %s", err.Error())
		return
	}

	fPath := f.Name()
	defer os.Remove(fPath)

	cfg.SSHKeyPath = fPath

	assertValidateSuccess(cfg, t)
}

// non-existing out directory
func TestValidateOutDirectory(t *testing.T) {
	cfg := setupBasicTest(t)
	os.RemoveAll(cfg.OutDirectory)
	assertValidateFailure(cfg, t)
}

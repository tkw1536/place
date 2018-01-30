package verify

import (
	"testing"
)

func TestCommandExistsWithEmpty(t *testing.T) {
	// verifying that a command exists when allowempty is true
	actual := Command([]string{"/bin/sh", "-c", "echo"}, true)
	if actual != nil {
		t.Errorf("Command Verification should succeed, but fails: %s", actual.Error())
	}
}

func TestCommandExistsWithoutEmpty(t *testing.T) {
	// verifying that a command exists when allow empty is false
	actual := Command([]string{"/bin/sh", "-c", "echo"}, false)
	if actual != nil {
		t.Errorf("Command Verification should succeed, but fails: %s", actual.Error())
	}
}

func TestCommandMissingWithEmpty(t *testing.T) {
	// verifying that an invalid command fails with allowEmpty being true
	actual := Command([]string{"/junk-command"}, true)
	if actual == nil {
		t.Errorf("Command Verification should fail, but succeeds")
	}
}

func TestCommandMissingWithoutEmpty(t *testing.T) {
	// verifying that an invalid command fails with allowEmpty being false
	actual := Command([]string{"/junk-command"}, false)
	if actual == nil {
		t.Errorf("Command Verification should fail, but succeeds")
	}
}

func TestCommandEmptyWithEmpty(t *testing.T) {
	// verifying that an empty command is allowed when the flag is set
	actual := Command([]string{}, true)
	if actual != nil {
		t.Errorf("Command Verification should succeed, but fails: %s", actual.Error())
	}
}

func TestCommandEmptyWithoutEmpty(t *testing.T) {
	// verifying that an empty command is not allowed when the flag is not set
	actual := Command([]string{}, false)
	if actual == nil {
		t.Errorf("Command Verification should fail, but succeeds")
	}
}

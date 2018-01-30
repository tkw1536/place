package command

import (
	"io/ioutil"
	"os"
	"testing"
	"time"
)

func TestNonCommand(t *testing.T) {
	// create a non-executable file
	f, err := ioutil.TempFile(os.TempDir(), "test")
	if err != nil {
		t.Errorf("Unable to create testing file: %s", err.Error())
		return
	}

	fPath := f.Name()
	defer os.Remove(fPath)

	_, err = WithTimeout(time.Second, fPath)
	if err == nil {
		t.Errorf("expected command to fail, but suceeded")
	}
}

func TestRegularCommand(t *testing.T) {
	_, err := WithTimeout(time.Second, "true")
	if err != nil {
		t.Errorf("expected command to suceeded, but failed: %s", err.Error())
	}
}

func TestIrregularCommand(t *testing.T) {
	_, err := WithTimeout(time.Second, "false")
	if err == nil {
		t.Errorf("expected command to fail, but suceeded")
	}
}

func TestTimedOut(t *testing.T) {
	_, err := WithTimeout(time.Second, "sleep", "10")
	if err == nil {
		t.Errorf("expected command to timeout, but suceeded")
	}
}

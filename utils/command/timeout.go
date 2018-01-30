package command

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

// WithTimeout runs a command with a specific timeout
func WithTimeout(timeout time.Duration, args ...string) (*exec.Cmd, error) {
	// create command to run
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// start the command or die
	if err := cmd.Start(); err != nil {
		return nil, fmt.Errorf("failed to run command: %s", err.Error())
	}

	// create and wait for a result
	var err error
	done := make(chan error)
	go func() { done <- cmd.Wait() }()

	select {
	case e := <-done:
		err = e
	case <-time.After(timeout):
		cmd.Process.Kill()
		err = fmt.Errorf("failed to run command %s, timeout expired after %d second(s)", args, timeout/time.Second)
	}
	return cmd, err
}

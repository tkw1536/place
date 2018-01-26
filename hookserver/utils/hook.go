package utils

import (
	"os/exec"
	"os"
	"sync"
	"log"
	"time"
	"fmt"
)

var lock sync.Mutex

// runs the actual hook to update stuff
func RunHook (hook []string, timeout int, Info *log.Logger) {
	Info.Print("queuing hook")

	// we only run one hook at a time
	lock.Lock()
	defer lock.Unlock()

	Info.Print("running hook\n")

	// create command to run
	cmd := exec.Command(hook[0], hook[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// start the command or die
	if err := cmd.Start(); err != nil {
		Info.Printf("failed to start hook: %s\n", err.Error())
		return
	}

	// create a timeout for the error
	var err error
	done := make(chan error)
	go func() { done <- cmd.Wait() }()
	select {
		case e := <-done:
			err = e
		case <-time.After(time.Duration(timeout) * time.Second):
			cmd.Process.Kill()
			err = fmt.Errorf("timeout expired after %d second(s)", timeout)
	}

	// error handling
	if err != nil {
		Info.Printf("hook failed to run: %s\n", err.Error())
	} else {
		Info.Print("hook finished")
	}
}
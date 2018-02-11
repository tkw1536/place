package main

// runChild runs the child Process with the given arguments
import (
	"os"
	"os/exec"

	"github.com/tkw1536/place/utils"
)

func runChild() int {
	cmd := childCommand[0]
	args := childCommand[1:]
	args = append(args, configPath)

	utils.Logger.Printf("Running child process %v with args %v", cmd, args)

	child := exec.Command(cmd, args...)

	// Loop through the output of the child process
	child.Stdout = os.Stdout
	child.Stderr = os.Stderr
	child.Stdin = os.Stdin

	err := child.Run()

	if err == nil {
		return 0
	}

	utils.Logger.Printf("Child process failed to exit cleanly: %v", err)
	return 1
}

package verify

import (
	"fmt"
)

// Command checks that a command is valid
func Command(command []string, allowEmpty bool) error {
	if len(command) == 0 {
		if allowEmpty {
			return nil
		}
		return fmt.Errorf("Command must be defined")
	}

	return File(command[0])
}

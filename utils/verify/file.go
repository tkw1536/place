package verify

import (
	"fmt"
	"os"
)

// File verifies that a path exists
func File(path string) error {
	stat, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("Path does not exist: %q", path)
	}

	if !stat.Mode().IsRegular() {
		return fmt.Errorf("Path is not a file: %q", path)
	}

	return nil
}

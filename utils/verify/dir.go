package verify

import (
	"fmt"
	"os"
)

// Dir verifies that a directory exists
func Dir(path string) error {
	stat, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("Path does not exist: %q", path)
	}

	if !stat.Mode().IsDir() {
		return fmt.Errorf("Path is not a directory: %q", path)
	}

	return nil
}

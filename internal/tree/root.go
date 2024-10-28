package tree

import (
	"fmt"
	"os"
	"path/filepath"
)

func getRootDirectory(starting string) (string, error) {
	if starting == "" {
		currentDir, err := os.Getwd()
		if err != nil {
			return "", fmt.Errorf("error getting current directory: %w", err)
		}
		starting = currentDir
	}

	for {
		filePath := filepath.Join(starting, "templar.yaml")

		// Check if the file exists
		if _, err := os.Stat(filePath); err == nil {
			return starting, nil
		}

		parentDir := filepath.Dir(starting)
		if starting == parentDir {
			break
		}

		starting = parentDir
	}

	return "", fmt.Errorf("file 'templar.yaml' not found in any parent directories")
}

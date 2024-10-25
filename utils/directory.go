package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetRootDirectory() (string, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("error getting current directory: %w", err)
	}
	for {
		filePath := filepath.Join(currentDir, "templar.yaml")

		// Check if the file exists
		if _, err := os.Stat(filePath); err == nil {
			// File found, return the full path
			return currentDir, nil
		}

		if currentDir == filepath.Dir(currentDir) {
			break
		}

		currentDir = filepath.Dir(currentDir)
	}

	// File not found
	return "", fmt.Errorf("file 'templar.yaml' not found in any parent directories")
}

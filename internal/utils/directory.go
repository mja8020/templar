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
			return currentDir, nil
		}

		parentDir := filepath.Dir(currentDir)
		if currentDir == parentDir {
			break
		}

		currentDir = parentDir
	}

	return "", fmt.Errorf("file 'templar.yaml' not found in any parent directories")
}

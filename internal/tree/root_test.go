package tree

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestGetRootDirectory tests the GetRootDirectory function
func TestGetRootDirectory(t *testing.T) {
	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "testrootdir")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	// Create a nested directory structure
	nestedDir := filepath.Join(tempDir, "nested", "deeper")
	err = os.MkdirAll(nestedDir, os.ModePerm)
	require.NoError(t, err)

	// Resolve symlinks for nestedDir
	// This allows tests to run on native osx, as /tmp links to /private/tmp
	nestedDir, err = filepath.EvalSymlinks(nestedDir)
	require.NoError(t, err)

	// Set up test cases
	testCases := []struct {
		name           string
		setup          func() string
		expectedOutput string
		expectedError  string
	}{
		{
			name: "File in current directory",
			setup: func() string {
				// Create the templar.yaml file in the nested directory
				filePath := filepath.Join(nestedDir, "templar.yaml")
				os.WriteFile(filePath, []byte("content"), 0o644)
				return nestedDir
			},
			expectedOutput: nestedDir,
			expectedError:  "",
		},
		{
			name: "File in parent directory",
			setup: func() string {
				// Create the templar.yaml file in the parent directory ("nested")
				parentDir := filepath.Dir(nestedDir)
				filePath := filepath.Join(parentDir, "templar.yaml")
				os.WriteFile(filePath, []byte("content"), 0o644)
				return nestedDir
			},
			expectedOutput: filepath.Dir(nestedDir),
			expectedError:  "",
		},
		{
			name: "File not found",
			setup: func() string {
				// Ensure no templar.yaml file is present
				return nestedDir
			},
			expectedOutput: "",
			expectedError:  "file 'templar.yaml' not found in any parent directories",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Set up the environment for the test case
			startDir := tc.setup()

			// Call the function
			result, err := getRootDirectory(startDir)

			// Clean up by removing the templar.yaml file if it exists
			for _, dir := range []string{nestedDir, filepath.Dir(nestedDir)} {
				filePath := filepath.Join(dir, "templar.yaml")
				os.Remove(filePath) // Ignore errors here as the file might not exist
			}

			if tc.expectedError == "" {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedOutput, result)
			} else {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tc.expectedError)
			}
		})
	}
}

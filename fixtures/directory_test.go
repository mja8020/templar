package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/mja8020/templar/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestGetRootDirectory tests the GetRootDirectory function
func TestGetRootDirectory(t *testing.T) {
	// Create a temporary directory for testing
	tempDir, err := ioutil.TempDir("", "testrootdir")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir) // Clean up the temporary directory afterward

	// Create a nested directory structure
	nestedDir := filepath.Join(tempDir, "nested", "deeper")
	err = os.MkdirAll(nestedDir, os.ModePerm)
	require.NoError(t, err)

	// Set up test cases
	testCases := []struct {
		name           string
		setup          func() string // function to set up the environment
		expectedOutput string
		expectedError  string
	}{
		{
			name: "File in current directory",
			setup: func() string {
				// Create the templar.yaml file in the nested directory
				filePath := filepath.Join(nestedDir, "templar.yaml")
				ioutil.WriteFile(filePath, []byte("content"), 0o644)
				return nestedDir
			},
			expectedOutput: nestedDir,
			expectedError:  "",
		},
		{
			name: "File in parent directory",
			setup: func() string {
				// Create the templar.yaml file in the parent directory
				parentDir := filepath.Dir(nestedDir) // This should be "nested"
				filePath := filepath.Join(parentDir, "templar.yaml")
				ioutil.WriteFile(filePath, []byte("content"), 0o644)
				return nestedDir // Start from the deeper directory
			},
			expectedOutput: filepath.Dir(nestedDir), // Expect the parent directory
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

			// Change to the start directory
			err := os.Chdir(startDir)
			require.NoError(t, err)

			// Call the function
			result, err := utils.GetRootDirectory()

			// Normalize the paths
			normalizedResult, err1 := filepath.EvalSymlinks(result)
			require.NoError(t, err1)

			normalizedExpected, err2 := filepath.EvalSymlinks(tc.expectedOutput)
			require.NoError(t, err2)

			// Assert the output
			if tc.expectedError == "" {
				assert.NoError(t, err)
				assert.Equal(t, normalizedExpected, normalizedResult)
			} else {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tc.expectedError)
			}
		})
	}
}

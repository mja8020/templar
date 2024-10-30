package exec

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRunSuccess(t *testing.T) {
	// Define a simple executable command that should succeed
	executable := "echo"
	args := []string{"Hello, world!"}
	env := map[string]string{}
	stream := false
	fileOutputStream := "os.Stdout"

	result, err := Run(executable, env, args, stream, fileOutputStream)
	require.NoError(t, err, "expected no error during execution")
	assert.Equal(t, 0, result.ExitCode, "expected exit code 0")
	assert.Equal(t, "Hello, world!\n", result.StdOut, "unexpected stdout content")
}

func TestRunCommandNotFound(t *testing.T) {
	// Use a non-existent command
	executable := "nonexistentcommand"
	args := []string{}
	env := map[string]string{}
	stream := false
	fileOutputStream := "os.Stdout"

	_, err := Run(executable, env, args, stream, fileOutputStream)
	assert.Error(t, err, "expected error for non-existent command")
}

func TestRunPermissionDenied(t *testing.T) {
	// Use an existing file that isn't executable
	executable := "/etc/passwd" // A non-executable file on Unix systems
	args := []string{}
	env := map[string]string{}
	stream := false
	fileOutputStream := "os.Stdout"

	_, err := Run(executable, env, args, stream, fileOutputStream)
	assert.Error(t, err, "expected permission error for non-executable file")
}

func TestRunWithStreaming(t *testing.T) {
	// Create a temporary file for streaming output
	tmpFile, err := os.CreateTemp("", "output.txt")
	require.NoError(t, err, "failed to create temp file")
	defer os.Remove(tmpFile.Name())

	// Test a simple command with streaming
	executable := "echo"
	args := []string{"Stream test"}
	env := map[string]string{}
	stream := true
	fileOutputStream := tmpFile.Name()

	result, err := Run(executable, env, args, stream, fileOutputStream)
	require.NoError(t, err, "expected no error during execution")
	assert.Equal(t, 0, result.ExitCode, "expected exit code 0")

	// Verify output was streamed to the file
	content, err := os.ReadFile(tmpFile.Name())
	require.NoError(t, err, "failed to read temp file")
	assert.Equal(t, "Stream test\n", string(content), "unexpected file content")
}

func TestRunErrorOutput(t *testing.T) {
	executable := "ls"
	args := []string{"/nonexistent_directory"}
	env := map[string]string{}
	stream := false
	fileOutputStream := "os.Stdout"

	result, _ := Run(executable, env, args, stream, fileOutputStream)

	assert.Contains(t, result.StdErr, "No such file or directory", "unexpected stderr content")
}

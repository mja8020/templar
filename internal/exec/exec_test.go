package exec

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRunSuccess(t *testing.T) {
	// Create a context with a timeout for the command
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Define the executable and arguments
	executable := "echo"
	args := []string{"Hello, world!"}
	env := map[string]string{}
	stream := false
	fileOutputStream := "os.Stdout"

	// Call the Run function with context
	result, err := Run(ctx, executable, env, args, stream, fileOutputStream)
	require.NoError(t, err, "expected no error during execution")
	assert.Equal(t, 0, result.ExitCode, "expected exit code 0")
	assert.Equal(t, "Hello, world!\n", result.StdOut, "unexpected stdout content")
}

func TestRunCommandNotFound(t *testing.T) {
	// Create a context with a timeout for the command
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Use a non-existent command
	executable := "nonexistentcommand"
	args := []string{}
	env := map[string]string{}
	stream := false
	fileOutputStream := "os.Stdout"

	_, err := Run(ctx, executable, env, args, stream, fileOutputStream)
	assert.Error(t, err, "expected error for non-existent command")
}

func TestRunPermissionDenied(t *testing.T) {
	// Create a context with a timeout for the command
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Use an existing file that isn't executable
	executable := "/etc/passwd" // A non-executable file on Unix systems
	args := []string{}
	env := map[string]string{}
	stream := false
	fileOutputStream := "os.Stdout"

	_, err := Run(ctx, executable, env, args, stream, fileOutputStream)
	assert.Error(t, err, "expected permission error for non-executable file")
}

func TestRunWithStreaming(t *testing.T) {
	// Create a context with a timeout for the command
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

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

	result, err := Run(ctx, executable, env, args, stream, fileOutputStream)
	require.NoError(t, err, "expected no error during execution")
	assert.Equal(t, 0, result.ExitCode, "expected exit code 0")

	// Verify output was streamed to the file
	content, err := os.ReadFile(tmpFile.Name())
	require.NoError(t, err, "failed to read temp file")
	assert.Equal(t, "Stream test\n", string(content), "unexpected file content")
}

func TestRunErrorOutput(t *testing.T) {
	// Create a context with a timeout for the command
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	executable := "ls"
	args := []string{"/nonexistent_directory"}
	env := map[string]string{}
	stream := false
	fileOutputStream := "os.Stdout"

	result, _ := Run(ctx, executable, env, args, stream, fileOutputStream)

	assert.Contains(t, result.StdErr, "No such file or directory", "unexpected stderr content")
}

package exec

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExec_Success(t *testing.T) {
	// Test a command that should succeed
	executable := "/bin/bash"
	environment := map[string]string{
		"TEST_VAR": "Hello, World!",
	}
	args := []string{"-c", "echo $TEST_VAR"}

	// Execute the command
	result, err := Exec(executable, environment, args)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, "Hello, World!\n", result.StdOut)
	assert.Equal(t, "", result.StdErr)
	assert.Equal(t, 0, result.ExitCode)
}

func TestExec_CommandNotFound(t *testing.T) {
	// Test a command that does not exist
	executable := "/bin/doesnotexist"
	environment := map[string]string{}
	args := []string{}

	// Execute the command
	result, err := Exec(executable, environment, args)

	// Assertions
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to start command")
	assert.Equal(t, "", result.StdOut)
	assert.Equal(t, "", result.StdErr)
	assert.NotEqual(t, 0, result.ExitCode)
}

func TestExec_FailWithStdErr(t *testing.T) {
	// Test a command that should fail and produce stderr output
	executable := "/bin/bash"
	environment := map[string]string{}
	args := []string{"-c", "ls /nonexistent"}

	// Execute the command
	result, err := Exec(executable, environment, args)

	// Assertions
	assert.Error(t, err)
	assert.Contains(t, result.StdErr, "No such file or directory")
	assert.Equal(t, 2, result.ExitCode) // Expect exit code 2 for "no such file or directory"
}

func TestExec_WithEnvironmentVariables(t *testing.T) {
	// Test a command that uses environment variables
	executable := "/bin/bash"
	environment := map[string]string{
		"MY_VAR": "TestValue",
	}
	args := []string{"-c", "echo $MY_VAR"}

	// Execute the command
	result, err := Exec(executable, environment, args)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, "TestValue\n", result.StdOut)
	assert.Equal(t, "", result.StdErr)
	assert.Equal(t, 0, result.ExitCode)
}

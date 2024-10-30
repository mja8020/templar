package exec

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"syscall"
)

type Execute struct {
	StdOut   string
	StdErr   string
	ExitCode int
}

func Run(ctx context.Context, executable string, environment map[string]string, args []string, stream bool, fileOutputStream string) (Execute, error) {
	var result Execute

	// Find execpath
	execPath, err := exec.LookPath(executable)
	if err != nil {
		return result, fmt.Errorf("executable not found or not in PATH: %s", executable)
	}

	// Set up command
	cmd := exec.CommandContext(ctx, execPath, args...)

	// Set up environment variables
	env := os.Environ()
	for key, value := range environment {
		env = append(env, fmt.Sprintf("%s=%s", key, value))
	}
	cmd.Env = env

	var outputStream io.Writer
	if fileOutputStream == "os.Stdout" {
		outputStream = os.Stdout
	} else {
		// Assume it's a file path
		file, err := os.Create(fileOutputStream)
		if err != nil {
			return result, fmt.Errorf("failed to open file for output: %v", err)
		}
		defer file.Close()
		outputStream = file
	}

	// Get stdout and stderr pipes
	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		result.ExitCode = -1
		return result, fmt.Errorf("failed to get stdout pipe: %v", err)
	}
	stderrPipe, err := cmd.StderrPipe()
	if err != nil {
		result.ExitCode = -1
		return result, fmt.Errorf("failed to get stderr pipe: %v", err)
	}

	// Start the command
	if err := cmd.Start(); err != nil {
		result.ExitCode = -1 // Set a non-zero exit code on start failure
		return result, fmt.Errorf("failed to start command: %v", err)
	}

	// Channel to signal when output collection is done
	done := make(chan struct{}, 2)

	// Function to collect output quietly or stream it
	collectOutput := func(pipe io.Reader, output *string) {
		var buf bytes.Buffer
		scanner := bufio.NewScanner(pipe)
		for scanner.Scan() {
			line := scanner.Text()
			if stream && outputStream != nil {
				fmt.Fprintln(outputStream, line) // Write to outputStream
			}
			buf.WriteString(line + "\n")
		}
		if err := scanner.Err(); err != nil {
			fmt.Printf("Error reading pipe: %v\n", err)
		}
		*output = buf.String()
		done <- struct{}{}
	}

	// Collect stdout and stderr concurrently
	go collectOutput(stdoutPipe, &result.StdOut)
	go collectOutput(stderrPipe, &result.StdErr)

	// Wait for the command to finish
	if err := cmd.Wait(); err != nil {
		fmt.Printf("Detected error from cmd.Wait(): %v\n", err)
		if exitError, ok := err.(*exec.ExitError); ok {
			if status, ok := exitError.Sys().(syscall.WaitStatus); ok {
				result.ExitCode = status.ExitStatus()
			}
		} else {
			result.ExitCode = -1
		}
	} else {
		result.ExitCode = 0
	}

	// Wait for both stdout and stderr to be fully collected
	<-done
	<-done

	return result, err
}

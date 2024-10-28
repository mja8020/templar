package exec

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"syscall"
)

type Execute struct {
	StdOut   string
	StdErr   string
	ExitCode int
}

func isExecutable(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.Mode()&0111 != 0 // Check for any executable bit
}

func Run(executable string, environment map[string]string, args []string, stream bool) (Execute, error) {
	var result Execute

	execPath, err := exec.LookPath(executable)
	if err != nil {
		return result, fmt.Errorf("executable not found or not in PATH: %s", executable)
	}
	if !isExecutable(execPath) {
		return result, fmt.Errorf("executable is not permitted to run: %s", execPath)
	}

	// Create the command
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/C", fmt.Sprintf("%s %s", execPath, strings.Join(args, " ")))
	default: // Linux and macOS
		cmd = exec.Command("/bin/bash", "-c", fmt.Sprintf("%s %s", execPath, strings.Join(args, " ")))
	}

	// Set up environment variables
	env := os.Environ()
	for key, value := range environment {
		env = append(env, fmt.Sprintf("%s=%s", key, value))
	}
	cmd.Env = env

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
	done := make(chan struct{})

	if stream != false {
		// Stream stdout
		go func() {
			scanner := bufio.NewScanner(stdoutPipe)
			for scanner.Scan() {
				line := scanner.Text()
				fmt.Println(line)            // Print each line to stdout
				result.StdOut += line + "\n" // Collect stdout output
			}
			done <- struct{}{}
		}()

		// Stream stderr
		go func() {
			scanner := bufio.NewScanner(stderrPipe)
			for scanner.Scan() {
				line := scanner.Text()
				fmt.Fprintln(os.Stderr, line) // Print each line to stderr
				result.StdErr += line + "\n"  // Collect stderr output
			}
			done <- struct{}{}
		}()
	}

	// Wait for the command to finish
	if err := cmd.Wait(); err != nil {
		// Capture the exit code if the command fails
		if exitError, ok := err.(*exec.ExitError); ok {
			if status, ok := exitError.Sys().(syscall.WaitStatus); ok {
				result.ExitCode = status.ExitStatus()
			}
		} else {
			result.ExitCode = -1
		}
		return result, err
	}

	result.ExitCode = 0

	// Wait for both stdout and stderr to be fully collected
	<-done
	<-done

	return result, nil
}

package command

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
)

type Command interface {
	Run(name string, args ...string) (string, error)
	RunAsync(name string, args ...string) (string, error)
}

type cmd struct{}

// RunAsync implements Command.
func (*cmd) RunAsync(name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", fmt.Errorf("error creating StdoutPipe: %v", err)
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return "", fmt.Errorf("error creating StderrPipe: %v", err)
	}

	if err := cmd.Start(); err != nil {
		return "", fmt.Errorf("error starting command: %v", err)
	}

	go printStream("stdout", stdout)
	go printStream("stderr", stderr)

	if err := cmd.Wait(); err != nil {
		return "", fmt.Errorf("Command finished with error: %v", err)
	}
	return "", nil
}

// Run implements Command.
func (*cmd) Run(name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	output, err := cmd.CombinedOutput()

	// Check if the command was successful
	if err != nil {
		// Print the standard error output
		fmt.Printf("Error output: %s\n", cmd.Stderr)

		return "", err
	}

	return string(output), err
}

func printStream(streamName string, reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Printf("[%s] %s\n", streamName, scanner.Text())
	}
}

func New() Command {
	return &cmd{}
}

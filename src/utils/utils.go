package utils

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func CommandExists(command string) bool {
	_, err := exec.LookPath(command)
	return err == nil
}

func DirExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}

	return info.IsDir()
}

func ExecuteCommand(verbose bool, command string, args ...string) error {
	var stdoutBuf, stderrBuf bytes.Buffer
	cmd := exec.Command(command, args...)
	cmd.Stdout = &stdoutBuf
	cmd.Stderr = &stderrBuf
	err := cmd.Run()

	if err != nil {
		fmt.Printf("Error executing command: %s %v | %v\n", command, args, err)
		if verbose {
			fmt.Println("Command output:", stderrBuf.String())
		}
	}

	if verbose {
		fmt.Println("Command output:", stdoutBuf.String())
	}

	return err
}

package utils

import (
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

func ExecuteCommand(command string, args ...string) error {
	cmd := exec.Command(command, args...)
	err := cmd.Run()
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err != nil {
		fmt.Printf("Error executing command: %s %v | %v\n", command, args, err)
	}

	return err
}

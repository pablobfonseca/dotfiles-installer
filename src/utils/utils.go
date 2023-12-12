package utils

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

var DotfilesRepo = "https://github.com/pablobfonseca/dotfiles.git"

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
	var stdoutBuf, stderrBuf bytes.Buffer
	cmd := exec.Command(command, args...)
	cmd.Stdout = &stdoutBuf
	cmd.Stderr = &stderrBuf
	err := cmd.Run()

	if err != nil {
		fmt.Println("Error executing command:", err)
		fmt.Println("Command output:", stdoutBuf.String())
	}

	return err
}

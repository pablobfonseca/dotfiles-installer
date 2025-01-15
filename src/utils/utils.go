package utils

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func CommandExists(command string) bool {
	_, err := exec.LookPath(command)
	return err == nil
}

func FileExists(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}

	return true
}

func Confirm(prompt string) bool {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("%s (y/n): ", prompt)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}
		input = strings.TrimSpace(strings.ToLower(input))
		if input == "y" || input == "yes" {
			return true
		} else if input == "n" || input == "no" {
			return false
		} else {
			fmt.Println("Please enter 'y' or 'n'.")
		}
	}
}

func DirExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}

	return info.IsDir()
}

func ClearTerminal() {
	fmt.Print("\033[H\033[2J]]")
}

func ExecuteCommand(command string, args ...string) error {
	ClearTerminal()
	cmd := exec.Command(command, args...)
	var out strings.Builder
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if err := cmd.Run(); err != nil {
		fmt.Printf("Error executing command: %s %v | %v\n", command, args, err)
		return err
	}

	fmt.Print(out.String())
	return nil
}

func SymlinkFiles(src, dest string) error {
	if err := os.Symlink(src, dest); err != nil {
		return err
	}

	return nil
}

func FindInOutput(output, query string) (bool, string) {
	lines := strings.Split(output, "\n")
	for _, line := range lines {
		if strings.Contains(line, query) {
			fields := strings.Fields(line)
			if len(fields) > 0 {
				return true, fields[0]
			}
		}
	}

	return false, ""
}

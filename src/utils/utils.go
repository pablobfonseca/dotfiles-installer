package utils

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strings"
)

var isDryRunFunc func() bool
var isNonInteractiveMode bool = false
var forceMode bool = false

func SetDryRunChecker(f func() bool) {
	isDryRunFunc = f
}

func IsDryRun() bool {
	if isDryRunFunc != nil {
		return isDryRunFunc()
	}
	return false
}

func SetNonInteractiveMode(nonInteractive bool) {
	isNonInteractiveMode = nonInteractive
}

func IsNonInteractive() bool {
	return isNonInteractiveMode
}

func SetForceMode(force bool) {
	forceMode = force
}

func IsForceMode() bool {
	return forceMode
}

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
	return ConfirmWithDryRunBehavior(prompt, false)
}

// ConfirmDestructive should be used for operations that modify/delete files
func ConfirmDestructive(prompt string) bool {
	return ConfirmWithDryRunBehavior(prompt, true)
}

func ConfirmWithDryRunBehavior(prompt string, skipOnDryRun bool) bool {
	// In force mode, always return true
	if IsForceMode() {
		InfoMessage(fmt.Sprintf("%s -> Forced (--force flag)", prompt))
		return true
	}

	// In dry-run mode, handle based on operation type
	if IsDryRun() {
		if skipOnDryRun {
			fmt.Printf("[DRY RUN] Would ask: %s -> Skipping destructive operation\n", prompt)
			return false // Skip destructive operations in dry-run
		} else {
			fmt.Printf("[DRY RUN] Would ask: %s -> Simulating 'yes'\n", prompt)
			return true // Simulate 'yes' for non-destructive operations
		}
	}

	// In non-interactive mode (like TUI), default to 'yes' to avoid blocking
	if IsNonInteractive() {
		InfoMessage(fmt.Sprintf("%s -> Auto-confirmed (non-interactive mode)", prompt))
		return true
	}

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
	if IsDryRun() {
		fmt.Printf("[DRY RUN] Would execute: %s %s\n", command, strings.Join(args, " "))
		return nil
	}

	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if err := cmd.Run(); err != nil {
		fmt.Printf("Error executing command: %s %v | %v\n", command, args, err)
		return err
	}

	return nil
}

func SymlinkFiles(src, dest string) error {
	if IsDryRun() {
		fmt.Printf("[DRY RUN] Would create symlink: %s -> %s\n", dest, src)
		return nil
	}

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

func GetCurrentUser() string {
	currentUser, err := user.Current()
	if err != nil {
		return "unknown"
	}
	return currentUser.Username
}

func RemoveAllFiles(path string) error {
	if IsDryRun() {
		fmt.Printf("[DRY RUN] Would remove: %s\n", path)
		return nil
	}

	if IsNonInteractive() {
		InfoMessage("Removing existing file/directory: " + path)
	}

	return os.RemoveAll(path)
}

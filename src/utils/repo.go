package utils

import (
	"os"
	"os/exec"

	"github.com/pablobfonseca/dotfiles/src/config"
)

func CloneRepoIfNotExists(repo, dest string) {
	if DirExists(dest) {
		SkipMessage("Clone destination already exists")
		return
	}

	InfoMessage("Cloning...")
	if err := ExecuteCommand("git", "clone", repo, dest); err != nil {
		ErrorMessage("Error cloning the repository", err)
	}
}

func UpdateRepository() error {
	cmd := exec.Command("git", "status", "--short")
	cmd.Dir = config.DotfilesConfigDir()

	out, err := cmd.Output()
	if err != nil {
		return err
	}

	if string(out[:]) == "" {
		if err := pullFromRepo(); err != nil {
			return err
		}
	} else {
		err := stash("-u")
		if err != nil {
			return err
		}

		err = pullFromRepo()
		if err != nil {
			return err
		}

		err = stash("pop")
		if err != nil {
			return err
		}
	}

	return nil
}

func pullFromRepo() error {
	cmd := exec.Command("git", "pull")
	cmd.Dir = config.DotfilesConfigDir()
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func stash(args ...string) error {
	stashArgs := append([]string{"stash"}, args...)

	cmd := exec.Command("git", stashArgs...)
	cmd.Dir = config.DotfilesConfigDir()
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

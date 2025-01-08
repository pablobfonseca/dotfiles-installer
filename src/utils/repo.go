package utils

import (
	"github.com/pablobfonseca/dotfiles/src/config"
)

func CloneRepoIfNotExists(repo, dest string) {
	if repo == "" {
		repo = config.RepositoryUrl()
	}

	if dest == "" {
		dest = config.DotfilesConfigDir()
	}

	if DirExists(dest) {
		SkipMessage("Clone destination already exists")
		return
	}

	InfoMessage("Cloning...")
	if err := ExecuteCommand("git", "clone", repo, dest); err != nil {
		ErrorMessage("Error cloning the repository", err)
	}
}

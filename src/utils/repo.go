package utils

import (
	"github.com/pablobfonseca/dotfiles/src/config"
)

func CloneRepoIfNotExists(verbose bool) {
	if DirExists(config.DotfilesConfigDir()) {
		SkipMessage("Dotfiles directory already exists")
		return
	}

	InfoMessage("Dotfiles directory does not exists, cloning...")
	if err := ExecuteCommand(verbose, "git", "clone", config.RepositoryUrl(), config.DotfilesConfigDir()); err != nil {
		ErrorMessage("Error cloning the repository", err)
	}
}

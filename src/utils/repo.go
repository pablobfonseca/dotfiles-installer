package utils

func CloneRepoIfNotExists() {
	if DirExists(DotfilesPath) {
		SkipMessage("Dotfiles directory already exists")
		return
	}
	InfoMessage("Dotfiles directory does not exists, cloning...")
	if err := ExecuteCommand("git", "clone", DotfilesRepo, DotfilesPath); err != nil {
		ErrorMessage("Error cloning the repository", err)
	}
}

package dotfiles

import (
	"os"
	"path"

	"github.com/pablobfonseca/dotfiles-cli/src/utils"
	"github.com/vbauerster/mpb/v7"
)

func CloneRepo(p *mpb.Progress) {
	bar := utils.NewBar("Cloning dotfiles repo", 1, p)

	if err := utils.ExecuteCommand("git", "clone", utils.DotfilesRepo, path.Join(os.Getenv("HOME"), ".dotfiles")); err != nil {
		utils.ErrorMessage("Error cloning the repository", err)
	}
	bar.Increment()
}

func DeleteRepo(p *mpb.Progress) {
	bar := utils.NewBar("Deleting dotfiles repo", 1, p)

	if err := utils.ExecuteCommand("rm", "-rf", path.Join(os.Getenv("HOME"), ".dotfiles")); err != nil {
		utils.ErrorMessage("Error deleting the repository", err)
	}
	bar.Increment()
}

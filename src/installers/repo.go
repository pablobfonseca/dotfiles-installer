package dotfiles

import (
	"github.com/pablobfonseca/dotfiles/src/config"
	"github.com/pablobfonseca/dotfiles/src/utils"
	"github.com/vbauerster/mpb/v7"
)

func CloneRepo(p *mpb.Progress, verbose bool) {
	bar := utils.NewBar("Cloning dotfiles repo", 1, p)

	if err := utils.ExecuteCommand(verbose, "git", "clone", config.RepositoryUrl(), config.DotfilesConfigDir()); err != nil {
		utils.ErrorMessage("Error cloning the repository", err)
	}
	bar.Increment()
}

func DeleteRepo(p *mpb.Progress, verbose bool) {
	bar := utils.NewBar("Deleting dotfiles repo", 1, p)

	if err := utils.ExecuteCommand(verbose, "rm", "-rf", config.DotfilesConfigDir()); err != nil {
		utils.ErrorMessage("Error deleting the repository", err)
	}
	bar.Increment()
}

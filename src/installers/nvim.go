package dotfiles

import (
	"github.com/pablobfonseca/dotfiles/src/config"
	"github.com/pablobfonseca/dotfiles/src/utils"
	"github.com/vbauerster/mpb/v7"
)

func UninstallNvim(uninstallApp bool, p *mpb.Progress, verbose bool) {
	if !utils.CommandExists("nvim") {
		utils.SkipMessage("nvim is not installed")
		return
	}

	if uninstallApp {
		uninstallBar := utils.NewBar("Uninstalling nvim", 1, p)

		if err := utils.ExecuteCommand(verbose, "brew", "uninstall", "neovim"); err != nil {
			utils.ErrorMessage("Error uninstalling nvim", err)
		}
		uninstallBar.Increment()
	}

	removeFilesBar := utils.NewBar("Removing nvim files", 1, p)
	if err := utils.ExecuteCommand(verbose, "rm", "-rf", config.NvimConfigDir()); err != nil {
		utils.ErrorMessage("Error removing nvim files", err)
	}
	removeFilesBar.Increment()
}

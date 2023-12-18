package dotfiles

import (
	"os"
	"path"

	"github.com/pablobfonseca/dotfiles/src/config"
	"github.com/pablobfonseca/dotfiles/src/utils"
	"github.com/vbauerster/mpb/v7"
)

var nvChadRepo = "https://github.com/NvChad/NvChad"

func InstallNvim(p *mpb.Progress, verbose bool) {
	if nvimInstalled() {
		utils.SkipMessage("nvim already installed")
	} else {
		installNvimBar := utils.NewBar("Installing nvim", 1, p)
		if err := utils.ExecuteCommand(verbose, "brew", "install", "neovim"); err != nil {
			utils.ErrorMessage("Error installing nvim", err)
		}
		installNvimBar.Increment()
	}

	utils.CloneRepoIfNotExists(verbose)

	if utils.DirExists(config.NvimConfigDir()) {
		utils.SkipMessage("nvim files already exists")
		return
	}
	installNvChadBar := utils.NewBar("Installing NvChad", 1, p)

	if err := utils.ExecuteCommand(verbose, "git", "clone", "--depth", "1", nvChadRepo, config.NvimConfigDir()); err != nil {
		utils.ErrorMessage("Error cloning the repository", err)
	}
	installNvChadBar.Increment()

	symlinkBar := utils.NewBar("Symlinking files", 1, p)

	src := path.Join(config.DotfilesConfigDir(), "nvim", "custom")
	dest := path.Join(config.NvimConfigDir(), "lua", "custom")
	if err := os.Symlink(src, dest); err != nil {
		utils.ErrorMessage("Error creating symlink", err)
	}
	symlinkBar.Increment()

}

func UninstallNvim(uninstallApp bool, p *mpb.Progress, verbose bool) {
	if !nvimInstalled() {
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

func nvimInstalled() bool {
	return utils.CommandExists("nvim")
}

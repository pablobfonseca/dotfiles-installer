package dotfiles

import (
	"os"
	"path"

	"github.com/pablobfonseca/dotfiles-cli/src/utils"
	"github.com/vbauerster/mpb/v7"
)

var dotfilesPath = path.Join(os.Getenv("HOME"), ".dotfiles")
var nvimConfigPath = path.Join(os.Getenv("HOME"), ".config", "nvim")
var nvChadRepo = "https://github.com/NvChad/NvChad"

func InstallNvim(p *mpb.Progress) {
	if nvimInstalled() {
		utils.SkipMessage("nvim already installed")
		return
	}

	if !utils.DirExists(dotfilesPath) {
		cloningBar := utils.NewBar("Cloning dotfiles", 1, p)
		utils.InfoMessage("Dotfiles directory does not exists, cloning...")
		if err := utils.ExecuteCommand("git", "clone", utils.DotfilesRepo, dotfilesPath); err != nil {
			utils.ErrorMessage("Error cloning the repository", err)
		}
		cloningBar.Increment()
	}

	installNvimBar := utils.NewBar("Installing nvim", 1, p)
	if err := utils.ExecuteCommand("brew", "install", "neovim"); err != nil {
		utils.ErrorMessage("Error installing nvim", err)
	}
	installNvimBar.Increment()

	if utils.DirExists(nvimConfigPath) {
		utils.SkipMessage("nvim files already exists")
		return
	}
	installNvChadBar := utils.NewBar("Installing NvChad", 1, p)

	if err := utils.ExecuteCommand("git", "clone", "--depth", "1", nvChadRepo, nvimConfigPath); err != nil {
		utils.ErrorMessage("Error cloning the repository", err)
	}
	installNvChadBar.Increment()

	symlinkBar := utils.NewBar("Symlinking files", 1, p)

	src := path.Join(dotfilesPath, "nvim", "custom")
	dest := path.Join(nvimConfigPath, "lua", "custom")
	if err := os.Symlink(src, dest); err != nil {
		utils.ErrorMessage("Error creating symlink", err)
	}
	symlinkBar.Increment()

}

func UninstallNvim(p *mpb.Progress) {
	if !nvimInstalled() {
		utils.SkipMessage("nvim is not installed")
		return
	}

	uninstallBar := utils.NewBar("Uninstalling nvim", 2, p)

	if err := utils.ExecuteCommand("brew", "uninstall", "neovim"); err != nil {
		utils.ErrorMessage("Error uninstalling nvim", err)
	}
	uninstallBar.Increment()

	if err := utils.ExecuteCommand("rm", "-rf", nvimConfigPath); err != nil {
		utils.ErrorMessage("Error removing nvim files", err)
	}
	uninstallBar.Increment()
}

func nvimInstalled() bool {
	return utils.CommandExists("nvim")
}

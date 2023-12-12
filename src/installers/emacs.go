package dotfiles

import (
	"os"
	"path"

	"github.com/pablobfonseca/dotfiles-cli/src/utils"
	"github.com/vbauerster/mpb/v7"
)

var emacsConfigPath = path.Join(os.Getenv("HOME"), ".emacs.d")

func InstallEmacs(p *mpb.Progress) {
	if emacsInstalled() {
		utils.SkipMessage("Emacs is already installed")
	} else {
		installEmacsBar := utils.NewBar("Installing emacs", 1, p)

		if err := utils.ExecuteCommand("brew", "install", "--cask", "emacs"); err != nil {
			utils.ErrorMessage("Error installing emacs symlink", err)
		}
		installEmacsBar.Increment()

	}

	utils.CloneRepoIfNotExists()

	symlinkBar := utils.NewBar("Symlinking files", 1, p)

	src := path.Join(utils.DotfilesPath, "emacs.d")
	dest := path.Join(emacsConfigPath)
	if err := os.Symlink(src, dest); err != nil {
		utils.ErrorMessage("Error creating symlink", err)
	}
	symlinkBar.Increment()
}

func UninstallEmacs(p *mpb.Progress) {
	if !emacsInstalled() {
		utils.SkipMessage("Emacs is not installed")
		return
	}

	uninstallBar := utils.NewBar("Uninstalling emacs", 2, p)

	if err := utils.ExecuteCommand("brew", "uninstall", "emacs"); err != nil {
		utils.ErrorMessage("Error uninstalling emacs", err)
	}
	uninstallBar.Increment()

	if err := utils.ExecuteCommand("rm", "-rf", emacsConfigPath); err != nil {
		utils.ErrorMessage("Error removing emacs files", err)
	}
	uninstallBar.Increment()
}

func emacsInstalled() bool {
	return utils.DirExists(path.Join("/Applications", "Emacs.app"))
}

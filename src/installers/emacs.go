package dotfiles

import (
	"os"
	"path"

	"github.com/pablobfonseca/dotfiles/src/config"
	"github.com/pablobfonseca/dotfiles/src/utils"
	"github.com/vbauerster/mpb/v7"
)

func InstallEmacs(p *mpb.Progress, verbose bool) {
	if emacsInstalled() {
		utils.SkipMessage("Emacs is already installed")
	} else {
		installEmacsBar := utils.NewBar("Installing emacs", 1, p)

		if err := utils.ExecuteCommand(verbose, "brew", "install", "--cask", "emacs"); err != nil {
			utils.ErrorMessage("Error installing emacs symlink", err)
		}
		installEmacsBar.Increment()

	}

	utils.CloneRepoIfNotExists(verbose)

	symlinkBar := utils.NewBar("Symlinking files", 1, p)

	src := path.Join(config.DotfilesConfigDir(), "emacs.d")
	dest := path.Join(config.EmacsConfigDir())
	if err := os.Symlink(src, dest); err != nil {
		utils.ErrorMessage("Error creating symlink", err)
	}
	symlinkBar.Increment()
}

func UninstallEmacs(uninstallApp bool, p *mpb.Progress, verbose bool) {
	if !emacsInstalled() {
		utils.SkipMessage("Emacs is not installed")
		return
	}

	uninstallBar := utils.NewBar("Uninstalling emacs", 1, p)

	if uninstallApp {
		if err := utils.ExecuteCommand(verbose, "brew", "uninstall", "emacs"); err != nil {
			utils.ErrorMessage("Error uninstalling emacs", err)
		}
		uninstallBar.Increment()

	}

	removeFilesBar := utils.NewBar("Removing emacs files", 1, p)
	if err := utils.ExecuteCommand(verbose, "rm", "-rf", config.EmacsConfigDir()); err != nil {
		utils.ErrorMessage("Error removing emacs files", err)
	}
	removeFilesBar.Increment()
}

func emacsInstalled() bool {
	return utils.DirExists(path.Join("/Applications", "Emacs.app"))
}

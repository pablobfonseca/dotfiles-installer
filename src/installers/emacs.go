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

	if utils.DirExists(emacsConfigPath) {
		utils.SkipMessage("Emacs folder already exists")
		return
	}

	cloningBar := utils.NewBar("Cloning emacs files", 1, p)
	utils.InfoMessage("Emacs directory does not exists, cloning...")
	if err := utils.ExecuteCommand("git", "clone", "https://github.com/pablobfonseca/emacs.d", emacsConfigPath); err != nil {
		utils.ErrorMessage("Error cloning the repository", err)
	}
	cloningBar.Increment()
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

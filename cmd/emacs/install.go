package emacs

import (
	"os"
	"path"

	"github.com/pablobfonseca/dotfiles/src/config"
	"github.com/pablobfonseca/dotfiles/src/utils"
	"github.com/spf13/cobra"
	"github.com/vbauerster/mpb/v7"
)

var InstallEmacsCmd = &cobra.Command{
	Use:   "emacs",
	Short: "Install emacs files",
	Run: func(cmd *cobra.Command, args []string) {
		p := mpb.New()

		verbose, _ := cmd.Flags().GetBool("verbose")

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
	},
}

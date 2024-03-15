package emacs

import (
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

		utils.CloneRepoIfNotExists(verbose, config.EmacsRepositoryUrl(), config.EmacsConfigDir())
	},
}

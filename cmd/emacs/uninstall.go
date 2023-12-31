package emacs

import (
	"github.com/pablobfonseca/dotfiles/src/config"
	"github.com/pablobfonseca/dotfiles/src/utils"
	"github.com/spf13/cobra"
	"github.com/vbauerster/mpb/v7"
)

var UninstallEmacsCmd = &cobra.Command{
	Use:   "emacs",
	Short: "Uninstall emacs files",
	Run: func(cmd *cobra.Command, args []string) {
		p := mpb.New()

		verbose, _ := cmd.Flags().GetBool("verbose")
		uninstallApp, _ := cmd.Flags().GetBool("uninstall-app")

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
	},
}

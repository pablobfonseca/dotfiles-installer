package vim

import (
	"github.com/pablobfonseca/dotfiles/src/config"
	"github.com/pablobfonseca/dotfiles/src/utils"
	"github.com/spf13/cobra"
	"github.com/vbauerster/mpb/v7"
)

var UninstallVimCmd = &cobra.Command{
	Use:   "vim",
	Short: "Uninstall vim files",
	Run: func(cmd *cobra.Command, args []string) {
		p := mpb.New()

		verbose, _ := cmd.Flags().GetBool("verbose")
		uninstallApp, _ := cmd.Flags().GetBool("uninstall-app")

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
	},
}

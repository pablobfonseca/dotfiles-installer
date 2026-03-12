package cmd

import (
	"github.com/spf13/cobra"

	"github.com/pablobfonseca/dotfiles/cmd/nvim"
	"github.com/pablobfonseca/dotfiles/src/config"
	"github.com/pablobfonseca/dotfiles/src/utils"
)

var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstall the dotfiles",
	Long:  "Uninstall the dotfiles. You can uninstall all the dotfiles or just some of them.",
	Run: func(cmd *cobra.Command, args []string) {
		dotfilesDir := config.DotfilesConfigDir()
		if !utils.Confirm("Are you sure you want to delete " + dotfilesDir + "?") {
			utils.SkipMessage("Uninstall cancelled")
			return
		}

		if err := utils.RemoveAllFiles(dotfilesDir); err != nil {
			utils.ErrorMessage("Error deleting the repository", err)
		}
		utils.SuccessMessage("Dotfiles uninstalled successfully")
	},
}

func init() {
	rootCmd.AddCommand(uninstallCmd)

	uninstallCmd.AddCommand(nvim.UninstallNvimCmd)
}

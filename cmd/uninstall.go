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
		if err := utils.ExecuteCommand("rm", "-rf", config.DotfilesConfigDir()); err != nil {
			utils.ErrorMessage("Error deleting the repository", err)
		}
	},
}

var uninstallApp bool

func init() {
	rootCmd.AddCommand(uninstallCmd)

	uninstallCmd.AddCommand(nvim.UnInstallNvimCmd)
}

package nvim

import (
	"os"
	"path"

	"github.com/pablobfonseca/dotfiles/src/utils"
	"github.com/spf13/cobra"
)

var UnInstallNvimCmd = &cobra.Command{
	Use:   "nvim",
	Short: "Uninstall nvim files",
	Run: func(cmd *cobra.Command, args []string) {
		configDir, _ := os.UserConfigDir()
		uninstallApp, _ := cmd.Flags().GetBool("uninstall-app")

		if !utils.CommandExists("nvim") {
			utils.SkipMessage("nvim is not installed")
			return
		}

		if uninstallApp {
			if err := utils.ExecuteCommand("brew", "uninstall", "neovim"); err != nil {
				utils.ErrorMessage("Error uninstalling nvim", err)
			}
		}

		if err := utils.ExecuteCommand("rm", "-rf", path.Join(configDir, "nvim")); err != nil {
			utils.ErrorMessage("Error removing nvim files", err)
		}
	},
}

package nvim

import (
	"path/filepath"

	"github.com/pablobfonseca/dotfiles/src/config"
	"github.com/pablobfonseca/dotfiles/src/utils"
	"github.com/spf13/cobra"
)

var UninstallNvimCmd = &cobra.Command{
	Use:   "nvim",
	Short: "Uninstall nvim files",
	Run: func(cmd *cobra.Command, args []string) {
		configDir, err := config.ConfigDir()
		if err != nil {
			utils.ErrorMessage("Error getting config directory", err)
			return
		}
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

		if err := utils.RemoveAllFiles(filepath.Join(configDir, "nvim")); err != nil {
			utils.ErrorMessage("Error removing nvim files", err)
		}
	},
}

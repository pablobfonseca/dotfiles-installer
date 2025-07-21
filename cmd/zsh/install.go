package zsh

import (
	"github.com/pablobfonseca/dotfiles/src/installer"
	"github.com/pablobfonseca/dotfiles/src/utils"
	"github.com/spf13/cobra"
)

var InstallZshCmd = &cobra.Command{
	Use:   "zsh",
	Short: "Install zsh configuration files and source them",
	Long:  "Install zsh configuration files (.zshrc, .zprofile, .zlogin) and automatically source the .zshrc to apply changes immediately.",
	Run: func(cmd *cobra.Command, args []string) {
		err := installer.SetupZsh()
		if err != nil {
			utils.ErrorMessage("Error creating symlink", err)
		}
	},
}

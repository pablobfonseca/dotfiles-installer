package zsh

import (
	"github.com/pablobfonseca/dotfiles/src/installer"
	"github.com/pablobfonseca/dotfiles/src/utils"
	"github.com/spf13/cobra"
)

var InstallZshCmd = &cobra.Command{
	Use:   "zsh",
	Short: "Install zsh files",
	Run: func(cmd *cobra.Command, args []string) {
		err := installer.SetupZsh()
		if err != nil {
			utils.ErrorMessage("Error creating symlink", err)
		}
	},
}

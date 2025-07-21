package karabiner

import (
	"github.com/pablobfonseca/dotfiles/src/installer"
	"github.com/pablobfonseca/dotfiles/src/utils"
	"github.com/spf13/cobra"
)

var InstallKarabinerCmd = &cobra.Command{
	Use:   "karabiner",
	Short: "Install Karabiner-Elements and sync configuration",
	Long:  "Install Karabiner-Elements keyboard customization tool and sync configuration files from your dotfiles repository.",
	Run: func(cmd *cobra.Command, args []string) {
		// Install Karabiner-Elements
		err := installer.InstallKarabiner()
		if err != nil {
			utils.ErrorMessage("Error installing Karabiner-Elements", err)
			return
		}

		// Setup configuration
		err = installer.SetupKarabiner()
		if err != nil {
			utils.ErrorMessage("Error setting up Karabiner configuration", err)
			return
		}

		utils.SuccessMessage("Karabiner-Elements installed and configured successfully!")
	},
}
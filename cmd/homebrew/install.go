package homebrew

import (
	"fmt"

	"github.com/pablobfonseca/dotfiles/src/config"
	"github.com/pablobfonseca/dotfiles/src/utils"
	"github.com/spf13/cobra"
)

var InstallHomebrewCmd = &cobra.Command{
	Use:   "homebrew",
	Short: "Install homebrew",
	Run: func(cmd *cobra.Command, args []string) {
		if utils.CommandExists("brew") {
			utils.SkipMessage("Homebrew already installed")
		} else {
			if err := utils.ExecuteCommand("/bin/bash", "-c", "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"); err != nil {
				fmt.Println("Error installing homebrew:", err)
				return
			}
		}

		utils.InfoMessage("Installing packages...")
		if err := utils.ExecuteCommand("\\cat", config.DotfilesConfigDir(), "/homebrew", "|", "xargs", "brew", "install"); err != nil {
			fmt.Println("Error installing packages:", err)
			return
		}
	},
}

package homebrew

import (
	"fmt"

	"github.com/pablobfonseca/dotfiles/src/installer"
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
			if err := installer.InstallHomebrew(); err != nil {
				fmt.Println("Error installing homebrew:", err)
				return
			}
		}
	},
}

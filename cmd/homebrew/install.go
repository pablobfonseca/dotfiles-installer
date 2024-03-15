package homebrew

import (
	"fmt"

	"github.com/pablobfonseca/dotfiles/src/config"
	"github.com/pablobfonseca/dotfiles/src/utils"
	"github.com/spf13/cobra"
	"github.com/vbauerster/mpb/v7"
)

var InstallHomebrewCmd = &cobra.Command{
	Use:   "homebrew",
	Short: "Install homebrew",
	Run: func(cmd *cobra.Command, args []string) {
		p := mpb.New()
		verbose, _ := cmd.Flags().GetBool("verbose")
		bar := utils.NewBar("Installing homebrew", 1, p)

		if utils.CommandExists("brew") {
			utils.SkipMessage("Homebrew already installed")
		} else {
			if err := utils.ExecuteCommand(verbose, "/bin/bash", "-c", "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"); err != nil {
				fmt.Println("Error installing homebrew:", err)
				return
			}
			bar.Increment()
		}

		installPackagesBar := utils.NewBar("Installing packages", 1, p)
		utils.InfoMessage("Installing packages...")
		if err := utils.ExecuteCommand(verbose, "\\cat", config.DotfilesConfigDir(), "/homebrew", "|", "xargs", "brew", "install"); err != nil {
			fmt.Println("Error installing packages:", err)
			return
		}
		installPackagesBar.Increment()
	},
}

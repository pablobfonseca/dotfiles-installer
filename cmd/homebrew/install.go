package homebrew

import (
	"fmt"

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

		if err := utils.ExecuteCommand(verbose, "/bin/bash", "-c", "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"); err != nil {
			fmt.Println("Error installing homebrew:", err)
			return
		}
		bar.Increment()
	},
}

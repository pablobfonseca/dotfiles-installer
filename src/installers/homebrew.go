package dotfiles

import (
	"fmt"

	"github.com/pablobfonseca/dotfiles/src/utils"
	"github.com/vbauerster/mpb/v7"
)

func InstallHomebrew(p *mpb.Progress, verbose bool) {
	bar := utils.NewBar("Installing homebrew", 1, p)

	if err := utils.ExecuteCommand(verbose, "/bin/bash", "-c", "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"); err != nil {
		fmt.Println("Error installing homebrew:", err)
		return
	}
	bar.Increment()
}

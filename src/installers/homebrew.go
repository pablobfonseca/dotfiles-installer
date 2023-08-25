package dotfiles

import (
	"fmt"

	utils "github.com/pablobfonseca/dotfiles-cli/src"
	"github.com/vbauerster/mpb/v7"
	"github.com/vbauerster/mpb/v7/decor"
)

func InstallHomebrew(p *mpb.Progress) {
	bar := p.AddBar(100,
		mpb.PrependDecorators(
			decor.Name("Installing homebrew", decor.WC{W: len("Installing homebre") + 1, C: decor.DidentRight}),
		),
		mpb.AppendDecorators(
			decor.OnComplete(decor.EwmaETA(decor.ET_STYLE_GO, 60), "done"),
		),
	)

	if err := utils.ExecuteCommand("/bin/bash", "-c", "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"); err != nil {
		fmt.Println("Error installing homebrew:", err)
		return
	}
	bar.Increment()
}

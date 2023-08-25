package dotfiles

import (
	"fmt"
	"os"
	"path"

	utils "github.com/pablobfonseca/dotfiles-cli/src"
	"github.com/vbauerster/mpb/v7"
	"github.com/vbauerster/mpb/v7/decor"
)

func CloneRepo(p *mpb.Progress) {
	bar := p.AddBar(100,
		mpb.PrependDecorators(
			decor.Name("Cloning repository", decor.WC{W: len("Cloning") + 1, C: decor.DidentRight}),
		),
		mpb.AppendDecorators(
			decor.OnComplete(decor.EwmaETA(decor.ET_STYLE_GO, 60), "done"),
		),
	)

	if err := utils.ExecuteCommand("git", "clone", "https://github.com/pablobfonseca/dotfiles.git", path.Join(os.Getenv("HOME"), ".dotfiles")); err != nil {
		fmt.Println("Error cloning the repository:", err)
		return
	}
	bar.Increment() // 100 % since cloning is done
}

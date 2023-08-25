package dotfiles

import (
	"fmt"
	"os"
	"path"

	"github.com/vbauerster/mpb/v7"
	"github.com/vbauerster/mpb/v7/decor"
)

func InstallZsh(p *mpb.Progress) {
	bar := p.AddBar(100,
		mpb.PrependDecorators(
			decor.Name("Symlinking zsh files", decor.WC{W: len("Symlinking zsh files") + 1, C: decor.DidentRight}),
		),
		mpb.AppendDecorators(
			decor.OnComplete(decor.EwmaETA(decor.ET_STYLE_GO, 60), "done"),
		),
	)

	for _, file := range []string{"zshrc", "zshenv"} {
		src := path.Join(os.Getenv("HOME"), ".dotfiles", "zsh", file)
		dest := path.Join(os.Getenv("HOME"), "."+file)
		if err := os.Symlink(src, dest); err != nil {
			fmt.Println("Error creating symlink:", err)
			return
		}
	}
	bar.Increment()
}

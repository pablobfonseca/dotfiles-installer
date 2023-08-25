package dotfiles

import (
	"fmt"
	"os"
	"path"

	utils "github.com/pablobfonseca/dotfiles-cli/src"
	"github.com/vbauerster/mpb/v7"
	"github.com/vbauerster/mpb/v7/decor"
)

func InstallEmacs(p *mpb.Progress) {
	installEmacs := p.AddBar(100,
		mpb.PrependDecorators(
			decor.Name("Installing emacs", decor.WC{W: len("Installing emacs") + 1, C: decor.DidentRight}),
		),
		mpb.AppendDecorators(
			decor.OnComplete(decor.EwmaETA(decor.ET_STYLE_GO, 60), "done"),
		),
	)

	if err := utils.ExecuteCommand("brew", "--cask", "install", "emacs"); err != nil {
		fmt.Println("Error installing emacs:", err)
		return
	}
	installEmacs.Increment()

	bar := p.AddBar(100,
		mpb.PrependDecorators(
			decor.Name("Symlinking emacs folder", decor.WC{W: len("Symlinking emacs folder") + 1, C: decor.DidentRight}),
		),
		mpb.AppendDecorators(
			decor.OnComplete(decor.EwmaETA(decor.ET_STYLE_GO, 60), "done"),
		),
	)

	src := path.Join(os.Getenv("HOME"), ".dotfiles", "emacs.d")
	dest := path.Join(os.Getenv("HOME"), ".emacs.d")
	if err := os.Symlink(src, dest); err != nil {
		fmt.Println("Error creating symlink:", err)
		return
	}
	bar.Increment()
}

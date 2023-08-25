package dotfiles

import (
	"fmt"
	"os"
	"path"

	utils "github.com/pablobfonseca/dotfiles-cli/src"
	"github.com/vbauerster/mpb/v7"
	"github.com/vbauerster/mpb/v7/decor"
)

func InstallNvim(p *mpb.Progress) {
	installNvChadBar := p.AddBar(100,
		mpb.PrependDecorators(
			decor.Name("Installing nvchad", decor.WC{W: len("Installing nvchad") + 1, C: decor.DidentRight}),
		),
		mpb.AppendDecorators(
			decor.OnComplete(decor.EwmaETA(decor.ET_STYLE_GO, 60), "done"),
		),
	)

	if err := utils.ExecuteCommand("git", "clone", "https://github.com/NvChad/NvChad", path.Join(os.Getenv("HOME"), ".config", "nvim", "--depth", "1")); err != nil {
		fmt.Println("Error cloning the repository:", err)
		return
	}
	installNvChadBar.Increment()

	bar := p.AddBar(100,
		mpb.PrependDecorators(
			decor.Name("Symlinking nvim custom folder", decor.WC{W: len("Symlinking nvim custom folder") + 1, C: decor.DidentRight}),
		),
		mpb.AppendDecorators(
			decor.OnComplete(decor.EwmaETA(decor.ET_STYLE_GO, 60), "done"),
		),
	)

	src := path.Join(os.Getenv("HOME"), ".dotfiles", "nvim", "custom")
	dest := path.Join(os.Getenv("HOME"), ".config", "nvim", "lua", "custom")
	if err := os.Symlink(src, dest); err != nil {
		fmt.Println("Error creating symlink:", err)
		return
	}
	bar.Increment()
}

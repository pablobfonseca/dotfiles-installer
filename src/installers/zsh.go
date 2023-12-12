package dotfiles

import (
	"os"
	"path"

	"github.com/pablobfonseca/dotfiles-cli/src/utils"
	"github.com/vbauerster/mpb/v7"
)

func InstallZsh(p *mpb.Progress) {
	bar := utils.NewBar("Symlinking zsh files", 1, p)

	for _, file := range []string{"zshrc", "zshenv"} {
		src := path.Join(os.Getenv("HOME"), ".dotfiles", "zsh", file)
		dest := path.Join(os.Getenv("HOME"), "."+file)
		if err := os.Symlink(src, dest); err != nil {
			utils.ErrorMessage("Error creating symlink", err)
			return
		}
	}
	bar.Increment()
}

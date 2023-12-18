package dotfiles

import (
	"os"
	"path"

	"github.com/pablobfonseca/dotfiles/src/config"
	"github.com/pablobfonseca/dotfiles/src/utils"
	"github.com/vbauerster/mpb/v7"
)

var zshConfigPath = path.Join(os.Getenv("HOME"), "zsh")

func InstallZsh(p *mpb.Progress) {
	bar := utils.NewBar("Symlinking zsh files", 1, p)

	for _, file := range []string{"zshrc", "zshenv"} {
		src := path.Join(config.DotfilesConfigDir(), "zsh", file)
		dest := path.Join(os.Getenv("HOME"), "."+file)

		utils.InfoMessage("Syncing " + src + " to " + dest)
		if _, err := os.Stat(dest); err == nil {
			utils.SkipMessage("File already exists: " + dest)
			continue
		}

		if err := os.Symlink(src, dest); err != nil {
			utils.ErrorMessage("Error creating symlink", err)
			return
		}
	}
	bar.Increment()
}

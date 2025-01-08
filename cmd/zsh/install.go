package zsh

import (
	"os"
	"path"

	"github.com/pablobfonseca/dotfiles/src/config"
	"github.com/pablobfonseca/dotfiles/src/utils"
	"github.com/spf13/cobra"
)

var InstallZshCmd = &cobra.Command{
	Use:   "zsh",
	Short: "Install zsh files",
	Run: func(cmd *cobra.Command, args []string) {
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
	},
}

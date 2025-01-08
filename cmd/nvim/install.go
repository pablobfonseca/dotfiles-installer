package nvim

import (
	"os"
	"path"

	"github.com/pablobfonseca/dotfiles/src/config"
	"github.com/pablobfonseca/dotfiles/src/utils"
	"github.com/spf13/cobra"
)

var InstallNvimCmd = &cobra.Command{
	Use:   "nvim",
	Short: "Install nvim files",
	Run: func(cmd *cobra.Command, args []string) {
		if utils.CommandExists("nvim") {
			utils.SkipMessage("nvim already installed")
		} else {
			if err := utils.ExecuteCommand("brew", "install", "nvim"); err != nil {
				utils.ErrorMessage("Error installing nvim", err)
			}
		}

		utils.CloneRepoIfNotExists("", "")

		if utils.DirExists(config.NvimConfigDir()) {
			utils.SkipMessage("nvim files already exists")
			return
		}

		src := path.Join(config.DotfilesConfigDir(), "nvim")
		dest := path.Join(config.NvimConfigDir())
		if err := os.Symlink(src, dest); err != nil {
			utils.ErrorMessage("Error creating symlink", err)
		}

		utils.SuccessMessage("nvim files synced")
	},
}

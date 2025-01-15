package nvim

import (
	"log"
	"os"
	"path"

	"github.com/pablobfonseca/dotfiles/src/config"
	"github.com/pablobfonseca/dotfiles/src/installer"
	"github.com/pablobfonseca/dotfiles/src/utils"
	"github.com/spf13/cobra"
)

var InstallNvimCmd = &cobra.Command{
	Use:   "nvim",
	Short: "Install nvim files",
	Run: func(cmd *cobra.Command, args []string) {
		configDir, _ := os.UserConfigDir()
		err := installer.InstallNvim()
		if err != nil {
			log.Fatalf("[nvim]: %v", err)
		}

		utils.CloneRepoIfNotExists(config.RepositoryUrl(), config.DotfilesConfigDir())

		nvimDir := path.Join(configDir, "nvim")
		if utils.DirExists(nvimDir) {
			if utils.Confirm("Nvim files already exists, do you want to override them?") {
				err := os.Remove(nvimDir)
				if err != nil {
					utils.ErrorMessage("[dotfiles]: error removing dir", err)
				}
			} else {
				utils.SkipMessage("nvim files already exists")
				return
			}
		}

		src := path.Join(config.DotfilesConfigDir(), "nvim")
		dest := path.Join(configDir, "nvim")
		err = utils.SymlinkFiles(src, dest)
		if err != nil {
			utils.ErrorMessage("[dotfiles]: symlink error", err)
		}

		utils.SuccessMessage("nvim files synced")
	},
}

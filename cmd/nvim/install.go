package nvim

import (
	"log"
	"path/filepath"

	"github.com/pablobfonseca/dotfiles/src/config"
	"github.com/pablobfonseca/dotfiles/src/installer"
	"github.com/pablobfonseca/dotfiles/src/utils"
	"github.com/spf13/cobra"
)

var InstallNvimCmd = &cobra.Command{
	Use:   "nvim",
	Short: "Install nvim files",
	Run: func(cmd *cobra.Command, args []string) {
		configDir, err := config.ConfigDir()
		if err != nil {
			log.Fatalf("[nvim]: failed to get config directory: %v", err)
		}
		err = installer.InstallNvim()
		if err != nil {
			log.Fatalf("[nvim]: %v", err)
		}

		utils.CloneRepoIfNotExists(config.RepositoryUrl(), config.DotfilesConfigDir())

		nvimDir := filepath.Join(configDir, "nvim")
		if utils.DirExists(nvimDir) {
			if utils.ConfirmDestructive("Nvim files already exists, do you want to override them?") {
				err := utils.RemoveAllFiles(nvimDir)
				if err != nil {
					utils.ErrorMessage("[dotfiles]: error removing dir", err)
					return
				}
			} else {
				utils.SkipMessage("nvim files already exists")
				return
			}
		}

		src := filepath.Join(config.DotfilesConfigDir(), "nvim")
		dest := filepath.Join(configDir, "nvim")
		err = utils.SymlinkFiles(src, dest)
		if err != nil {
			utils.ErrorMessage("[dotfiles]: symlink error", err)
		}

		utils.SuccessMessage("nvim files synced")
	},
}

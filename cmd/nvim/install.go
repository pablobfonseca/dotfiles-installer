package nvim

import (
	"os"
	"path"

	"github.com/pablobfonseca/dotfiles/src/config"
	"github.com/pablobfonseca/dotfiles/src/utils"
	"github.com/spf13/cobra"
	"github.com/vbauerster/mpb/v7"
)

var InstallNvimCmd = &cobra.Command{
	Use:   "nvim",
	Short: "Install nvim files",
	Run: func(cmd *cobra.Command, args []string) {
		p := mpb.New()

		verbose, _ := cmd.Flags().GetBool("verbose")

		if utils.CommandExists("nvim") {
			utils.SkipMessage("nvim already installed")
		} else {
			InstallNvimBar := utils.NewBar("Installing nvim", 1, p)
			if err := utils.ExecuteCommand(verbose, "brew", "install", "nvim"); err != nil {
				utils.ErrorMessage("Error installing nvim", err)
			}
			InstallNvimBar.Increment()
		}

		utils.CloneRepoIfNotExists(verbose)

		if utils.DirExists(config.NvimConfigDir()) {
			utils.SkipMessage("nvim files already exists")
			return
		}

		symlinkBar := utils.NewBar("Symlinking files", 1, p)

		src := path.Join(config.DotfilesConfigDir(), "nvim")
		dest := path.Join(config.NvimConfigDir())
		if err := os.Symlink(src, dest); err != nil {
			utils.ErrorMessage("Error creating symlink", err)
		}
		symlinkBar.Increment()

		utils.SuccessMessage("nvim files synced")
	},
}

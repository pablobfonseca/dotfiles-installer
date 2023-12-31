package vim

import (
	"os"
	"path"

	"github.com/pablobfonseca/dotfiles/src/config"
	"github.com/pablobfonseca/dotfiles/src/utils"
	"github.com/spf13/cobra"
	"github.com/vbauerster/mpb/v7"
)

const NVCHAD_REPO = "https://github.com/NvChad/NvChad"

var InstallVimCmd = &cobra.Command{
	Use:   "vim",
	Short: "Install vim files",
	Run: func(cmd *cobra.Command, args []string) {
		p := mpb.New()

		verbose, _ := cmd.Flags().GetBool("verbose")

		if utils.CommandExists("nvim") {
			utils.SkipMessage("nvim already installed")
		} else {
			installNvimBar := utils.NewBar("Installing nvim", 1, p)
			if err := utils.ExecuteCommand(verbose, "brew", "install", "neovim"); err != nil {
				utils.ErrorMessage("Error installing nvim", err)
			}
			installNvimBar.Increment()
		}

		utils.CloneRepoIfNotExists(verbose)

		if utils.DirExists(config.NvimConfigDir()) {
			utils.SkipMessage("nvim files already exists")
			return
		}
		installNvChadBar := utils.NewBar("Installing NvChad", 1, p)

		if err := utils.ExecuteCommand(verbose, "git", "clone", "--depth", "1", NVCHAD_REPO, config.NvimConfigDir()); err != nil {
			utils.ErrorMessage("Error cloning the repository", err)
		}
		installNvChadBar.Increment()

		symlinkBar := utils.NewBar("Symlinking files", 1, p)

		src := path.Join(config.DotfilesConfigDir(), "nvim", "custom")
		dest := path.Join(config.NvimConfigDir(), "lua", "custom")
		if err := os.Symlink(src, dest); err != nil {
			utils.ErrorMessage("Error creating symlink", err)
		}
		symlinkBar.Increment()
	},
}

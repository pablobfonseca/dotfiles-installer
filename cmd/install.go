package cmd

import (
	"github.com/spf13/cobra"

	"github.com/pablobfonseca/dotfiles/cmd/homebrew"
	"github.com/pablobfonseca/dotfiles/cmd/nvim"
	"github.com/pablobfonseca/dotfiles/cmd/zsh"
	"github.com/pablobfonseca/dotfiles/src/config"
	"github.com/pablobfonseca/dotfiles/src/installer"
	"github.com/pablobfonseca/dotfiles/src/utils"
)

var installCmd = &cobra.Command{
	Use:     "install",
	Short:   "Install the dotfiles",
	Example: "dotfiles install nvim",
	Long:    "Install the dotfiles. You can install all the dotfiles or just some of them.",
	Run: func(cmd *cobra.Command, args []string) {
		err := installer.SetupZsh()
		if err != nil {
			utils.ErrorMessage("[setup:zsh]:", err)
		}

		if utils.DirExists(config.DotfilesConfigDir()) {
			if utils.Confirm("Dotfiles already exists, do you want to check for updates?") {
				if err := utils.UpdateRepository(); err != nil {
					utils.ErrorMessage("Error updating repository", err)
				}
			}
		} else {
			utils.InfoMessage("Cloning dotfiles repository")
			if err := utils.ExecuteCommand("git", "clone", config.RepositoryUrl(), config.DotfilesConfigDir()); err != nil {
				utils.ErrorMessage("Error cloning the repository", err)
			}
		}

		err = installer.InstallHomebrew()
		if err != nil {
			utils.ErrorMessage("[homebrew]: %v", err)
		}

	},
}

func init() {
	rootCmd.AddCommand(installCmd)

	installCmd.AddCommand(nvim.InstallNvimCmd)
	installCmd.AddCommand(zsh.InstallZshCmd)
	installCmd.AddCommand(homebrew.InstallHomebrewCmd)
}

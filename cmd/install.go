package cmd

import (
	"github.com/spf13/cobra"

	"github.com/pablobfonseca/dotfiles/cmd/homebrew"
	"github.com/pablobfonseca/dotfiles/cmd/nvim"
	"github.com/pablobfonseca/dotfiles/cmd/zsh"
	"github.com/pablobfonseca/dotfiles/src/config"
	"github.com/pablobfonseca/dotfiles/src/utils"
)

var installCmd = &cobra.Command{
	Use:     "install",
	Short:   "Install the dotfiles",
	Example: "dotfiles install nvim",
	Long:    "Install the dotfiles. You can install all the dotfiles or just some of them.",
	Run: func(cmd *cobra.Command, args []string) {
		if utils.DirExists(config.DotfilesConfigDir()) {
			utils.SkipMessage("Dotfiles repo already exists")
		} else {
			if err := utils.ExecuteCommand("git", "clone", config.RepositoryUrl(), config.DotfilesConfigDir()); err != nil {
				utils.ErrorMessage("Error cloning the repository", err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(installCmd)

	installCmd.AddCommand(nvim.InstallNvimCmd)
	installCmd.AddCommand(zsh.InstallZshCmd)
	installCmd.AddCommand(homebrew.InstallHomebrewCmd)
}

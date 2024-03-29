package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vbauerster/mpb/v7"

	"github.com/pablobfonseca/dotfiles/cmd/emacs"
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
		p := mpb.New()

		if utils.DirExists(config.DotfilesConfigDir()) {
			utils.SkipMessage("Dotfiles repo already exists")
		} else {
			bar := utils.NewBar("Cloning dotfiles repo", 1, p)

			if err := utils.ExecuteCommand(verbose, "git", "clone", config.RepositoryUrl(), config.DotfilesConfigDir()); err != nil {
				utils.ErrorMessage("Error cloning the repository", err)
			}
			bar.Increment()

			p.Wait()
		}
	},
}

func init() {
	rootCmd.AddCommand(installCmd)

	installCmd.AddCommand(nvim.InstallNvimCmd)
	installCmd.AddCommand(emacs.InstallEmacsCmd)
	installCmd.AddCommand(zsh.InstallZshCmd)
	installCmd.AddCommand(homebrew.InstallHomebrewCmd)
}

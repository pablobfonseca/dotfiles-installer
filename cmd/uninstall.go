package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vbauerster/mpb/v7"

	"github.com/pablobfonseca/dotfiles/cmd/emacs"
	"github.com/pablobfonseca/dotfiles/cmd/vim"
	"github.com/pablobfonseca/dotfiles/src/config"
	"github.com/pablobfonseca/dotfiles/src/utils"
)

var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstall the dotfiles",
	Long:  "Uninstall the dotfiles. You can uninstall all the dotfiles or just some of them.",
	Run: func(cmd *cobra.Command, args []string) {
		p := mpb.New()

		bar := utils.NewBar("Deleting dotfiles repo", 1, p)

		if err := utils.ExecuteCommand(verbose, "rm", "-rf", config.DotfilesConfigDir()); err != nil {
			utils.ErrorMessage("Error deleting the repository", err)
		}
		bar.Increment()

		p.Wait()
	},
}

var uninstallApp bool

func init() {
	rootCmd.AddCommand(uninstallCmd)

	uninstallCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Verbose output")

	uninstallCmd.AddCommand(vim.UninstallVimCmd)
	uninstallCmd.AddCommand(emacs.UninstallEmacsCmd)
}

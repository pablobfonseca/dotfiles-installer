/*
Copyright © 2023 Pablo Fonseca <pablofonseca777@gmail.com>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vbauerster/mpb/v7"

	dotfiles "github.com/pablobfonseca/dotfiles/src/installers"
)

var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstall the dotfiles",
	Long: `Uninstall the dotfiles. You can uninstall all the dotfiles or just some of them.
    Example: dotfiles uninstall --nvim
             dotfiles uninstall --emacs
  `,
	Run: func(cmd *cobra.Command, args []string) {
		p := mpb.New()

		all, _ := cmd.Flags().GetBool("all")
		nvim, _ := cmd.Flags().GetBool("nvim")
		emacs, _ := cmd.Flags().GetBool("emacs")

		if all {
			uninstallAll(p)
		}
		if nvim {
			dotfiles.UninstallNvim(uninstallApp, p, verbose)
		}
		if emacs {
			dotfiles.UninstallEmacs(uninstallApp, p, verbose)
		}

		p.Wait()
	},
}

func uninstallAll(p *mpb.Progress) {
	dotfiles.DeleteRepo(p, verbose)
	dotfiles.UninstallNvim(uninstallApp, p, verbose)
	dotfiles.UninstallEmacs(uninstallApp, p, verbose)
}

var uninstallApp bool

func init() {
	rootCmd.AddCommand(uninstallCmd)

	uninstallCmd.Flags().BoolVarP(&uninstallApp, "prune", "p", false, "Also uninstall the app")

	uninstallCmd.Flags().BoolP("all", "a", false, "Uninstall all dotfiles")
	uninstallCmd.Flags().BoolP("nvim", "n", false, "Uninstall nvim files")
	uninstallCmd.Flags().BoolP("emacs", "e", false, "Uninstall emacs files")
}

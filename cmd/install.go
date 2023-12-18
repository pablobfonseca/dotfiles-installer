/*
Copyright © 2023 Pablo Fonseca <pablofonseca777@gmail.com>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vbauerster/mpb/v7"

	dotfiles "github.com/pablobfonseca/dotfiles/src/installers"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install the dotfiles",
	Long: `Install the dotfiles. You can install all the dotfiles or just some of them.
    Example: dotfiles install --all
             dotfiles install --nvim
  `,
	Run: func(cmd *cobra.Command, args []string) {
		p := mpb.New()

		all, _ := cmd.Flags().GetBool("all")
		nvim, _ := cmd.Flags().GetBool("nvim")
		emacs, _ := cmd.Flags().GetBool("emacs")
		zsh, _ := cmd.Flags().GetBool("zsh")

		if all {
			installAll(p)
		}
		if nvim {
			dotfiles.InstallNvim(p, verbose)
		}
		if emacs {
			dotfiles.InstallEmacs(p, verbose)
		}
		if zsh {
			dotfiles.InstallZsh(p)
		}

		p.Wait()
	},
}

func installAll(p *mpb.Progress) {
	dotfiles.CloneRepo(p, verbose)
	dotfiles.InstallHomebrew(p, verbose)
	dotfiles.InstallNvim(p, verbose)
	dotfiles.InstallZsh(p)
	dotfiles.InstallEmacs(p, verbose)
}

func init() {
	rootCmd.AddCommand(installCmd)

	installCmd.Flags().BoolP("all", "a", false, "Install all the dotfiles")
	installCmd.Flags().BoolP("nvim", "n", false, "Install nvim files")
	installCmd.Flags().BoolP("emacs", "e", false, "Install emacs files")
	installCmd.Flags().BoolP("zsh", "z", false, "Install zsh files")
}

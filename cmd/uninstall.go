/*
Copyright © 2023 Pablo Fonseca <pablofonseca777@gmail.com>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vbauerster/mpb/v7"

	dotfiles "github.com/pablobfonseca/dotfiles-cli/src/installers"
)

var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstall the dotfiles",
	Long: `Uninstall the dotfiles. You can uninstall all the dotfiles or just some of them.
    Example: dotfiles uninstall --all
             dotfiles uninstall --nvim
  `,
	Run: func(cmd *cobra.Command, args []string) {
		p := mpb.New()

		// all, _ := cmd.Flags().GetBool("all")
		nvim, _ := cmd.Flags().GetBool("nvim")
		// emacs, _ := cmd.Flags().GetBool("emacs")
		// zsh, _ := cmd.Flags().GetBool("zsh")

		// if all {
		// 	installAll(p)
		// }
		if nvim {
			dotfiles.UninstallNvim(p)
		}
		// if emacs {
		// 	dotfiles.InstallEmacs(p)
		// }
		// if zsh {
		// 	dotfiles.InstallZsh(p)
		// }

		p.Wait()
	},
}

// func installAll(p *mpb.Progress) {
// 	dotfiles.CloneRepo(p)
// 	dotfiles.InstallHomebrew(p)
// 	dotfiles.InstallNvim(p)
// 	dotfiles.InstallZsh(p)
// 	dotfiles.InstallEmacs(p)
// }

func init() {
	rootCmd.AddCommand(uninstallCmd)

	// uninstallCmd.Flags().BoolP("all", "a", false, "Install all the dotfiles")
	uninstallCmd.Flags().BoolP("nvim", "n", false, "Uninstall nvim files")
	// uninstallCmd.Flags().BoolP("emacs", "e", false, "Install emacs files")
	// uninstallCmd.Flags().BoolP("zsh", "z", false, "Install zsh files")
}

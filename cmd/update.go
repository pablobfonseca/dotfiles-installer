package cmd

import (
	updaters "github.com/pablobfonseca/dotfiles-cli/src/updaters"
	"github.com/spf13/cobra"
	"github.com/vbauerster/mpb/v7"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update the dotfiles",
	Long: `Update the dotfiles. You can update all the dotfiles or just some of them.
    Example: dotfiles update --all
              dotfiles update --nvim
              dotfiles update --brew
  `,
	Run: func(cmd *cobra.Command, args []string) {
		p := mpb.New()

		nvim, _ := cmd.Flags().GetBool("nvim")
		brew, _ := cmd.Flags().GetBool("brew")

		if nvim {
			updaters.UpdateNvim(p)
		}

		if brew {
			updaters.UpdateBrew(p)
		}

		p.Wait()
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	updateCmd.Flags().BoolP("nvim", "n", false, "Update nvim files")
	updateCmd.Flags().BoolP("brew", "b", false, "Update homebrew")
}

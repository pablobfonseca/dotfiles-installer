package vim

import (
	"fmt"

	"github.com/spf13/cobra"
)

var VimRootCmd = &cobra.Command{
	Use:   "vim",
	Short: "vim is a text editor",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("vim called")
	},
}

func init() {
	VimRootCmd.AddCommand(vimInstallCmd)
}

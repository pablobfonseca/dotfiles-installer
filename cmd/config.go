package cmd

import (
	"fmt"

	"github.com/enescakir/emoji"
	"github.com/pablobfonseca/dotfiles/src/config"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Show configuration",
	Long:  `Display current dotfiles configuration settings.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%v Dotfiles Configuration\n\n", emoji.Gear)
		fmt.Printf("  Repository URL: %s\n", config.RepositoryUrl())
		fmt.Printf("  Directory: %s\n", config.DotfilesConfigDir())
		
		// Validate current config
		if err := config.ValidateConfig(); err != nil {
			fmt.Printf("\n%v Configuration Error: %v\n", emoji.Warning, err)
		} else {
			fmt.Printf("\n%v Configuration is valid\n", emoji.CheckMark)
		}
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
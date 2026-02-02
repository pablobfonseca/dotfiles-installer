package cmd

import (
	"github.com/pablobfonseca/dotfiles/src/utils/prompts"
	"github.com/spf13/cobra"
)

var dashboardCmd = &cobra.Command{
	Use:   "dashboard",
	Short: "Launch the installation dashboard",
	Long:  "Launch an interactive dashboard to view and manage your dotfiles installation",
	Run: func(cmd *cobra.Command, args []string) {
		// Setup repository first
		setupRepository()
		
		// Launch dashboard TUI
		prompts.LaunchDashboard()
	},
}

func init() {
	rootCmd.AddCommand(dashboardCmd)
}
package cmd

import (
	"github.com/pablobfonseca/dotfiles/src/utils"
	"github.com/spf13/cobra"
)

var backupCmd = &cobra.Command{
	Use:     "backup",
	Short:   "Backup the dotfiles",
	Example: "dotfiles backup --resource nvim",
	Long:    "Backup the dotfiles. You can backup all the dotfiles or just some of them.",
	Run: func(cmd *cobra.Command, args []string) {
		resource, err := cmd.Flags().GetString("resource")
		if err != nil {
			utils.ErrorMessage("Error getting resource flag", err)
			return
		}

		if resource == "" {
			utils.InfoMessage("No resource specified, backing up all dotfiles")
			// TODO: Implement backup all functionality
			return
		}

		utils.InfoMessage("Backing up", resource)
		// TODO: Implement backup functionality for specific resource
	},
}

func init() {
	rootCmd.AddCommand(backupCmd)

	backupCmd.Flags().StringP("resource", "r", "", "Backup a specific resource (e.g: nvim, zsh)")
}

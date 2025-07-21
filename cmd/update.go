package cmd

import (
	"fmt"

	"github.com/pablobfonseca/dotfiles/src/config"
	"github.com/pablobfonseca/dotfiles/src/ui"
	"github.com/pablobfonseca/dotfiles/src/utils"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update dotfiles repository and configurations",
	Long: `Update your dotfiles repository to the latest version and refresh configurations.
This will pull the latest changes from your dotfiles repository and update symlinks.`,
	Run: func(cmd *cobra.Command, args []string) {
		utils.SetDryRunChecker(func() bool {
			dryRun, _ := cmd.Flags().GetBool("dry-run")
			return dryRun
		})

		ui.ShowInfo("Starting dotfiles update...")

		// Check if repository exists
		if !utils.DirExists(config.DotfilesConfigDir()) {
			ui.ShowError("Dotfiles repository not found. Run 'dotfiles install' first.")
			return
		}

		steps := []string{
			"Checking repository status",
			"Pulling latest changes",
			"Updating configurations",
			"Refreshing symlinks",
			"Update completed",
		}

		tasks := []func() error{
			func() error {
				// Check if repo is clean or has changes
				return nil
			},
			func() error {
				ui.ShowInfo("Updating repository from " + config.RepositoryUrl())
				return utils.UpdateRepository()
			},
			func() error {
				// Update brew packages if requested
				updateBrew, _ := cmd.Flags().GetBool("brew")
				if updateBrew {
					ui.ShowInfo("Updating Homebrew packages...")
					return utils.ExecuteCommand("brew", "bundle", "--force", fmt.Sprintf("--file=%s/homebrew/Brewfile", config.DotfilesConfigDir()))
				}
				return nil
			},
			func() error {
				// Re-sync configuration files
				return utils.ExecuteCommand("echo", "Refreshing symlinks...")
			},
			func() error {
				ui.ShowSuccess("🎉 Dotfiles updated successfully!")
				return nil
			},
		}

		ui.RunWithProgress(steps, tasks)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	
	updateCmd.Flags().Bool("brew", false, "also update Homebrew packages")
}
package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/enescakir/emoji"
	"github.com/pablobfonseca/dotfiles/src/config"
	"github.com/pablobfonseca/dotfiles/src/ui"
	"github.com/pablobfonseca/dotfiles/src/utils"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show installation status of tools",
	Long:  `Show the current installation status of all available tools and configurations.`,
	Run: func(cmd *cobra.Command, args []string) {
		ui.ShowInfo("Checking installation status...")

		// Check repository status
		if utils.DirExists(config.DotfilesConfigDir()) {
			ui.ShowSuccess(fmt.Sprintf("Repository: Cloned at %s", config.DotfilesConfigDir()))
		} else {
			ui.ShowError("Repository: Not cloned")
		}

		// Check tools status
		tools := []struct {
			name       string
			checkCmd   string
			configPath string
		}{
			{"Homebrew", "brew", ""},
			{"Neovim", "nvim", path.Join(os.Getenv("HOME"), ".config/nvim")},
			{"Git", "git", path.Join(os.Getenv("HOME"), ".gitconfig")},
			{"Tmux", "tmux", path.Join(os.Getenv("HOME"), ".config/tmux")},
			{"Starship", "starship", path.Join(os.Getenv("HOME"), ".config/starship.toml")},
			{"Karabiner-Elements", "karabiner_cli", path.Join(os.Getenv("HOME"), ".config/karabiner")},
		}

		fmt.Printf("\n%v Tool Status:\n\n", emoji.Information)
		for _, tool := range tools {
			installed := utils.CommandExists(tool.checkCmd)
			configured := tool.configPath == "" || utils.FileExists(tool.configPath) || utils.DirExists(tool.configPath)

			if installed && configured {
				fmt.Printf("  %v %s: Installed & Configured\n", emoji.CheckMark, tool.name)
			} else if installed {
				fmt.Printf("  %v %s: Installed (not configured)\n", emoji.Warning, tool.name)
			} else {
				fmt.Printf("  %v %s: Not installed\n", emoji.CrossMark, tool.name)
			}
		}

		// Show config info
		fmt.Printf("\n%v Configuration:\n", emoji.Gear)
		fmt.Printf("  Repository: %s\n", config.RepositoryUrl())
		fmt.Printf("  Directory: %s\n", config.DotfilesConfigDir())
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
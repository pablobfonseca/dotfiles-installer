package cmd

import (
	"github.com/spf13/cobra"

	"github.com/pablobfonseca/dotfiles/cmd/homebrew"
	"github.com/pablobfonseca/dotfiles/cmd/nvim"
	"github.com/pablobfonseca/dotfiles/cmd/zsh"
	"github.com/pablobfonseca/dotfiles/src/config"
	"github.com/pablobfonseca/dotfiles/src/installer"
	"github.com/pablobfonseca/dotfiles/src/ui"
	"github.com/pablobfonseca/dotfiles/src/utils"
)

var installCmd = &cobra.Command{
	Use:     "install",
	Short:   "Install the dotfiles with beautiful TUI",
	Example: "dotfiles install --interactive",
	Long:    "Install the dotfiles with an interactive terminal UI. You can install all tools or select specific ones.",
	Run: func(cmd *cobra.Command, args []string) {
		utils.SetDryRunChecker(func() bool {
			dryRun, _ := cmd.Flags().GetBool("dry-run")
			return dryRun
		})

		// Show banner
		ui.ShowBanner()
		ui.ShowWelcome()

		interactive, _ := cmd.Flags().GetBool("interactive")
		
		if interactive {
			// Interactive mode with tool selection
			tools := []ui.Tool{
				{Name: "Homebrew", Desc: "Package manager for macOS", Installed: utils.CommandExists("brew")},
				{Name: "Neovim", Desc: "Modern Vim-based text editor with config", Installed: utils.CommandExists("nvim")},
				{Name: "Zsh", Desc: "Z shell configuration files", Installed: utils.FileExists("/Users/" + utils.GetCurrentUser() + "/.zshrc")},
				{Name: "Wezterm", Desc: "GPU-accelerated terminal emulator", Installed: utils.CommandExists("wezterm")},
				{Name: "Tmux", Desc: "Terminal multiplexer configuration", Installed: utils.CommandExists("tmux")},
				{Name: "Starship", Desc: "Cross-shell prompt configuration", Installed: utils.CommandExists("starship")},
				{Name: "Git Config", Desc: "Git configuration files", Installed: utils.FileExists("/Users/" + utils.GetCurrentUser() + "/.gitconfig")},
			}

			selectedTools, err := ui.RunToolSelector(tools)
			if err != nil {
				ui.ShowError("Tool selection failed: " + err.Error())
				return
			}

			if len(selectedTools) == 0 {
				ui.ShowInfo("No tools selected. Exiting.")
				return
			}

			// Run installation with progress
			runInteractiveInstall(selectedTools)
		} else {
			// Standard installation with progress
			runStandardInstall()
		}
	},
}

func runStandardInstall() {
	steps := []string{
		"Setting up repository",
		"Installing Homebrew",
		"Setting up Zsh configuration", 
		"Installing configuration files",
		"Finalizing installation",
	}

	tasks := []func() error{
		func() error {
			if utils.DirExists(config.DotfilesConfigDir()) {
				return utils.UpdateRepository()
			} else {
				return utils.ExecuteCommand("git", "clone", config.RepositoryUrl(), config.DotfilesConfigDir())
			}
		},
		func() error {
			return installer.InstallHomebrew()
		},
		func() error {
			return installer.SetupZsh()
		},
		func() error {
			return installer.InstallConfigFiles()
		},
		func() error {
			ui.ShowSuccess("🎉 Installation completed successfully!")
			return nil
		},
	}

	ui.RunWithProgress(steps, tasks)
}

func runInteractiveInstall(selectedTools []ui.Tool) {
	steps := make([]string, 0, len(selectedTools)+2)
	tasks := make([]func() error, 0, len(selectedTools)+2)

	// Always setup repository first
	steps = append(steps, "Setting up repository")
	tasks = append(tasks, func() error {
		if utils.DirExists(config.DotfilesConfigDir()) {
			return utils.UpdateRepository()
		} else {
			return utils.ExecuteCommand("git", "clone", config.RepositoryUrl(), config.DotfilesConfigDir())
		}
	})

	// Add selected tools
	for _, tool := range selectedTools {
		steps = append(steps, "Installing "+tool.Name)
		
		switch tool.Name {
		case "Homebrew":
			tasks = append(tasks, func() error { return installer.InstallHomebrew() })
		case "Neovim":
			tasks = append(tasks, func() error { return installer.InstallNvim() })
		case "Zsh":
			tasks = append(tasks, func() error { return installer.SetupZsh() })
		case "Wezterm":
			tasks = append(tasks, func() error { return installer.SetupWezterm() })
		case "Tmux":
			tasks = append(tasks, func() error { return installer.SetupTmux() })
		case "Starship":
			tasks = append(tasks, func() error { return installer.SetupStarship() })
		case "Git Config":
			tasks = append(tasks, func() error { return installer.SetupGit() })
		default:
			tasks = append(tasks, func() error { return nil })
		}
	}

	// Finalize
	steps = append(steps, "Finalizing installation")
	tasks = append(tasks, func() error {
		ui.ShowSuccess("🎉 Selected tools installed successfully!")
		return nil
	})

	ui.RunWithProgress(steps, tasks)
}

func init() {
	rootCmd.AddCommand(installCmd)
	
	installCmd.Flags().BoolP("interactive", "i", false, "run interactive installation with tool selection")

	installCmd.AddCommand(nvim.InstallNvimCmd)
	installCmd.AddCommand(zsh.InstallZshCmd)
	installCmd.AddCommand(homebrew.InstallHomebrewCmd)
}

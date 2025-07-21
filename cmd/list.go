package cmd

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/enescakir/emoji"
	"github.com/pablobfonseca/dotfiles/src/ui"
	"github.com/pablobfonseca/dotfiles/src/utils"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available tools to install",
	Long:  `List all available tools that can be installed with this dotfiles manager.`,
	Run: func(cmd *cobra.Command, args []string) {
		ui.ShowInfo("Available tools to install")

		tools := []struct {
			name        string
			description string
			command     string
			installed   bool
		}{
			{"Homebrew", "Package manager for macOS", "dotfiles install homebrew", utils.CommandExists("brew")},
			{"Neovim", "Modern Vim-based text editor with config", "dotfiles install nvim", utils.CommandExists("nvim")},
			{"Zsh", "Z shell configuration files", "dotfiles install zsh", utils.FileExists("/Users/"+utils.GetCurrentUser()+"/.zshrc")},
			{"Wezterm", "GPU-accelerated terminal emulator", "dotfiles install --interactive", utils.CommandExists("wezterm")},
			{"Tmux", "Terminal multiplexer configuration", "dotfiles install --interactive", utils.CommandExists("tmux")},
			{"Starship", "Cross-shell prompt configuration", "dotfiles install --interactive", utils.CommandExists("starship")},
			{"Karabiner-Elements", "Keyboard customization tool", "dotfiles install karabiner", utils.CommandExists("karabiner_cli")},
			{"Aerospace", "Window manager for macOS", "dotfiles install --interactive", utils.CommandExists("aerospace")},
			{"Git Config", "Git configuration files", "dotfiles install --interactive", utils.FileExists("/Users/"+utils.GetCurrentUser()+"/.gitconfig")},
		}

		// Style definitions
		toolStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("36")).Bold(true)
		descStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Italic(true)
		cmdStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("220")).Bold(true)
		installedStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("34"))
		notInstalledStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("196"))

		fmt.Println()
		for _, tool := range tools {
			status := "Not Installed"
			statusStyle := notInstalledStyle
			
			if tool.installed {
				status = "Installed"
				statusStyle = installedStyle
			}

			fmt.Printf("  %v %s %s\n", 
				emoji.Gear, 
				toolStyle.Render(tool.name),
				statusStyle.Render(fmt.Sprintf("[%s]", status)),
			)
			fmt.Printf("    %s\n", descStyle.Render(tool.description))
			fmt.Printf("    %s %s\n\n", 
				emoji.RightArrow,
				cmdStyle.Render(tool.command),
			)
		}

		// Usage tips
		tipStyle := lipgloss.NewStyle().
			Foreground(lipgloss.Color("220")).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("220")).
			Padding(1, 2).
			MarginTop(1)

		tips := fmt.Sprintf(
			"%s Quick Commands:\n\n"+
				"• %s Install everything interactively\n"+
				"• %s Preview installation\n"+
				"• %s Check what's installed\n"+
				"• %s View configuration",
			emoji.Information,
			cmdStyle.Render("dotfiles install --interactive"),
			cmdStyle.Render("dotfiles install --dry-run"),
			cmdStyle.Render("dotfiles status"),
			cmdStyle.Render("dotfiles config"),
		)

		fmt.Println(tipStyle.Render(tips))
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
package cmd

import (
	"os"
	"path"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/pablobfonseca/dotfiles/src/config"
	"github.com/pablobfonseca/dotfiles/src/utils"
	"github.com/pablobfonseca/dotfiles/src/utils/prompts"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "dotfiles",
	Short: "🚀 Personal dotfiles installer",
	Long: `🚀 Dotfiles Installer

A CLI tool to manage your personal dotfiles and development environment setup.
Install your favorite tools, configurations, and settings with an interactive terminal UI.

Features:
• 🔧 Interactive tool selection
• 📊 Installation status tracking  
• 🏃 Dry-run mode for safe testing
• ⚙️  Configuration management
• 🔄 Update and rollback capabilities

Examples:
  dotfiles install --interactive    # Interactive installation with tool selection
  dotfiles install --dry-run       # Preview what would be installed
  dotfiles status                   # Check installation status
  dotfiles list                     # List available tools
  dotfiles config                   # Show current configuration`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func IsDryRun() bool {
	return dryRun
}

var cfgFile = ""
var dryRun = false

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/dotfiles/config.toml)")
	rootCmd.PersistentFlags().BoolVarP(&dryRun, "dry-run", "n", false, "show what would be done without executing")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		if err != nil {
			utils.ErrorMessage("Something went wrong", err)
		}

		configDir := path.Join(home, ".config/", "dotfiles/")

		if _, err := os.Stat(configDir); os.IsNotExist(err) {
			err := os.MkdirAll(configDir, 0755)
			if err != nil {
				utils.ErrorMessage("Error creating the config dir", err)
			}
		}

		viper.AddConfigPath(configDir)
		viper.SetConfigType("toml")
		viper.SetConfigName("config")
	}

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			p := tea.NewProgram(prompts.ConfigPrompt())

			if _, err := p.Run(); err != nil {
				utils.ErrorMessage("[Config prompt error]: Something went wrong", err)
			}

		} else {
			utils.ErrorMessage("Something went wrong", err)
		}
	} else {
		// Validate config after reading
		if err := config.ValidateConfig(); err != nil {
			utils.ErrorMessage("Invalid configuration", err)
		}
	}
}

package cmd

import (
	"os"
	"path"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/pablobfonseca/dotfiles/src/utils"
	"github.com/pablobfonseca/dotfiles/src/utils/prompts"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "dotfiles",
	Short: "Install dotfiles from a git repository",
	Long:  `dotfiles is a CLI tool to install dotfiles from a git repository.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var cfgFile = ""

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/dotfiles/config.toml)")
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

		if _, err := os.Stat(path.Join(home, ".config/", "dotfiles/")); os.IsNotExist(err) {
			err := os.Mkdir(configDir, 0755)
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
	}
}

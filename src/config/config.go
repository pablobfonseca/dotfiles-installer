package config

import "github.com/spf13/viper"

func RepositoryUrl() string {
	return "https://github.com/" + viper.GetString("dotfiles.repository")
}

func DotfilesConfigDir() string {
	return viper.GetString("dotfiles.default_dir")
}

func NvimConfigDir() string {
	return viper.GetString("dotfiles.nvim.config_dir")
}

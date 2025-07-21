package config

import (
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/spf13/viper"
)

func RepositoryUrl() string {
	return "https://github.com/" + viper.GetString("dotfiles.repository")
}

func DotfilesConfigDir() string {
	return viper.GetString("dotfiles.default_dir")
}

func ValidateConfig() error {
	repo := viper.GetString("dotfiles.repository")
	if repo == "" {
		return fmt.Errorf("repository is required")
	}

	// Basic validation for repository format (should be username/repo)
	parts := strings.Split(repo, "/")
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		return fmt.Errorf("repository should be in format 'username/repository'")
	}

	// Validate URL format
	repoURL := RepositoryUrl()
	_, err := url.Parse(repoURL)
	if err != nil {
		return fmt.Errorf("invalid repository URL: %v", err)
	}

	// Validate directory path
	dir := DotfilesConfigDir()
	if dir == "" {
		return fmt.Errorf("dotfiles directory is required")
	}

	// Expand path if it starts with ~
	if strings.HasPrefix(dir, "~") {
		home, err := os.UserHomeDir()
		if err != nil {
			return fmt.Errorf("cannot get home directory: %v", err)
		}
		dir = strings.Replace(dir, "~", home, 1)
		viper.Set("dotfiles.default_dir", dir)
	}

	return nil
}

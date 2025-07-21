package installer

import (
	"fmt"
	"os"
	"path"

	"github.com/pablobfonseca/dotfiles/src/config"
	"github.com/pablobfonseca/dotfiles/src/utils"
)

const homebrewInstallUrl = "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"

func InstallHomebrew() error {
	if utils.CommandExists("brew") {
		utils.SkipMessage("Homebrew is already installed")
		if utils.Confirm("Do you want to install/reinstall packages from Brewfile?") {
			return installBrewPackages()
		}
	} else {
		err := utils.ExecuteCommand("/bin/bash", "-c", homebrewInstallUrl)
		if err != nil {
			return err
		}
		if err := installBrewPackages(); err != nil {
			return err
		}
	}

	return nil
}

func installBrewPackages() error {
	utils.InfoMessage("Installing brew packages...")
	if err := utils.ExecuteCommand("brew", "bundle", fmt.Sprintf("--file=%s", path.Join(config.DotfilesConfigDir(), "/homebrew/Brewfile"))); err != nil {
		return err
	}

	return nil
}

func InstallNvim() error {
	if utils.CommandExists("nvim") {
		return nil
	}

	return utils.ExecuteCommand("brew", "install", "nvim")
}

func SetupStarship() error {
	configDir, _ := os.UserConfigDir()

	src := path.Join(config.DotfilesConfigDir(), "starship", "starship.toml")
	dest := path.Join(configDir, "starship.toml")

	err := install(src, dest)
	if err != nil {
		return err
	}

	return nil
}

func SetupAerospace() error {
	configDir, _ := os.UserConfigDir()

	src := path.Join(config.DotfilesConfigDir(), "aerospace", "aerospace.toml")
	dest := path.Join(configDir, "aerospace.toml")

	err := install(src, dest)
	if err != nil {
		return err
	}

	return nil
}

func SetupGit() error {
	home, _ := os.UserHomeDir()

	for _, file := range []string{"gitconfig", "gitignore", "global_ignore", "gitmessage"} {
		src := path.Join(config.DotfilesConfigDir(), "zsh", file)
		dest := path.Join(home, "."+file)

		err := install(src, dest)
		if err != nil {
			return err
		}
	}

	return nil
}

func SetupZsh() error {
	home, _ := os.UserHomeDir()

	for _, file := range []string{"zshrc", "zprofile", "zlogin"} {
		src := path.Join(config.DotfilesConfigDir(), "zsh", file)
		dest := path.Join(home, "."+file)

		err := install(src, dest)
		if err != nil {
			return err
		}
	}

	// Source the .zshrc file to apply changes immediately
	zshrcPath := path.Join(home, ".zshrc")
	if utils.FileExists(zshrcPath) {
		utils.InfoMessage("Sourcing .zshrc to apply changes...")
		
		// Try to source the file - use a safer approach with exec
		sourceCmd := fmt.Sprintf(". %s", zshrcPath)
		if err := utils.ExecuteCommand("/bin/zsh", "-c", sourceCmd); err != nil {
			// Don't fail the installation if sourcing fails, just provide helpful info
			utils.InfoMessage("Note: Please restart your terminal or run 'source ~/.zshrc' to apply changes")
		} else {
			utils.SuccessMessage("✅ Zsh configuration loaded successfully!")
		}
	}

	return nil
}

func SetupWezterm() error {
	configDir, _ := os.UserConfigDir()

	src := path.Join(config.DotfilesConfigDir(), "wezterm")
	dest := path.Join(configDir, "wezterm")

	return install(src, dest)
}

func SetupTmux() error {
	configDir, _ := os.UserConfigDir()

	src := path.Join(config.DotfilesConfigDir(), "tmux")
	dest := path.Join(configDir, "tmux")

	err := install(src, dest)

	utils.CloneRepoIfNotExists("https://github.com/tmux-plugins/tmp", path.Join(dest, "plugins", "tmp"))

	return err
}

func InstallKarabiner() error {
	if !utils.CommandExists("karabiner_cli") {
		utils.InfoMessage("Installing Karabiner-Elements...")
		if err := utils.ExecuteCommand("brew", "install", "--cask", "karabiner-elements"); err != nil {
			return err
		}
	} else {
		utils.SkipMessage("Karabiner-Elements is already installed")
	}
	
	return nil
}

func SetupKarabiner() error {
	configDir, _ := os.UserConfigDir()
	
	// Karabiner config is typically in ~/.config/karabiner/
	src := path.Join(config.DotfilesConfigDir(), "karabiner")
	dest := path.Join(configDir, "karabiner")
	
	err := install(src, dest)
	if err != nil {
		return err
	}
	
	// Restart Karabiner-Elements to load new configuration
	utils.InfoMessage("Restarting Karabiner-Elements to load new configuration...")
	if err := utils.ExecuteCommand("launchctl", "kickstart", "-k", "gui/$(id -u)/org.pqrs.karabiner.karabiner_console_user_server"); err != nil {
		utils.InfoMessage("Note: Please restart Karabiner-Elements manually to load the new configuration")
	}
	
	return nil
}

func InstallConfigFiles() error {
	err := SetupWezterm()
	if err != nil {
		return err
	}
	err = SetupStarship()
	if err != nil {
		return err
	}
	err = SetupTmux()
	if err != nil {
		return err
	}
	err = SetupAerospace()
	if err != nil {
		return err
	}
	err = SetupGit()
	if err != nil {
		return err
	}
	err = InstallKarabiner()
	if err != nil {
		return err
	}
	err = SetupKarabiner()
	if err != nil {
		return err
	}

	return err
}

func install(src, dest string) error {
	utils.InfoMessage("Syncing " + src + " to " + dest)
	if _, err := os.Stat(dest); err == nil {
		if utils.ConfirmDestructive(fmt.Sprintf("File %s already exists, do you want to replace it?", dest)) {
			// Remove existing file or directory
			if err := utils.RemoveAllFiles(dest); err != nil {
				utils.ErrorMessage(fmt.Sprintf("[%s]: error removing existing file/directory", src), err)
				return err
			}
			if err := utils.SymlinkFiles(src, dest); err != nil {
				return err
			}
		} else {
			utils.SkipMessage("File already exists: " + dest)
		}
	} else {
		if err := utils.SymlinkFiles(src, dest); err != nil {
			return err
		}
	}

	return nil
}

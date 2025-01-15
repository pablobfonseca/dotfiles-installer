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

func SetupZsh() error {
	home, _ := os.UserHomeDir()

	for _, file := range []string{"zshrc", "zprofile", "zlogin"} {
		src := path.Join(config.DotfilesConfigDir(), "zsh", file)
		dest := path.Join(home, "."+file)

		utils.InfoMessage("Syncing " + src + " to " + dest)
		if _, err := os.Stat(dest); err == nil {
			if utils.Confirm(fmt.Sprintf("File %s already exists, do you want to replace it?", dest)) {
				if err := os.Remove(dest); err != nil {
					utils.ErrorMessage("[zsh]: error symlinking file", err)
				}
				utils.SymlinkFiles(src, dest)
				continue
			} else {
				utils.SkipMessage("File already exists: " + dest)
				continue
			}
		} else {
			if err := utils.SymlinkFiles(src, dest); err != nil {
				return err
			}
		}
	}

	return nil
}

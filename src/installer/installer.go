package installer

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"

	"github.com/pablobfonseca/dotfiles/src/config"
	"github.com/pablobfonseca/dotfiles/src/utils"
)

const homebrewInstallUrl = "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"

type Application struct {
	name          string
	src           string
	dest          string
	multipleFiles bool
	files         []string
}

func NewApp(name, src, dest string, multipleFiles bool, files []string) *Application {
	return &Application{
		name:          name,
		src:           src,
		dest:          dest,
		multipleFiles: multipleFiles,
		files:         files,
	}
}

func (a *Application) Install() error {
	if a.multipleFiles {
		for _, file := range a.files {
			src := filepath.Join(config.DotfilesConfigDir(), a.name, file)

			err := install(src, a.dest)
			if err != nil {
				return err
			}
		}
		return nil
	}

	return install(a.src, a.dest)
}

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
	if err := utils.ExecuteCommand("brew", "bundle", fmt.Sprintf("--file=%s", filepath.Join(config.DotfilesConfigDir(), "/homebrew/Brewfile"))); err != nil {
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
	configDir, err := config.ConfigDir()
	if err != nil {
		return fmt.Errorf("failed to get config directory: %w", err)
	}

	src := filepath.Join(config.DotfilesConfigDir(), "starship", "starship.toml")
	dest := filepath.Join(configDir, "starship.toml")

	return install(src, dest)
}

func SetupGit() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get home directory: %w", err)
	}

	for _, file := range []string{"gitconfig", "gitignore", "global_ignore", "gitmessage"} {
		src := filepath.Join(config.DotfilesConfigDir(), "git", file)
		dest := filepath.Join(home, "."+file)

		err := install(src, dest)
		if err != nil {
			return err
		}
	}

	return nil
}

func SetupZsh() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get home directory: %w", err)
	}

	for _, file := range []string{"zshrc", "zprofile", "zlogin"} {
		src := filepath.Join(config.DotfilesConfigDir(), "zsh", file)
		dest := filepath.Join(home, "."+file)

		err := install(src, dest)
		if err != nil {
			return err
		}
	}

	// Source the .zshrc file to apply changes immediately
	zshrcPath := filepath.Join(home, ".zshrc")
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

func SetupGhostty() error {
	configDir, err := config.ConfigDir()
	if err != nil {
		return fmt.Errorf("failed to get config directory: %w", err)
	}

	src := filepath.Join(config.DotfilesConfigDir(), "ghostty")
	dest := filepath.Join(configDir, "ghostty")

	return install(src, dest)
}

func InstallCyberpunkTheme() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get home directory: %w", err)
	}

	dest := filepath.Join(home, ".cyberpunk-theme")
	utils.CloneRepoIfNotExists("https://github.com/pablobfonseca/cyberpunk-theme", dest)

	installScript := filepath.Join(dest, "install.sh")
	if !utils.FileExists(installScript) {
		return fmt.Errorf("install script not found at %s", installScript)
	}

	return utils.ExecuteCommand("/bin/bash", installScript)
}

func SetupTmux() error {
	configDir, err := config.ConfigDir()
	if err != nil {
		return fmt.Errorf("failed to get config directory: %w", err)
	}

	src := filepath.Join(config.DotfilesConfigDir(), "tmux")
	dest := filepath.Join(configDir, "tmux")

	err = install(src, dest)

	utils.CloneRepoIfNotExists("https://github.com/tmux-plugins/tmp", filepath.Join(dest, "plugins", "tmp"))

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
	configDir, err := config.ConfigDir()
	if err != nil {
		return fmt.Errorf("failed to get config directory: %w", err)
	}

	// Karabiner config is typically in ~/.config/karabiner/
	src := filepath.Join(config.DotfilesConfigDir(), "karabiner")
	dest := filepath.Join(configDir, "karabiner")

	if err := install(src, dest); err != nil {
		return err
	}

	// Restart Karabiner-Elements to load new configuration.
	// Resolve UID in Go instead of relying on shell expansion, which exec.Command does not perform.
	currentUser, err := user.Current()
	if err != nil {
		utils.InfoMessage("Note: Please restart Karabiner-Elements manually to load the new configuration")
		return nil
	}
	utils.InfoMessage("Restarting Karabiner-Elements to load new configuration...")
	launchctlTarget := fmt.Sprintf("gui/%s/org.pqrs.karabiner.karabiner_console_user_server", currentUser.Uid)
	if err := utils.ExecuteCommand("launchctl", "kickstart", "-k", launchctlTarget); err != nil {
		utils.InfoMessage("Note: Please restart Karabiner-Elements manually to load the new configuration")
	}

	return nil
}

func InstallConfigFiles() error {
	configDir, err := config.ConfigDir()
	if err != nil {
		return fmt.Errorf("failed to get config directory: %w", err)
	}

	apps := []Application{
		{name: "ghostty", src: filepath.Join(config.DotfilesConfigDir(), "ghostty"), dest: filepath.Join(configDir, "ghostty")},
		{name: "starship", src: filepath.Join(config.DotfilesConfigDir(), "starship", "starship.toml"), dest: filepath.Join(configDir, "starship.toml")},
		{name: "tmux", src: filepath.Join(config.DotfilesConfigDir(), "tmux"), dest: filepath.Join(configDir, "tmux")},
	}

	for _, app := range apps {
		utils.InfoMessage("Installing %s", app.name)
		err := app.Install()
		if err != nil {
			return err
		}
		utils.SuccessMessage(fmt.Sprintf("Installed %s", app.name))
	}

	utils.InfoMessage("Setting up Git")
	if err = SetupGit(); err != nil {
		return err
	}
	utils.SuccessMessage("Git setup complete")

	return nil
}

func install(src, dest string) error {
	utils.InfoMessage("Syncing " + src + " to " + dest)
	if _, err := os.Stat(dest); err == nil {
		if utils.ConfirmDestructive(fmt.Sprintf("File %s already exists, do you want to replace it?", dest)) {
			// Remove existing file or directory
			if err := utils.RemoveAllFiles(dest); err != nil {
				return fmt.Errorf("[%s]: error removing file: %w", src, err)
			}
			if err := utils.SymlinkFiles(src, dest); err != nil {
				return fmt.Errorf("[%s]: error symlinking file: %w", src, err)
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

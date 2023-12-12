package dotfiles

import (
	"fmt"
	"log"
	"os"
	"path"

	utils "github.com/pablobfonseca/dotfiles-cli/src"
	"github.com/vbauerster/mpb/v7"
)

var nvimConfigPath = path.Join(os.Getenv("HOME"), ".config", "nvim")
var nvChadRepo = "https://github.com/NvChad/NvChad"

func InstallNvim(p *mpb.Progress) {
	if nvimInstalled() {
		fmt.Println("nvim already installed, skipping...")
		return
	} else {
		installNvimBar := NewBar("Installing nvim", 1, p)
		if err := utils.ExecuteCommand("brew", "install", "neovim"); err != nil {
			log.Fatal("Error installing nvim:", err)
		}
		installNvimBar.Increment()
	}

	if utils.DirExists(nvimConfigPath) {
		fmt.Println("nvim files already exists, skipping...")
		return
	} else {
		installNvChadBar := NewBar("Installing NvChad", 1, p)

		if err := utils.ExecuteCommand("git", "clone", "--depth", "1", nvChadRepo, nvimConfigPath); err != nil {
			log.Fatal("Error cloning the repository:", err)
		}
		installNvChadBar.Increment()

		symlinkBar := NewBar("Symlinking files", 1, p)

		src := path.Join(os.Getenv("HOME"), ".dotfiles", "nvim", "custom")
		dest := path.Join(nvimConfigPath, "lua", "custom")
		if err := os.Symlink(src, dest); err != nil {
			log.Fatal("Error creating symlink:", err)
		}
		symlinkBar.Increment()
	}

}

func UninstallNvim(p *mpb.Progress) {
	if !nvimInstalled() {
		fmt.Println("nvim is not installed, skipping...")
		return
	}

	uninstallBar := NewBar("Uninstalling nvim", 2, p)

	if err := utils.ExecuteCommand("brew", "uninstall", "neovim"); err != nil {
		log.Fatal("Error uninstalling nvim:", err)
	}
	uninstallBar.Increment()

	if err := utils.ExecuteCommand("rm", "-rf", nvimConfigPath); err != nil {
		log.Fatal("Error removing nvim files:", err)
	}
	uninstallBar.Increment()
}

func nvimInstalled() bool {
	return utils.CommandExists("nvim")
}

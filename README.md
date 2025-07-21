# 🚀 Dotfiles Installer

A CLI tool to manage your personal dotfiles and development environment setup with an interactive terminal UI.

![Banner](https://img.shields.io/badge/Go-1.23+-blue.svg)
![License](https://img.shields.io/badge/license-MIT-green.svg)
![Status](https://img.shields.io/badge/status-active-brightgreen.svg)

## ✨ Features

- 🔧 **Interactive Tool Selection** - Choose exactly what you want to install
- 📊 **Installation Status Tracking** - See what's installed and what's not
- 🏃 **Dry-run Mode** - Preview installations safely before executing
- ⚙️ **Configuration Management** - Validate and manage your settings
- 🔄 **Update & Rollback** - Keep your dotfiles current
- 🛡️ **Error Handling** - Robust error recovery and reporting

## 🚀 Quick Start

### Installation

```bash
# Build from source
git clone <this-repo>
cd dotfiles-installer
go build -o dotfiles
```

### Basic Usage

```bash
# Show help and all available commands
dotfiles --help

# Interactive installation with tool selection
dotfiles install --interactive

# List all available tools and their status
dotfiles list

# Preview what would be installed (safe mode)
dotfiles install --dry-run

# Check installation status
dotfiles status

# Show current configuration
dotfiles config

# Update your dotfiles
dotfiles update --brew
```

## 🔧 Available Tools

- **Homebrew** - Package manager for macOS
- **Neovim** - Modern Vim-based text editor with config
- **Zsh** - Z shell configuration files (.zshrc, .zprofile, .zlogin)
- **Git** - Git configuration files (.gitconfig, .gitignore, etc.)
- **Tmux** - Terminal multiplexer configuration
- **Starship** - Cross-shell prompt configuration
- **Wezterm** - GPU-accelerated terminal emulator
- **Karabiner-Elements** - Keyboard customization tool
- **Aerospace** - Window manager for macOS

## 📖 Commands

| Command     | Description                           | Example                          |
| ----------- | ------------------------------------- | -------------------------------- |
| `install`   | Install dotfiles (interactive or all) | `dotfiles install --interactive` |
| `list`      | Show available tools and status       | `dotfiles list`                  |
| `status`    | Check installation status             | `dotfiles status`                |
| `config`    | Show configuration settings           | `dotfiles config`                |
| `update`    | Update repository and packages        | `dotfiles update --brew`         |
| `uninstall` | Remove dotfiles                       | `dotfiles uninstall`             |

### Command Options

- `--interactive, -i` - Interactive installation with tool selection
- `--dry-run, -n` - Preview mode (show what would be done)
- `--config` - Specify custom config file
- `--brew` - Also update Homebrew packages (for update command)

## ⚙️ Configuration

The tool uses a TOML configuration file located at `~/.config/dotfiles/config.toml`:

```toml
[dotfiles]
repository = "pablobfonseca/dotfiles"
default_dir = "/Users/yourusername/.dotfiles"
```

### First Run Setup

On first run, you'll be prompted to configure:

- **Repository**: Your GitHub dotfiles repository (format: `username/repository`)
- **Directory**: Where to clone your dotfiles (e.g., `~/.dotfiles`)

## 🎨 Screenshots

### Interactive Installation

```
🚀 Select Tools to Install

  ❯ ☑️ 📦 Homebrew
    ☐ 📦 Neovim
    ☐ 📦 Zsh
    ☐ 📦 Git Config

    Modern terminal-based package manager for macOS

• space: toggle selection  • a: select all  • enter: install  • q: quit
```

### Progress Tracking

```
🚀 Dotfiles Installation Progress

██████████████████████████████████████████ 100%

  ✅ Setting up repository
  ✅ Installing Homebrew
  ✅ Setting up Zsh configuration
  ✅ Installing configuration files
  ⚙️ Finalizing installation...

✅ Working... (4/5 steps completed)
```

## 🛠️ Development

### Project Structure

```
dotfiles-installer/
├── cmd/                    # CLI commands
│   ├── install.go         # Install command with TUI
│   ├── list.go           # List available tools
│   ├── status.go         # Show status
│   ├── config.go         # Configuration management
│   └── update.go         # Update functionality
├── src/
│   ├── config/           # Configuration management
│   ├── installer/        # Installation logic
│   ├── ui/              # Terminal UI components
│   └── utils/           # Utility functions
└── main.go
```

### Building

```bash
# Build for current platform
go build -o dotfiles

# Build for release
make build

# Run tests
go test ./...

# Clean build artifacts
make clean
```

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- [Charm](https://github.com/charmbracelet) for the amazing TUI libraries
- [Cobra](https://github.com/spf13/cobra) for CLI framework
- [Viper](https://github.com/spf13/viper) for configuration management

---

**Built with ❤️ and Go** | **Personal dotfiles:** [pablobfonseca/dotfiles](https://github.com/pablobfonseca/dotfiles)

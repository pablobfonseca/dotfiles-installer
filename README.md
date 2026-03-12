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
- 🔄 **Update** - Keep your dotfiles current
- 🛡️ **Error Handling** - Robust error recovery and reporting

## 🚀 Quick Start

### Installation

#### Option 1: Quick Install Script (Easiest)

```bash
# One-liner install script (detects your OS/architecture)
curl -fsSL https://raw.githubusercontent.com/pablobfonseca/dotfiles-installer/main/install.sh | bash
```

#### Option 2: Download Pre-built Binary

```bash
# Download the latest release for macOS (Intel)
curl -L -o dotfiles https://github.com/pablobfonseca/dotfiles-installer/releases/latest/download/dotfiles-darwin-amd64

# Make it executable
chmod +x dotfiles

# Move to your PATH (optional)
sudo mv dotfiles /usr/local/bin/
```

**Available architectures:**
- **macOS Intel**: `dotfiles-darwin-amd64`
- **macOS Apple Silicon**: `dotfiles-darwin-arm64` 
- **Linux x86_64**: `dotfiles-linux-amd64`
- **Linux ARM64**: `dotfiles-linux-arm64`

#### Option 3: Build from Source

```bash
# Clone the repository
git clone https://github.com/pablobfonseca/dotfiles-installer.git
cd dotfiles-installer

# Build the binary
go build -o dotfiles

# Or use make
make build
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

### 🔧 Installation Notes

- **macOS**: No additional dependencies required
- **Linux**: Requires `curl` for downloading
- **PATH**: The installer automatically tries to install to `/usr/local/bin` which should be in your PATH
- **Permissions**: May require `sudo` for system-wide installation

### 🚀 First Run

After installation, run the setup:

```bash
# First time setup with interactive tool selection
dotfiles install --interactive

# Or install everything at once
dotfiles install
```

## 🔧 Available Tools

- **Homebrew** - Package manager for macOS
- **Neovim** - Modern Vim-based text editor with config
- **Zsh** - Z shell configuration files (.zshrc, .zprofile, .zlogin)
- **Git** - Git configuration files (.gitconfig, .gitignore, etc.)
- **Tmux** - Terminal multiplexer configuration
- **Starship** - Cross-shell prompt configuration
- **Ghostty** - Fast, native terminal emulator
- **Karabiner-Elements** - Keyboard customization tool
- **Cyberpunk Theme** - Cyberpunk color theme for terminal tools

## 📖 Commands

| Command       | Description                               | Example                          |
| ------------- | ----------------------------------------- | -------------------------------- |
| `install`     | Install dotfiles (interactive or all)     | `dotfiles install --interactive` |
| `list`        | Show available tools and status           | `dotfiles list`                  |
| `status`      | Check installation status                 | `dotfiles status`                |
| `config`      | Show configuration settings               | `dotfiles config`                |
| `update`      | Update repository and packages            | `dotfiles update --brew`         |
| `uninstall`   | Remove dotfiles                           | `dotfiles uninstall`             |
| `dashboard`   | Launch interactive TUI monitoring dashboard | `dotfiles dashboard`           |

### Subcommands

| Command                         | Description                                        |
| ------------------------------- | -------------------------------------------------- |
| `dotfiles install nvim`         | Install Neovim and config                          |
| `dotfiles install zsh`          | Install Zsh configuration                          |
| `dotfiles install homebrew`     | Install Homebrew and packages                      |
| `dotfiles install karabiner`    | Install Karabiner-Elements                         |
| `dotfiles uninstall nvim`       | Uninstall Neovim config (`--uninstall-app` to also remove the app) |

### Command Options

- `--interactive, -i` - Interactive installation with tool selection
- `--dry-run, -n` - Preview mode (show what would be done)
- `--force, -f` - Force overwrite existing files without confirmation
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
│   ├── homebrew/          # Homebrew subcommand
│   ├── nvim/              # Neovim install/uninstall subcommands
│   ├── zsh/               # Zsh subcommand
│   ├── karabiner/         # Karabiner subcommand
│   ├── install.go         # Install command with TUI
│   ├── list.go            # List available tools
│   ├── status.go          # Show status
│   ├── config.go          # Configuration management
│   ├── update.go          # Update functionality
│   ├── uninstall.go       # Uninstall command
│   └── dashboard.go       # Dashboard TUI command
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

## 📦 Releases

Pre-built binaries are automatically created for each release and available on the [GitHub Releases page](https://github.com/pablobfonseca/dotfiles-installer/releases).

**Supported platforms:**
- macOS (Intel & Apple Silicon)
- Linux (x86_64 & ARM64)

**Release artifacts:**
- `dotfiles-darwin-amd64` - macOS Intel
- `dotfiles-darwin-arm64` - macOS Apple Silicon  
- `dotfiles-linux-amd64` - Linux x86_64
- `dotfiles-linux-arm64` - Linux ARM64

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

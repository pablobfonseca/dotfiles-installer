#!/bin/bash

# Dotfiles Installer - Installation Script
# https://github.com/pablobfonseca/dotfiles-installer

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# GitHub repository info
REPO="pablobfonseca/dotfiles-installer"
GITHUB_URL="https://github.com/${REPO}"
RELEASES_URL="${GITHUB_URL}/releases"

# Functions
print_info() {
    echo -e "${BLUE}ℹ️  $1${NC}"
}

print_success() {
    echo -e "${GREEN}✅ $1${NC}"
}

print_error() {
    echo -e "${RED}❌ $1${NC}"
}

print_warning() {
    echo -e "${YELLOW}⚠️  $1${NC}"
}

# Detect OS and architecture
detect_platform() {
    local os=""
    local arch=""

    # Detect OS
    case "$(uname -s)" in
        Darwin)
            os="darwin"
            ;;
        Linux)
            os="linux"
            ;;
        *)
            print_error "Unsupported operating system: $(uname -s)"
            exit 1
            ;;
    esac

    # Detect architecture
    case "$(uname -m)" in
        x86_64)
            arch="amd64"
            ;;
        arm64|aarch64)
            arch="arm64"
            ;;
        *)
            print_error "Unsupported architecture: $(uname -m)"
            exit 1
            ;;
    esac

    echo "${os}-${arch}"
}

# Check if command exists
command_exists() {
    command -v "$1" >/dev/null 2>&1
}

# Main installation function
install_dotfiles() {
    print_info "🚀 Dotfiles Installer Setup"
    echo

    # Check dependencies
    if ! command_exists curl; then
        print_error "curl is required but not installed. Please install curl and try again."
        exit 1
    fi

    # Detect platform
    platform=$(detect_platform)
    print_info "Detected platform: $platform"

    # Set binary name
    binary_name="dotfiles-${platform}"
    
    # Download URL
    download_url="${RELEASES_URL}/latest/download/${binary_name}"
    
    # Installation directory
    install_dir="/usr/local/bin"
    binary_path="${install_dir}/dotfiles"
    
    print_info "Downloading from: $download_url"
    
    # Create temporary file
    temp_file=$(mktemp)
    
    # Download binary
    if curl -fL "$download_url" -o "$temp_file"; then
        print_success "Downloaded successfully"
    else
        print_error "Failed to download binary from $download_url"
        print_info "Please check if the release exists: $RELEASES_URL"
        rm -f "$temp_file"
        exit 1
    fi
    
    # Make executable
    chmod +x "$temp_file"
    
    # Check if we need sudo for installation
    if [ -w "$install_dir" ]; then
        mv "$temp_file" "$binary_path"
    else
        print_info "Installing to $install_dir (requires sudo)"
        sudo mv "$temp_file" "$binary_path"
    fi
    
    # Verify installation
    if [ -x "$binary_path" ]; then
        print_success "Installed successfully to $binary_path"
        
        # Check if directory is in PATH
        if echo "$PATH" | grep -q "$install_dir"; then
            print_success "Ready to use! Try: dotfiles --help"
        else
            print_warning "$install_dir is not in your PATH"
            print_info "Add this to your shell profile: export PATH=\"$install_dir:\$PATH\""
        fi
        
        echo
        print_info "🎉 Installation complete!"
        print_info "Get started with: dotfiles install --interactive"
        
    else
        print_error "Installation verification failed"
        exit 1
    fi
}

# Run installation
install_dotfiles
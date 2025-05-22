#!/bin/bash

# Garden Ecosystem Installation Script
# Install the Sprout CLI tool for weather context preservation

set -e

# Colors and formatting
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
BOLD='\033[1m'
RESET='\033[0m'

# Configuration
GITHUB_REPO="sprouted-dev/garden"
BINARY_NAME="sprout"
INSTALL_DIR="/usr/local/bin"
TEMP_DIR=$(mktemp -d)

# Platform detection
detect_platform() {
    case "$(uname -s)" in
        Darwin*)
            case "$(uname -m)" in
                arm64) echo "darwin-arm64" ;;
                x86_64) echo "darwin-amd64" ;;
                *) 
                    echo -e "${RED}❌ Unsupported macOS architecture: $(uname -m)${RESET}" >&2
                    exit 1
                    ;;
            esac
            ;;
        Linux*)
            case "$(uname -m)" in
                aarch64|arm64) echo "linux-arm64" ;;
                x86_64|amd64) echo "linux-amd64" ;;
                *)
                    echo -e "${RED}❌ Unsupported Linux architecture: $(uname -m)${RESET}" >&2
                    exit 1
                    ;;
            esac
            ;;
        CYGWIN*|MINGW32*|MSYS*|MINGW*)
            echo "windows-amd64"
            ;;
        *)
            echo -e "${RED}❌ Unsupported operating system: $(uname -s)${RESET}" >&2
            exit 1
            ;;
    esac
}

# Check if command exists
command_exists() {
    command -v "$1" >/dev/null 2>&1
}

# Print header
print_header() {
    echo -e "${BOLD}${BLUE}"
    echo "╭─────────────────────────────────────────────────────────────────────╮"
    echo "│                        🌱 Garden Ecosystem                         │"
    echo "│                     Sprout CLI Installation                        │"
    echo "╰─────────────────────────────────────────────────────────────────────╯"
    echo -e "${RESET}"
}

# Print success message
print_success() {
    echo -e "${GREEN}${BOLD}"
    echo "╭─────────────────────────────────────────────────────────────────────╮"
    echo "│                    ✅ Installation Complete!                       │"
    echo "╰─────────────────────────────────────────────────────────────────────╯"
    echo -e "${RESET}"
    echo -e "${GREEN}Sprout CLI has been installed successfully!${RESET}"
    echo ""
    echo -e "${YELLOW}Quick Start:${RESET}"
    echo -e "  ${BLUE}sprout weather${RESET}                    # Show current development context"
    echo -e "  ${BLUE}sprout weather --for-ai${RESET}           # Get AI-friendly context"  
    echo -e "  ${BLUE}sprout weather --onboard-ai${RESET}       # Full AI assistant onboarding"
    echo ""
    echo -e "${YELLOW}Documentation:${RESET}"
    echo -e "  ${BLUE}https://sprouted.dev${RESET}              # Official website"
    echo -e "  ${BLUE}https://github.com/$GITHUB_REPO${RESET}  # Source code"
    echo ""
    echo -e "${YELLOW}Get started with Weather System:${RESET}"
    echo -e "  ${BLUE}cd your-project && sprout weather${RESET}"
}

# Clean up on exit
cleanup() {
    if [ -d "$TEMP_DIR" ]; then
        rm -rf "$TEMP_DIR"
    fi
}
trap cleanup EXIT

# Main installation function
install_sprout() {
    print_header
    
    echo -e "${YELLOW}🔍 Detecting platform...${RESET}"
    PLATFORM=$(detect_platform)
    echo -e "${GREEN}✅ Platform detected: $PLATFORM${RESET}"
    
    # Check prerequisites
    echo -e "${YELLOW}🔧 Checking prerequisites...${RESET}"
    
    if ! command_exists curl && ! command_exists wget; then
        echo -e "${RED}❌ Neither curl nor wget found. Please install one of them.${RESET}" >&2
        exit 1
    fi
    
    # Determine download command
    if command_exists curl; then
        DOWNLOAD_CMD="curl -fsSL"
    else
        DOWNLOAD_CMD="wget -qO-"
    fi
    
    echo -e "${GREEN}✅ Prerequisites satisfied${RESET}"
    
    # Get latest release information
    echo -e "${YELLOW}📡 Fetching latest release information...${RESET}"
    
    if command_exists curl; then
        LATEST_RELEASE=$(curl -fsSL "https://api.github.com/repos/$GITHUB_REPO/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')
    else
        LATEST_RELEASE=$(wget -qO- "https://api.github.com/repos/$GITHUB_REPO/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')
    fi
    
    if [ -z "$LATEST_RELEASE" ]; then
        echo -e "${YELLOW}⚠️  Could not fetch latest release. Using 'latest' tag.${RESET}"
        LATEST_RELEASE="latest"
        DOWNLOAD_URL="https://github.com/$GITHUB_REPO/releases/latest/download/$BINARY_NAME-$PLATFORM"
    else
        echo -e "${GREEN}✅ Latest version: $LATEST_RELEASE${RESET}"
        if [ "$PLATFORM" = "windows-amd64" ]; then
            DOWNLOAD_URL="https://github.com/$GITHUB_REPO/releases/download/$LATEST_RELEASE/$BINARY_NAME-$PLATFORM.exe"
        else
            DOWNLOAD_URL="https://github.com/$GITHUB_REPO/releases/download/$LATEST_RELEASE/$BINARY_NAME-$PLATFORM"
        fi
    fi
    
    # Download binary
    echo -e "${YELLOW}⬇️  Downloading Sprout CLI...${RESET}"
    echo -e "${BLUE}URL: $DOWNLOAD_URL${RESET}"
    
    cd "$TEMP_DIR"
    
    if [ "$PLATFORM" = "windows-amd64" ]; then
        BINARY_FILE="$BINARY_NAME.exe"
    else
        BINARY_FILE="$BINARY_NAME"
    fi
    
    if ! $DOWNLOAD_CMD "$DOWNLOAD_URL" > "$BINARY_FILE"; then
        echo -e "${RED}❌ Failed to download Sprout CLI${RESET}" >&2
        echo -e "${YELLOW}This might mean there's no pre-built binary for your platform yet.${RESET}"
        echo -e "${YELLOW}You can build from source: https://github.com/$GITHUB_REPO${RESET}"
        exit 1
    fi
    
    # Make executable (not needed on Windows)
    if [ "$PLATFORM" != "windows-amd64" ]; then
        chmod +x "$BINARY_FILE"
    fi
    
    echo -e "${GREEN}✅ Downloaded successfully${RESET}"
    
    # Install binary
    echo -e "${YELLOW}📦 Installing Sprout CLI...${RESET}"
    
    # Check if we need sudo
    if [ ! -w "$INSTALL_DIR" ]; then
        echo -e "${YELLOW}🔐 Administrator privileges required for installation to $INSTALL_DIR${RESET}"
        if ! sudo cp "$BINARY_FILE" "$INSTALL_DIR/"; then
            echo -e "${RED}❌ Failed to install to $INSTALL_DIR${RESET}" >&2
            
            # Try user-local installation
            LOCAL_BIN="$HOME/.local/bin"
            echo -e "${YELLOW}🏠 Trying user-local installation to $LOCAL_BIN${RESET}"
            
            mkdir -p "$LOCAL_BIN"
            if cp "$BINARY_FILE" "$LOCAL_BIN/"; then
                echo -e "${GREEN}✅ Installed to $LOCAL_BIN${RESET}"
                echo -e "${YELLOW}⚠️  Make sure $LOCAL_BIN is in your PATH${RESET}"
                echo -e "${BLUE}Add this to your shell profile:${RESET}"
                echo -e "${BLUE}export PATH=\"\$HOME/.local/bin:\$PATH\"${RESET}"
                INSTALL_DIR="$LOCAL_BIN"
            else
                echo -e "${RED}❌ Installation failed${RESET}" >&2
                exit 1
            fi
        else
            echo -e "${GREEN}✅ Installed to $INSTALL_DIR${RESET}"
        fi
    else
        if cp "$BINARY_FILE" "$INSTALL_DIR/"; then
            echo -e "${GREEN}✅ Installed to $INSTALL_DIR${RESET}"
        else
            echo -e "${RED}❌ Installation failed${RESET}" >&2
            exit 1
        fi
    fi
    
    # Verify installation
    echo -e "${YELLOW}🔍 Verifying installation...${RESET}"
    
    if command_exists "$BINARY_NAME"; then
        VERSION_OUTPUT=$($BINARY_NAME weather --help 2>&1 | head -1 || echo "Sprout CLI")
        echo -e "${GREEN}✅ Installation verified: $VERSION_OUTPUT${RESET}"
    else
        echo -e "${YELLOW}⚠️  Binary installed but not found in PATH${RESET}"
        echo -e "${YELLOW}You may need to restart your shell or add $INSTALL_DIR to PATH${RESET}"
    fi
    
    print_success
}

# Handle command line arguments
case "${1:-}" in
    --help|-h)
        echo "Garden Ecosystem Installation Script"
        echo ""
        echo "Usage: $0 [options]"
        echo ""
        echo "Options:"
        echo "  --help, -h          Show this help message"
        echo "  --version, -v       Show version information"
        echo ""
        echo "This script will:"
        echo "  1. Detect your platform (Linux, macOS, Windows)"
        echo "  2. Download the appropriate Sprout CLI binary"
        echo "  3. Install it to /usr/local/bin (or ~/.local/bin)"
        echo "  4. Verify the installation"
        echo ""
        echo "Examples:"
        echo "  curl -fsSL https://sprouted.dev/install.sh | sh"
        echo "  wget -qO- https://sprouted.dev/install.sh | sh"
        echo ""
        exit 0
        ;;
    --version|-v)
        echo "Garden Ecosystem Installation Script v1.0.0"
        exit 0
        ;;
    "")
        # No arguments, proceed with installation
        install_sprout
        ;;
    *)
        echo -e "${RED}❌ Unknown option: $1${RESET}" >&2
        echo "Use --help for usage information"
        exit 1
        ;;
esac
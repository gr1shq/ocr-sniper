# OCR Sniper

Fast screen OCR tool written in Go. Select text on your screen, and it's instantly copied to your clipboard.

## Features

- Instant OCR - Select area, text copied to clipboard
- Privacy First - Everything runs locally
- Cross-Platform - Works on Linux, macOS, Windows
- Single Binary - No runtime dependencies

## Prerequisites

You need Tesseract OCR installed on your system.

Linux (Ubuntu/Debian):
  sudo apt install tesseract-ocr tesseract-ocr-eng

macOS:
  brew install tesseract

Windows:
  Download from: https://github.com/UB-Mannheim/tesseract/wiki
  
  After installation, add Tesseract to your PATH:
  1. Open System Properties → Environment Variables
  2. Add C:\Program Files\Tesseract-OCR to PATH


## Installation

### Option 1: Download Pre-built Binary

1. Go to Releases page
2. Download for your OS
3. Run the binary

Linux/macOS - Make executable first:
  chmod +x ocr-sniper-linux
  ./ocr-sniper-linux

Windows:
  Double-click ocr-sniper-windows.exe

### Option 2: Build from Source

  git clone https://github.com/YOUR_USERNAME/ocr-sniper.git
  cd ocr-sniper
  go mod download
  go build -o ocr-sniper
  ./ocr-sniper

## Usage

1. Run the binary (or set up hotkey - see below)
2. Select text area with your mouse
3. Release mouse button
4. Text is automatically copied to clipboard!

### Platform Notes

| OS | Selection Mode | Tool Used |
|----|---------------|-----------|
| Linux | Interactive (click and drag) | gnome-screenshot |
| macOS | Interactive (click and drag) | screencapture |
| Windows | Full screen (v1.0) | Native API |

## Setting Up Global Hotkey

### Linux (GNOME)

1. Open Settings → Keyboard → View and Customize Shortcuts
2. Click "Add Custom Shortcut"
3. Name: OCR Sniper
4. Command: /full/path/to/ocr-sniper-linux
5. Set shortcut: Ctrl+Alt+O

### macOS

1. Open System Settings → Keyboard → Keyboard Shortcuts
2. Click "+" to add new shortcut
3. Select the app and assign hotkey

### Windows

1. Right-click ocr-sniper-windows.exe → Create Shortcut
2. Right-click shortcut → Properties
3. In "Shortcut key" field, press Ctrl+Alt+O
4. Click OK

## Troubleshooting

| Problem | Solution |
|---------|----------|
| Tesseract not found | Install Tesseract and add to PATH |
| Screenshot failed (Linux) | Make sure you're using GNOME desktop |
| Screenshot failed (Mac) | Grant Screen Recording permission in System Settings |
| Permission denied (Linux/Mac) | Run chmod +x ocr-sniper |
| No text detected | Try selecting larger area with clearer text |
| Windows captures full screen | Selection support coming in v1.1 |

## Building for Development

Install dependencies:
  go mod download

Run directly:
  go run main.go

Build for current OS:
  go build -o ocr-sniper

Build for all platforms:
  GOOS=linux GOARCH=amd64 go build -o ocr-sniper-linux
  GOOS=darwin GOARCH=amd64 go build -o ocr-sniper-mac
  GOOS=windows GOARCH=amd64 go build -o ocr-sniper-windows.exe

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments

- Tesseract OCR
- clipboard (github.com/atotto/clipboard)
- screenshot (github.com/kbinani/screenshot)
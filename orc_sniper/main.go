package main

import (
	"fmt"
	"image/png"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/kbinani/screenshot"
)

func main() {
	fmt.Println("✅ OCR Sniper Ready!")
	fmt.Println("📸 Select an area with your mouse...")
	fmt.Println()

	tempPath := "/tmp/ocr_capture.png"
	var err error

	osType := runtime.GOOS

	if osType == "linux" {
		err = captureLinux(tempPath)
	} else if osType == "darwin" {
		err = captureMac(tempPath)
	} else if osType == "windows" {
		err = captureWindows(tempPath)
	} else {
		fmt.Println("❌ Unsupported OS:", osType)
		return
	}

	if err != nil {
		fmt.Println("❌ Screenshot failed:", err)
		return
	}

	fmt.Println("✅ Screenshot captured!")

	// Check if Tesseract is installed
	_, err = exec.LookPath("tesseract")
	if err != nil {
		fmt.Println("❌ Tesseract not found!")
		fmt.Println("💡 Install: sudo apt install tesseract-ocr")
		return
	}

	// Run OCR using tesseract command line
	cmd := exec.Command("tesseract", tempPath, "stdout", "-l", "eng")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("❌ OCR failed:", err)
		return
	}

	text := strings.TrimSpace(string(output))

	fmt.Println()
	fmt.Println("📝 OCR Result:")
	fmt.Println("━━━━━━━━━━━━━━")
	fmt.Println(text)
	fmt.Println("━━━━━━━━━━━━━━")
	fmt.Println()

	if text != "" {
		clipboard.WriteAll(text)
		fmt.Println("📋 Copied to clipboard!")
	} else {
		fmt.Println("⚠️  No text detected")
	}

	os.Remove(tempPath)
}

func captureLinux(path string) error {
	cmd := exec.Command("gnome-screenshot", "-a", "-f", path)
	return cmd.Run()
}

func captureMac(path string) error {
	cmd := exec.Command("screencapture", "-i", "-s", path)
	return cmd.Run()
}

func captureWindows(path string) error {
	n := screenshot.NumActiveDisplays()
	if n == 0 {
		return fmt.Errorf("no displays found")
	}

	bounds := screenshot.GetDisplayBounds(0)

	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		return err
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		return err
	}

	fmt.Println("⚠️  Windows: Full screen captured (selection not supported yet)")
	return nil
}
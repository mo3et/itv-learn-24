package main

import (
	"fmt"
	"os"

	"github.com/dsoprea/go-exif/v3"
)

func main() {
	// check cli args
	if len(os.Args) < 2 {
		fmt.Println("plz provide the path to the image file")
		return
	}

	// Get file Path
	filepath := os.Args[1]

	// Read image file
	imgData, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println("Error reading image file:", err)
		return
	}

	// 解析 EXIF Data
	rawExif, err := exif.SearchAndExtractExif(imgData)
	if err != nil {
		fmt.Println("Error extracting EXIF data:", err)
		return
	}

	// Remove EXIF Data
	strippedData := exif.Re(rawExif)
}

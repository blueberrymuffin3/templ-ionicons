package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// Set the source and destination directories
	sourceDir := "ionicons/src/svg/"
	destinationDir := "data/"

	// Create the destination directory if it doesn't exist
	err := os.MkdirAll(destinationDir, 0755)
	if err != nil {
		fmt.Println("Error creating destination directory:", err)
		return
	}

	// Get a list of SVG files in the source directory
	files, err := filepath.Glob(filepath.Join(sourceDir, "*.svg"))
	if err != nil {
		fmt.Println("Error getting list of SVG files:", err)
		return
	}

	// Iterate through each SVG file
	for _, file := range files {
		// Read the content of the SVG file
		content, err := os.ReadFile(file)
		if err != nil {
			fmt.Println("Error reading file:", err)
			continue
		}

		// Parse the SVG content
		svg, err := extractBody(content)
		if err != nil {
			fmt.Println("Error parsing SVG:", err)
		}

		minified := minify(svg)

		// Create the destination file path
		destinationFile := filepath.Join(destinationDir, filepath.Base(file))

		// Write the SVG children content to the destination file
		err = os.WriteFile(destinationFile, []byte(minified), 0644)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			continue
		}
	}

	fmt.Println("Processing completed. SVG files saved in", destinationDir)
}

type SVG struct {
	XMLName xml.Name `xml:"svg"`
	Content string   `xml:",innerxml"`
}

func extractBody(content []byte) (string, error) {
	var svg SVG
	err := xml.Unmarshal(content, &svg)
	if err != nil {
		return "", err
	}
	return svg.Content, nil
}

func minify(content string) string {
	output := ""

	// For each line, remove line breaks and indentation
	for _, line := range strings.Split(content, "\n") {
		output += strings.TrimSpace(line)
	}

	return output
}

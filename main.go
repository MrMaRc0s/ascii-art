package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Ensure correct argument usage
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Print("Usage: go run <program.go> <input> + <banner>(optional)")
		os.Exit(2)
	}
	var banner string = "banner/standard.txt"
	if len(os.Args) == 3 {
		banner = os.Args[2]
	}
	// Read the file
	bannerFile, err := os.ReadFile(banner)
	if err != nil {
		fmt.Print("There was a problem reading the banner file template")
		os.Exit(1)
	}

	// Split the lines of the file into a slice
	lines := strings.Split(string(bannerFile), "\n")

	// Split the input string into lines, based on literal newline characters
	sepText := strings.Split(os.Args[1], "\\n") // Handling literal "\n" from the input

	// Print the ASCII Art based on the input
	PrintAsciiArtRecursive(sepText, lines)
}

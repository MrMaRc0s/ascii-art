package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Read the file
	bannerFile, err := os.ReadFile("banner/standard.txt")
	if err != nil {
		fmt.Print("There was a problem reading the banner file template")
		os.Exit(1)
	}

	// Ensure correct argument usage
	if len(os.Args) != 2 {
		fmt.Print("Usage: go run <program.go> <input>")
		os.Exit(2)
	}

	// Split the lines of the file into a slice
	lines := strings.Split(string(bannerFile), "\n")

	// Split the input string into lines, based on literal newline characters
	sepText := strings.Split(os.Args[1], "\\n") // Handling literal "\n" from the input

	// Print the ASCII Art based on the input
	PrintAsciiArtRecursive(sepText, lines)
}

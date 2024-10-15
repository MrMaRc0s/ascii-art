package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Read the file
	file, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Print("There was a problem opening our data")
		os.Exit(1)
	}

	// Ensure correct argument usage
	if len(os.Args) != 2 {
		fmt.Print("Usage: go run <program.go> <input>")
		os.Exit(2)
	}

	// Split the lines of the file into a slice
	lines := strings.Split(string(file), "\n")

	// Capture the input from the user
	inputText := os.Args[1]

	// Split the input string into lines, based on literal newline characters
	sepText := strings.Split(inputText, "\\n") // Handling literal "\n" from the input

	// Print the ASCII Art based on the input
	printAsciiArtRecursive(sepText, lines)
}

func printAsciiArtRecursive(sentences []string, textFile []string) {
	// Base case: If no sentences left, return
	if len(sentences) == 0 {
		return
	}

	// Process the first sentence
	if sentences[0] != "" {
		printSentenceAsciiRecursive(sentences[0], textFile, 1)
		fmt.Print()
	} else {
		// If the sentence is empty (i.e., there was a '\n'), print an empty line
		fmt.Println()
	}

	// Recursive call to process the remaining sentences
	printAsciiArtRecursive(sentences[1:], textFile)
}

func printSentenceAsciiRecursive(word string, textFile []string, h int) {
	// Base case: If we've processed all the lines for this character height, return
	if h > 8 {
		return
	}

	// Print the current line for each character in the word
	for i := 0; i < len(word); i++ {
		for lineIndex, line := range textFile {
			if lineIndex == (int(word[i])-32)*9+h {
				// Print the line corresponding to the current character and line number
				fmt.Print(line)
			}
		}
	}
	fmt.Println()

	// Recursive call for the next line height
	printSentenceAsciiRecursive(word, textFile, h+1)
}

/*func printAsciiArt(sentences []string, textFile []string) {
	// Loop through each sentence (split by '\n')
	for i, word := range sentences {
		if word == "" {
			// If the sentence is empty (i.e., there was a '\n'), print an empty line
			if i != 0 {
				fmt.Println()
			}
			continue
		}

		for h := 1; h < 9; h++ {
			for i := 0; i < len(word); i++ {
				for lineIndex, line := range textFile {
					if lineIndex == (int(word[i])-32)*9+h {
						// Print the line corresponding to the current character and line number
						fmt.Print(line)
					}
				}
			}
			fmt.Println()
		}
	}
}*/

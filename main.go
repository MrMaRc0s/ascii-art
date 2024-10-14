package main

import (
	"fmt"
	"os"

	// "strconv"
	"strings"
)

func main() {
	inputFile := "standard.txt"
	// inputFile = "standard.txt"
	file, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	text := "Sample"
	if len(os.Args) > 1 {
		text = os.Args[1]
	}
	fileString := string(file)
	fileString = strings.Replace(fileString, "\r", "", -1)
	lines := strings.Split(fileString, "\n")
	text_lines := strings.Split(text, "\\n")

	for l := 0; l < len(text_lines); l++ {
		runes := []rune(text_lines[l])

		for i := 0; i < 8; i++ {
			for j := 0; j < len(runes); j++ {
				// fmt.Print(j)
				pos := int(runes[j] - 32)
				pos2 := 1 + (int(runes[j])-32)*9 + i
				// fmt.Printf("%d ", pos2)
				letterLen := length(lines, pos)
				// fmt.Print(letterLen)
				fmt.Print(lines[pos2])

				printSpaces(letterLen - len(lines[pos2]))
			}
			fmt.Println()
		}
	}
}

func printSpaces(n int) {
	for i := 0; i < n; i++ {
		fmt.Printf(" ")
	}
}

func length(list []string, ind int) int {
	i := 1 + 9*ind
	// fmt.Print(i)
	// fmt.Print("IND"+strconv.Itoa(i))
	max := 1
	for ; i < 1+9*ind+9; i++ {
		lineLength := len(list[i])
		if lineLength > max {
			max = lineLength
		}
	}
	return max
}

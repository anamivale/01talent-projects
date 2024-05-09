package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	testFile, _ := os.ReadFile("standard.txt")
	inputString := os.Args[1]

	inputFile := strings.Split(string(testFile), "\n")

	output := AsciiArt(inputString, inputFile)
	fmt.Print(output)
}

func AsciiArt(input string, inputFile []string) string {
	var result strings.Builder
	var newLinesOnly strings.Builder
	sepInputString := strings.Split(input, "\\n")

	newLines := OnlyNewLines(sepInputString)
	if newLines != "false" {
		newLinesOnly.WriteString(newLines)
		return newLinesOnly.String()

	}

	for _, words := range sepInputString {
		if words == "" {
			result.WriteString("\n")
		} else {
			for i := 0; i < len(words); {
				for j := 0; j < 8; {
					start := (int(words[i]-32) * 9) + 1
					result.WriteString(inputFile[start+j])
					i++
					if i == len(words) {
						if j == 7 {
							result.WriteString("\n")
							break
						}
						result.WriteString("\n")
						j++
						i = 0

					}
				}
			}
		}
	}
	return result.String()
}

func OnlyNewLines(sepInputString []string) string {
	empty := ""
	for i, words := range sepInputString {
		if words != "" {
			return "false"
		}
		if i == 0 && words == "" {
			continue
		}
		empty += "\n"

	}

	return empty
}

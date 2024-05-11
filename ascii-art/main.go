package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// checkErr helps in Printing errors that occur when you use some package methods.
func checkErr(err error) {
	if err != nil {
		fmt.Println("You have the following error:", err)
	}
}

// CheckFileName helps get a file with the write extension.
func CheckFileName(fileName string) string {
	if filepath.Ext(fileName) != ".txt" {
		return ""
	}
	return string(fileName)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Only need 2 arguments but got %d\n", len(os.Args))
		return
	}
	File := "standard.txt"
	// File := os.Args[1]
	// if File != "standard.txt" || File != "shadow.txt" || File != "thinkeratoy.txt" {
	// 	fmt.Println("Banner file should be standard.txt or shadow.txt or thinkeratoy.txt")
	// }
	bannerFile := CheckFileName(File)
	if bannerFile == "" {
		fmt.Println("needs a file name not an empty strin")
		return
	}
	bannerContent, err := os.ReadFile(bannerFile)
	checkErr(err)

	inputString := os.Args[1]

	splitBannerContent := strings.Split(string(bannerContent), "\n")

	output := AsciiArt(inputString, splitBannerContent)
	fmt.Print(output)
}

// AsciiArt gives the art of the given words in the format required.
func AsciiArt(input string, inputFile []string) string {
	var result strings.Builder
	var newLinesOnly strings.Builder
	input = strings.ReplaceAll(input, "\n", "\\n")
	input = strings.ReplaceAll(input, "\\t", "    ")
	sepInputString := strings.Split(input, "\\n")

	newLines := OnlyNewLines(sepInputString)
	if newLines != "false" {
		newLinesOnly.WriteString(newLines)
		return newLinesOnly.String()

	} else {
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
	}
	return result.String()
}

// OnlyNewLines checks if the input string contains only new lines.
func OnlyNewLines(sepInputString []string) string {
	empty := ""
	for i, words := range sepInputString {
		if words != "" {
			return "false"
		}
		if words == "" && i == 0 {
			continue
		}
		empty += "\n"

	}

	return empty
}

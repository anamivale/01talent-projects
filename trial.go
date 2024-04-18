package main

import (
	"fmt"
	"regexp"
)

func main() {
	// Input string
	input := "AbrahaBmb"

	// Regular expression pattern to match 'b' (either capital or small)
	pattern := regexp.MustCompile("[bB]")

	// Find all occurrences of 'b' (either capital or small)
	indices := pattern.FindAllStringIndex(input, -1)

	// Get the index of the last occurrence of 'b'
	lastIndex := indices[len(indices)-1]

	// Get the substring from the input string using the index of the last 'b'
	lastB := input[lastIndex[0]:lastIndex[1]]

	// Print the result
	fmt.Println("Last 'b' (either capital or small):", lastB)

	// Input string
	input1 := "Araham b  B"

	// Regular expression pattern to match the first instance of the letter "b" (either capital or small)
	pattern1 := regexp.MustCompile("[bB]")

	// Find the first instance of the letter "b" (either capital or small)
	firstB := pattern1.ReplaceAllString(input1, )

	// Print the result
	fmt.Println("First 'b' (either capital or small):", firstB)
}

package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func hexToDec(words []string) []string {
	for i, word := range words {
		if word == "(hex)" {
			dec, err := strconv.ParseInt(words[i-1], 16, 64)
			checkErr(err)
			words[i-1] = strconv.Itoa(int(dec))
			words = append(words[0:i], words[i+1:]...)

		}
	}
	return words
}

func binToDec(words []string) []string {
	for i, word := range words {
		if word == "(bin)" {
			dec, err := strconv.ParseInt(words[i-1], 2, 64)
			checkErr(err)
			words[i-1] = strconv.Itoa(int(dec))
			words = append(words[0:i], words[i+1:]...)

		}
	}
	return words
}

func uppercase(words []string) []string {
	for i, word := range words {
		if word == "(up)" {
			words[i-1] = strings.ToUpper(words[i-1])
			words = append(words[:i], words[i+1:]...)
		}
	}
	return words
}

func capitalization(words []string) []string {
	for i := 0; i < len(words); i++ {
		// for i, word := range words {
		if words[i] == "(cap)" {
			words[i-1] = strings.ToUpper(string(words[i-1][0])) + strings.ToLower(words[i-1][1:])
			if i != len(words)-1 {
				words = append(words[0:i], words[i+1:]...)
			} else {
				words = words[:i]
			}
		}
	}
	return words
}

func lowercase(words []string) []string {
	for i, word := range words {
		if word == "(low)" {
			words[i-1] = strings.ToLower(words[i-1])
			words = append(words[0:i], words[i+1:]...)
		}
	}
	return words
}

func capapitalizeNWords(words []string) []string {
	for i := 0; i < len(words); i++ {
		if words[i] == "(cap," {
			nstr := words[i+1][:len(words[i+1])-1]
			n, err := strconv.Atoi(nstr)
			checkErr(err)
			if i > n {
				for j := i - n; j < i; j++ {
					words[j] = strings.ToUpper(string(words[j][0])) + strings.ToLower(words[j][1:])
				}
			}
			words = append(words[:i], words[i+2:]...)

		}
	}

	return words
}

func uppercaseNWords(words []string) []string {
	for i := 0; i < len(words); i++ {
		if words[i] == "(up," {
			nstr := words[i+1][:len(words[i+1])-1]
			n, err := strconv.Atoi(nstr)
			checkErr(err)
			if i > n {
				for j := i - n; j < i; j++ {
					words[j] = strings.ToUpper(words[j])
				}
			}
			words = append(words[:i], words[i+2:]...)

		}
	}

	return words
}

func lowercaseNWords(words []string) []string {
	for i := 0; i < len(words); i++ {
		if words[i] == "(low," {
			nstr := words[i+1][:len(words[i+1])-1]
			n, err := strconv.Atoi(nstr)
			checkErr(err)
			if i > n {
				for j := i - n; j < i; j++ {
					words[j] = strings.ToLower(words[j])
				}
			}
			words = append(words[:i], words[i+2:]...)

		}
	}

	return words
}

func Punctuating(sentence string) string {
	re := regexp.MustCompile(`(\s)([.,!?:;])(\S)`)
	sentence = re.ReplaceAllString(sentence, "$2 $3")
	re = regexp.MustCompile(`(\S)(\s)([.,!?:;])`)
	sentence = re.ReplaceAllString(sentence, "$1$3")

	return sentence
}

func articles(sentence string) string {
	words := strings.Fields(sentence)
	vowels := "aeiouhAEIOUH"
	for i, word := range words {
		if word == "a" || word == "A" {
			if strings.Contains(vowels, string(words[i+1][0])) {
				if word == "A" {
					words[i] = "An"
				} else if word == "a" {
					words[i] = "an"
				}
			}
		}
	}
	sentence = strings.Join(words, " ")
	return sentence
}

func main() {
	if len(os.Args) == 3 {
		if os.Args[1] == "sample.txt" && os.Args[2] == "result.txt" {
			inputFileContent, err := os.ReadFile(os.Args[1])
			checkErr(err)
			input := string(inputFileContent)
			indWords := strings.Fields(input)
			indWords = hexToDec(indWords)
			indWords = binToDec(indWords)
			indWords = uppercase(indWords)
			indWords = lowercase(indWords)
			indWords = capitalization(indWords)
			indWords = capapitalizeNWords(indWords)
			indWords = uppercaseNWords(indWords)
			indWords = lowercaseNWords(indWords)
			modifiedInput := strings.Join(indWords, " ")
			modifiedInput = Punctuating(modifiedInput)
			modifiedInput = articles(modifiedInput)

			output := []byte(modifiedInput)
			os.WriteFile(os.Args[2], output, 0o666)

		} else {
			fmt.Println("Incorrect input and output file")
		}
	}
}

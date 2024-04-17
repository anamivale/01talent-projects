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

func Marks(input string) string {
	re := regexp.MustCompile(`(')(\s)`)
	formattedText := re.ReplaceAllString(input, "$1")
	re = regexp.MustCompile(`(\s)('$)`)
	formattedText = re.ReplaceAllString(formattedText, "$2")

	return formattedText
}

// func hexdec()

func main() {
	if len(os.Args) > 1 && len(os.Args) <= 3 {
		file1 := os.Args[1]
		// file2 := os.Args[2]
		data1, err := os.ReadFile(file1)
		checkErr(err)
		var newSentence []string
		words := strings.Fields(string(data1))
		for i, word := range words {
			// vowels and consonants
			if word == "a" || word == "an" {
				if strings.ToLower(string(words[i+1][0])) == "a" || strings.ToLower(string(words[i+1][0])) == "e" || strings.ToLower(string(words[i+1][0])) == "i" || strings.ToLower(string(words[i+1][0])) == "o" || strings.ToLower(string(words[i+1][0])) == "u" || strings.ToLower(string(words[i+1][0])) == "h" {
					words[i] = "an"
				} else {
					words[i] = "a"
				}
			}
			if word == "A" || word == "An" {
				if strings.ToLower(string(words[i+1][0])) == "a" || strings.ToLower(string(words[i+1][0])) == "e" || strings.ToLower(string(words[i+1][0])) == "i" || strings.ToLower(string(words[i+1][0])) == "o" || strings.ToLower(string(words[i+1][0])) == "u" || strings.ToLower(string(words[i+1][0])) == "h" {
					words[i] = "An"
				} else {
					words[i] = "A"
				}
			}

			// (cap)
			if word == "(cap)" {
				words[i-1] = strings.ToUpper(string(words[i-1][0])) + strings.ToLower(words[i-1][1:])
			}
			// hextodec
			if word == "(hex)" {
				dec, err := strconv.ParseInt(words[i-1], 16, 64)
				checkErr(err)
				words[i-1] = strconv.Itoa(int(dec))
			}
			// bintodec
			if word == "(bin)" {
				bin, err := strconv.ParseInt(words[i-1], 2, 64)
				checkErr(err)
				words[i-1] = strconv.Itoa(int(bin))
			}
			// (up)
			if word == "(up)" {
				words[i-1] = strings.ToUpper(words[i-1])
			}
			// (low)
			if word == "(low)" {
				words[i-1] = strings.ToLower(words[i-1])
			}

			//(cap, n)
			if strings.HasPrefix(word, "(cap,") {
				nStr := words[i+1][0 : len(words[i+1])-1]
				n, err := strconv.Atoi(nStr)
				checkErr(err)
				if i >= n {
					for j := i - n; j < i; j++ {
						words[j] = strings.ToUpper(string(words[j][0])) + strings.ToLower(words[j][1:])
						continue
					}
				}
			}

			// (up, n)
			if strings.HasPrefix(word, "(up,") {

				nStr := words[i+1][0 : len(words[i+1])-1]
				n, err := strconv.Atoi(nStr)
				checkErr(err)

				if i >= n {
					for j := i - n; j < i; j++ {
						words[j] = strings.ToUpper(words[j])
						continue
					}
				}
			}
			// (low, n)
			if strings.HasPrefix(word, "(low,") {
				nStr := words[i+1][0 : len(words[i+1])-1]
				n, err := strconv.Atoi(nStr)
				checkErr(err)
				if i >= n {
					for j := i - n; j < i; j++ {
						words[j] = strings.ToLower(words[j])
						continue
					}
				}
			}

		}
		for i, word := range words {
			if word == "(cap)" || word == "(up)" || word == "(low)" || word == "(hex)" || word == "(bin)" {
				continue
			}
			if strings.HasPrefix(word, "(cap,") || strings.HasPrefix(word, "(up,") || strings.HasPrefix(word, "(low,") {
				words[i+1] = ""

				continue
			}
			newSentence = append(newSentence, word)
		}

		// puct ., ,, !, ?, : and ;
		x := strings.Join(newSentence, " ")
		y := strings.Fields(x)
		puctSentence := strings.Join(y, " ")
		var result []rune
		for i, letter := range puctSentence {
			if letter == ' ' && (puctSentence[i+1] == '.' || puctSentence[i+1] == ',' || puctSentence[i+1] == '!' || puctSentence[i+1] == '?' || puctSentence[i+1] == ':' || puctSentence[i+1] == ';') {
				continue
			} else if letter == '.' || letter == ',' || letter == '!' || letter == '?' || letter == ':' || letter == ';' {
				result = append(result, letter)
			} else {
				result = append(result, letter)
			}
		}
		var finalRest []rune

		for i, pnct := range result {
			if i > 0 {
				if (result[i-1] == '.' || result[i-1] == ',' || result[i-1] == '!' || result[i-1] == '?' || result[i-1] == ':' || result[i-1] == ';') && !(pnct == '.' || pnct == ',' || pnct == '!' || pnct == '?' || pnct == ':' || pnct == ';' || pnct == ' ') {
					finalRest = append(finalRest, ' ')
				}
			}
			finalRest = append(finalRest, pnct)

		}
		count := 0
		for _, mark := range finalRest {
			if mark == '\'' {
				count++
			}
		}
		var conText string
		if count == 2 {
			conText = Marks(string(finalRest))
		} else {
			conText = string(finalRest)
		}

		modifiedText := []byte(conText)
		modifiedText = append(modifiedText, '\n')
		err = os.WriteFile(os.Args[2], modifiedText, 0o644)
		checkErr(err)

	} else {
		fmt.Println("not done")
	}
}

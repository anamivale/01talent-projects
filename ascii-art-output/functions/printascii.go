package functions

import (
	"strings"
)

func Mapping(fileContent []string) map[rune][]string {
	mapBanner := make(map[rune][]string)
	runs := 32
	for i, str := range fileContent {
		if str == "" {
			if i == 0 {
				continue
			}
			runs++
			continue
		}
		mapBanner[rune(runs)] = append(mapBanner[rune(runs)], str)

	}
	return mapBanner
}

func PrintAscii(mapBanner map[rune][]string, userInput string) string {
	var outputStr strings.Builder
	userInput = strings.ReplaceAll(userInput, "\n", "\\n")
	// userInput = strings.ReplaceAll(userInput, "\t", "\\n")
	validStr := strings.Split(userInput, "\\n")

	for i, str := range validStr {
		if str == "" {
			if i==0{
				continue
			}
			outputStr.WriteString("\n")
		} else {
			 output := make([]string, 8)
			for _, char := range str {
				if value, ok := mapBanner[char]; ok{
				for i:= 0; i<8; i++ {
					output[i] += value[i]
				}

				} else {
					return "!ok"
				}
			}
			for i:= 0; i<8; i++ {
				outputStr.WriteString(output[i]+"\n")
			}
			

		}
	}

return outputStr.String()

}
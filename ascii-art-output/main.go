package main

import (
	"flag"
	"fmt"
	"os"
	fileProcessing "output/functions"
)

func main() {
	var outputFile string
	flag.StringVar(&outputFile, "output", "banner.txt", "banner file")
	flag.Parse()

	file := ""
	inputStr := os.Args[2]

	if !(len(os.Args) == 2 || len(os.Args) ==4 ){

		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
		return
	}
	if len(os.Args) == 2{
		file = "standard.txt"
		inputStr= os.Args[1]
	}
	
	file = os.Args[3]+".txt"
	x := fileProcessing.ValidateBanner(file)
	if x == nil {
		fmt.Println("file not valid or empty")
		return
	}
	y := fileProcessing.Mapping(x)
	z := fileProcessing.PrintAscii(y, inputStr)
	if z == "!ok"{
		fmt.Println("character not within the range (32-126)")
	}

	err := os.WriteFile(outputFile, []byte(z), 0o444)
	if err != nil {
		fmt.Println("unable to write to the output file")
		return
	}

}

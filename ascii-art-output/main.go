package main

import (
	"flag"
	"fmt"
	"os"
	file "output/functions"
)

func main() {
	var outputFile string
	flag.StringVar(&outputFile, "output", "standard.txt", "banner file")
	flag.Parse()
	fmt.Println(outputFile)
	if len(os.Args) != 2 || len(os.Args) !=4 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
	}
	x := file.ValidateBanner("standard.txt")
	if x == nil {
		fmt.Println("file not valid or empty")
		return
	}
	y := file.Mapping(x)
	z := file.PrintAscii(y, "val")
	if z == "!ok"{
		fmt.Println("character not within the range (32-126)")
	}

	fmt.Print(z)
}

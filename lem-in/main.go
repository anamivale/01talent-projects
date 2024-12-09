package main

import (
	"fmt"
	"os"

	"learn.zone01kisumu.ke/git/vmuhembe/lem-in/utils"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("only requires one argument")
		return
	}
	fileName := os.Args[1]

	content := utils.ReadFile(fileName)

	filledContents, err := utils.GroupContent(content)
	if err != nil {
		fmt.Println(err.Error())
	}
	x := utils.Searchpath(filledContents.Links, "start", "end")

	
	y := utils.RemoveInterPaths(x)
	fmt.Println(utils.FindPathsToUse(y, 10))

	for _, v := range y {
		fmt.Println(len(v), v)
	}
}

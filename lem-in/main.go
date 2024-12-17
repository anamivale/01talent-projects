package main

import (
	"fmt"

	"learn.zone01kisumu.ke/git/vmuhembe/lem-in/structs"
	"learn.zone01kisumu.ke/git/vmuhembe/lem-in/utils"
)

func main() {
	filename, errmsg := utils.ParseArgs()
	if errmsg != "" {
		fmt.Println(errmsg)
		return
	}
	Antcolony, err := utils.ParseFile(filename)
	if err != nil {
		fmt.Println("ERROR: invalid data format,", err)
		return
	}
	paths, antsperpath, turns := utils.FindPaths(Antcolony)
	moves := utils.MoveAnts(paths, antsperpath, turns)

	// Print the file contents
	fmt.Println(structs.FileContents)

	for _, move := range moves {
		fmt.Println(move)
	}
}

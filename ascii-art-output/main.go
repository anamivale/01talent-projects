package main

import (
	"fmt"
	file "output/functions"
)

func main() {
	x := file.ValidateBanner("shadow.txt")
	if x == nil {
		fmt.Println("file not valid or empty")
		return
	}
	fmt.Println(x[20])
}

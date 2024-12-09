package utils

import (
	"fmt"
)

func FindPathsToUse(paths [][]string, ants int) []int {
	var ls []int
	ls = append(ls, 0)
	if len(paths) == 1 {
		return ls
	}
	turns := len(paths[0]) + 1
	for i := 1; i < ants; i++ {
		antsPresent := countChar(ls, ls[len(ls)-1])
		turns = len(paths[ls[len(ls)-1]]) + antsPresent
		for j := 0; j < len(paths); j++ {
			tempRuns := len(paths[j])
			if !(turns > tempRuns ){

				ls = append(ls, j)
				fmt.Println(turns, tempRuns)

				break
			}

		}
	}
	fmt.Println(ls)
	return ls
}

func countChar(s []int, c int) int {
	count := 0
	for _, v := range s {
		if v == c {
			count++
		}
	}
	return count
}

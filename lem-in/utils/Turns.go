package utils

import "learn.zone01kisumu.ke/git/vmuhembe/lem-in/structs"

// finding maximum turns the paths to be optimized can have
func GenerateTurns(option map[int][]int, paths []structs.Path) int {
	maxturns := 0
	for i, path := range paths {
		rooms := len(path.Rooms) - 1
		ants := len(option[i])
		turns := rooms + ants - 1
		if turns > maxturns {
			maxturns = turns
		}
	}
	return maxturns
}

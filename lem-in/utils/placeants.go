package utils

import "learn.zone01kisumu.ke/git/vmuhembe/lem-in/structs"

//  function that assigns rooms to ants and places them
func PlaceAnts(colony *structs.AntColony, paths []structs.Path) map[int][]int {
	ants := colony.Ants
	pathants := make(map[int][]int)
	ant := 1
	for ants > 0 {
		if PlaceRecursively(ant, paths, pathants, 0) {
			ant++
			ants--
		}
	}
	return pathants
}

// helper function to placeants that places recursively until an optimal solution is found
func PlaceRecursively(ant int, paths []structs.Path, pathants map[int][]int, path int) bool {
	if ant == 1 || path == len(paths)-1 {
		pathants[path] = append(pathants[path], ant)
		return true
	}
	rooms := len(paths[path].Rooms) - 2
	antsinroom := len(pathants[path]) - 2
	det := rooms + antsinroom
	if det > (len(paths[path+1].Rooms)-2)+(len(pathants[path+1])-2) {
		if PlaceRecursively(ant, paths, pathants, path+1) {
			return true
		}
	} else {
		pathants[path] = append(pathants[path], ant)
		return true
	}
	return false
}

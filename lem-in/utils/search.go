package utils

import (
	"strings"
)

func Searchpath(edges []string, start, stop string) [][]string {
	graph := make(map[string][]string)

	// Build adjacency list from edges
	for _, edge := range edges {
		parts := strings.Split(edge, "-")
		if len(parts) != 2 {
			continue
		}
		graph[parts[0]] = append(graph[parts[0]], parts[1])
		graph[parts[1]] = append(graph[parts[1]], parts[0]) 
	}

	// BFS initialization
	queue := [][]string{{start}}
	var allPaths [][]string

	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]
		node := path[len(path)-1]

		// If we've reached the destination, save the path
		if node == stop {
			allPaths = append(allPaths, path)
			continue
		}

		// Add neighbors to the queue
		for _, neighbor := range graph[node] {
			if !contains(path, neighbor) { // Avoid cycles
				newPath := append([]string{}, path...) // Copy the path
				newPath = append(newPath, neighbor)
				queue = append(queue, newPath)
			}
		}
	}

	return allPaths
}

// Helper function to check if a slice contains a string
func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

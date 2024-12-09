package main

func RemoveInterPaths(paths [][]string) [][]string {
	var result [][]string
	visited := make([]bool, len(paths)) // Track visited paths

	for i := 0; i < len(paths); i++ {
		if visited[i] {
			continue // Skip already processed paths
		}

		result = append(result, paths[i]) // Retain this path
		visited[i] = true

		// Mark all interlocked paths as visited
		for j := i + 1; j < len(paths); j++ {
			if CheckInterPaths(paths[i], paths[j]) {
				visited[j] = true
			}
		}
	}

	return result
}

func CheckInterPaths(path1, path2 []string) bool {
	nodeSet := make(map[string]bool)

	for i := 1; i < len(path1)-1; i++ {
		nodeSet[path1[i]] = true
	}

	for j := 1; j < len(path2)-1; j++ {
		if nodeSet[path2[j]] {
			return true
		}
	}

	return false
}

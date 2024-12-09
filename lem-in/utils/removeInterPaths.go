package utils

func RemoveInterPaths(paths [][]string) [][]string {
	var result [][]string
	visited := make([]bool, len(paths)) 

	for i := 0; i < len(paths); i++ {
		if visited[i] {
			continue 
		}

		result = append(result, paths[i]) 
		visited[i] = true

		for j := i + 1; j < len(paths); j++ {
			if CheckInterPaths(paths[i], paths[j]) {
				visited[j] = true
			}
		}
	}

	return result
}

func CheckInterPaths(path1, path2 []string) bool {
	l := 0

	if len(path1) > len(path2) {
		l = len(path2)
	} else {
		l = len(path1)
	}
	for i := 1; i < l-1; i++ {
		for j := 1; j < l-1; j++ {
			if path1[i] == path2[j] {
				return true
			}
		}
	}

	return false
}

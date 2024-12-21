package Mosdef

func BreadthFirstSearch(graph map[string][]string, start, end string) [][]string {
	queue := [][]string{{start}} // Queue of paths
	visited := map[string]bool{start: true}
	shortestPaths := [][]string{}
	shortestLength := -1

	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]
		current := path[len(path)-1]

		// Stop if the path length exceeds the shortest found
		// if shortestLength != -1 && len(path) > shortestLength {
		// 	break
		// }

		if current == end {
			if shortestLength == -1 {
				shortestLength = len(path)
			}
			shortestPaths = append(shortestPaths, path)
			continue
		}

		for _, neighbor := range graph[current] {
			if !visited[neighbor] || neighbor == end {
				newPath := append([]string{}, path...)
				newPath = append(newPath, neighbor)
				queue = append(queue, newPath)
				visited[neighbor] = true
			}
		}
	}

	return shortestPaths
}

package main

// GetExtendedNeighbors takes in a pattern (read), the overlap graph and maxK.
// It returns the extendedNeighbors list. For each neighbor *n* in this list,
// distance between n and pattern is between 2 to maxK.
func GetExtendedNeighbors(pattern string, adjList map[string][]string, maxK int) []string {
	curNodes := adjList[pattern]
	visited := map[string]bool{}
	visited[pattern] = true

	for d := 2; d <= maxK; d++ {
		curNodes = ExpandNodes(curNodes, adjList, visited)
	}

	finalList := []string{}
	for node := range visited {
		if node != pattern {
			finalList = append(finalList, node)
		}
	}

	return finalList
}

func ExpandNodes(curNodes []string, adjList map[string][]string, visited map[string]bool) []string {
	nexNodes := []string{}
	for _, node := range curNodes {
		nexNodes = ExpandNode(node, adjList, nexNodes, visited)
	}
	return nexNodes
}

func ExpandNode(node string, adjList map[string][]string, nexNodes []string, visited map[string]bool) []string {
	for _, neighbor := range adjList[node] {
		if !visited[neighbor] {
			nexNodes = append(nexNodes, neighbor)
			visited[neighbor] = true
		}
	}

	return nexNodes
}

package main

// AverageOutDegree takes the adjacency list of a directed network.
// It returns the average outdegree of each node in the network.
// It does not include nodes with no outgoing edges in the average.
func AverageOutDegree(adjList map[string][]string) float64 {
	total := 0.0
	count := 0

	for _, adj := range adjList {
		if len(adj) > 0 {
			total += float64(len(adj))
			count++
		}
	}

	if count == 0.0 {
		return 0.0
	}

	return total / float64(count)
}

// AverageOutDegreeAllNodes takes as input the adjacency list of a directed network.
// It returns the average outdegree of each node in the network, including nodes with outdegree zero.
func AverageOutDegreeAllNodes(adjList map[string][]string) float64 {
	total := 0.0

	for _, adj := range adjList {
		total += float64(len(adj))
	}

	if len(adjList) == 0.0 {
		return 0.0
	}

	return total / float64(len(adjList))
}

package main

import (
	"slices"
)

// GetTrimmedNeighbors takes in a string pattern (read), an adjacency list and maxK.
// It returns all n-transitivity edges in the adjList of the current read (pattern) for n <= maxK.
func GetTrimmedNeighbors(pattern string, adjList map[string][]string, maxK int) []string {
	neighbors := adjList[pattern]
	extNeighbors := GetExtendedNeighbors(pattern, adjList, maxK)
	for i := 0; i < len(neighbors); i++ {
		if slices.Contains(extNeighbors, neighbors[i]) {
			neighbors = append(neighbors[:i], neighbors[i+1:]...)
		}
	}
	return neighbors
}

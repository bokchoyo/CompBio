package main

import "math/rand"

// ShuffleStrings takes a collection of strings patterns as input.
// It returns a random shuffle of the strings.
func ShuffleStrings(patterns []string) []string {
	array := make([]string, len(patterns))

	indices := rand.Perm(len(patterns))

	for i, index := range indices {
		array[i] = patterns[index]
	}

	return array
}

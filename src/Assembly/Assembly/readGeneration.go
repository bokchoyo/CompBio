package main

import (
	"math/rand"
)

// SimulateReadsClean takes a genome along with a read length and a probability.
// It returns a collection of strings resulting from simulating clean reads,
// where a given position is sampled with the given probability.
func SimulateReadsClean(genome string, readLength int, probability float64) []string {
	result := []string{}

	for i := 0; i < len(genome)-readLength+1; i++ {
		if rand.Float64() < probability {
			result = append(result, genome[i:i+readLength])
		}
	}

	return result
}

package main

import "math/rand"

// GenerateRandomGenome takes a parameter length and returns
// a random DNA string of this length where each nucleotide has equal probability.
func GenerateRandomGenome(length int) string {

	genome := ""

	for i := 0; i < length; i++ {
		randomizer := rand.Intn(4)
		if randomizer == 0 {
			genome += "A"
		} else if randomizer == 1 {
			genome += "C"
		} else if randomizer == 2 {
			genome += "T"
		} else {
			genome += "G"
		}
	}

	return genome
}

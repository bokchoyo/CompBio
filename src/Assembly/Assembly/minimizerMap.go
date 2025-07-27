package main

import "slices"

// StringIndex is a type that will map a minimizer string to its list of indices
// in a read set corresponding to reads with this minimizer.
type StringIndex map[string][]int

// BuildMinimizerMap takes a collection of reads, integer k and integer windowLength.
// It returns a map of k-mers to the indices of the reads in the list having this k-mer minimizer.
func BuildMinimizerMap(reads []string, k int, windowLength int) StringIndex {
	dict := make(StringIndex)

	for readIndex, read := range reads {
		curMin, curMinI := MinimizerAndIndex(read[:windowLength], k)
		for i := 1; i <= len(read)-windowLength; i++ {
			if curMinI < i {
				winMin, winMinI := MinimizerAndIndex(read[i:i+windowLength], k)
				curMin = winMin
				curMinI = i + winMinI
				if !slices.Contains(dict[winMin], readIndex) {
					dict[winMin] = append(dict[winMin], readIndex)
				}
			} else {
				newKmerI := i + windowLength - k
				if newKmerI+k > len(read) {
					continue
				}
				newKmer := read[newKmerI : newKmerI+k]
				if newKmer < curMin {
					curMin = newKmer
					curMinI = newKmerI
					if !slices.Contains(dict[newKmer], readIndex) {
						dict[newKmer] = append(dict[newKmer], readIndex)
					}
				}
			}
		}
	}

	return dict
}

func MinimizerAndIndex(text string, k int) (string, int) {
	n := len(text)
	minimizerIndex := 0
	minimizer := text[0:k]

	for i := 0; i <= n-k; i++ {
		if minimizer > text[i:i+k] {
			minimizer = text[i : i+k]
			minimizerIndex = i
		}
	}

	return minimizer, minimizerIndex
}

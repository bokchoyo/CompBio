package main

// SumPairsScore takes two multiple alignments as well as two indices, and scoring
// parameters. It returns the sum of pairs score of the corresponding columns
// in the two alignments, using the specified scoring parameters.
func SumPairsScore(align1 Alignment, align2 Alignment,
	idx1 int, idx2 int, match float64, mismatch float64, gap float64) float64 {
	score := 0.0

	for _, str1 := range align1 {
		char1 := str1[idx1 : idx1+1]
		for _, str2 := range align2 {
			char2 := str2[idx2 : idx2+1]
			if char1 == "-" && char2 == "-" {
				continue
			} else if char1 == char2 {
				score += match
			} else if char1 == "-" || char2 == "-" {
				score -= gap
			} else {
				score -= mismatch
			}
		}
	}

	return score
}

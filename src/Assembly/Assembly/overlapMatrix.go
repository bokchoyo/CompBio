package main

// OverlapScoringMatrix takes a collection of reads along with alignment penalties.
// It returns a matrix in which mtx[i][j] is the overlap alignment score of
// reads[i] with reads[j].
func OverlapScoringMatrix(reads []string, match, mismatch, gap float64) [][]float64 {
	matrix := make([][]float64, len(reads))
	for i := range matrix {
		matrix[i] = make([]float64, len(reads))
	}
	for r := 0; r < len(reads); r++ {
		for c := 0; c < len(reads); c++ {
			if r != c {
				matrix[r][c] = ScoreOverlapAlignment(reads[r], reads[c], match, mismatch, gap)
			}
		}

	}
	return matrix
}

package main

// EditDistanceMatrix takes as input a slice of strings patterns.
// It returns a matrix whose (i,j)th entry is the edit distance between
// the i-th and j-th strings in patterns.
func EditDistanceMatrix(patterns []string) [][]int {
	matrix := make([][]int, len(patterns))
	for i := range matrix {
		matrix[i] = make([]int, len(patterns))
	}

	for r, pattern := range patterns {
		for c := r + 1; c < len(patterns); c++ {
			d := EditDistance(pattern, patterns[c])
			matrix[r][c] = d
			matrix[c][r] = d
		}
	}

	return matrix
}

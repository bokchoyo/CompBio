package main

// LocalScoreTable takes two strings and alignment penalties. It returns a 2-D array
// holding dynamic programming scores for local alignment with these penalties.
func LocalScoreTable(str1, str2 string, match, mismatch, gap float64) [][]float64 {
	array := make([][]float64, len(str1)+1)

	for i := range array {
		array[i] = make([]float64, len(str2)+1)
	}

	for r := 1; r <= len(str1); r++ {
		for c := 1; c <= len(str2); c++ {
			if str1[r-1] == str2[c-1] {
				array[r][c] = MaxFloat(0, array[r-1][c-1]+match, array[r][c-1]-gap, array[r-1][c]-gap)
			} else {
				array[r][c] = MaxFloat(0, array[r-1][c-1]-mismatch, array[r][c-1]-gap, array[r-1][c]-gap)
			}
		}
	}

	return array
}

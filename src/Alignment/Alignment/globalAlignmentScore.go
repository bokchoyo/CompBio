package main

// GlobalScoreTable takes two strings and alignment penalties. It returns a 2-D array
// holding dynamic programming scores for global alignment with these penalties.
func GlobalScoreTable(str1, str2 string, match, mismatch, gap float64) [][]float64 {
	array := make([][]float64, len(str1)+1)

	for i := range array {
		array[i] = make([]float64, len(str2)+1)
	}

	for i := 0; i <= len(str1); i++ {
		array[i][0] -= float64(i) * gap
	}

	for i := 0; i <= len(str2); i++ {
		array[0][i] -= float64(i) * gap
	}

	for r := 1; r <= len(str1); r++ {
		for c := 1; c <= len(str2); c++ {
			if str1[r-1] == str2[c-1] {
				array[r][c] = MaxFloat(array[r-1][c-1]+match, array[r][c-1]-gap, array[r-1][c]-gap)
			} else {
				array[r][c] = MaxFloat(array[r-1][c-1]-mismatch, array[r][c-1]-gap, array[r-1][c]-gap)
			}
		}

	}

	return array
}

func MaxFloat(nums ...float64) float64 {
	m := 0.0

	for i, val := range nums {
		if i == 0 || val > m {
			m = val
		}
	}

	return m
}

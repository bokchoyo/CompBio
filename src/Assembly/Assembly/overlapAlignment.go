package main

//ALL PENALTIES POSITIVE

// ScoreOverlapAlignment takes two strings along with match, mismatch, and gap penalties.
// It returns the maximum score of an overlap alignment in which str1 is overlapped with str2.
// Assume we are overlapping a suffix of str1 with a prefix of str2.
func ScoreOverlapAlignment(str1, str2 string, match, mismatch, gap float64) float64 {
	array := make([][]float64, len(str1)+1)
	max := 0.0

	for i := range array {
		array[i] = make([]float64, len(str2)+1)
	}

	for r := 1; r <= len(str1); r++ {
		for c := 1; c <= len(str2); c++ {
			if str1[r-1] == str2[c-1] {
				array[r][c] = MaxFloat(array[r-1][c-1]+match, array[r][c-1]-gap, array[r-1][c]-gap)
			} else {
				array[r][c] = MaxFloat(array[r-1][c-1]-mismatch, array[r][c-1]-gap, array[r-1][c]-gap)
			}

			if r == len(str1) && array[r][c] > max {
				max = array[r][c]
			}
		}
	}

	return max
}

package main

// BinarizeMatrix takes a matrix of values and a threshold.
// It binarizes the matrix according to the threshold.
// If entries across main diagonal are both above threshold, only retain the bigger one.
func BinarizeMatrix(mtx [][]float64, threshold float64) [][]int {
	binMtx := make([][]int, len(mtx))

	for i := range mtx {
		binMtx[i] = make([]int, len(mtx[0]))
	}

	for r := 0; r < len(mtx); r++ {
		for c := r + 1; c < len(mtx[0]); c++ {
			if mtx[r][c] >= threshold && mtx[c][r] >= threshold {
				if mtx[r][c] >= mtx[c][r] {
					binMtx[r][c] = 1
				} else if mtx[c][r] > mtx[r][c] {
					binMtx[c][r] = 1
				}
			} else if mtx[r][c] >= threshold {
				binMtx[r][c] = 1
			} else if mtx[c][r] >= threshold {
				binMtx[c][r] = 1
			}
		}
	}

	return binMtx
}

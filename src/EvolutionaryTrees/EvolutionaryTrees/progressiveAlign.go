package main

// ProgressiveAlign takes two (multiple) alignments as input and
// returns a multiple alignment corresponding to combining the two
// alignments according to the Clustal dynamic programming heuristic.

func ProgressiveAlign(align1 Alignment, align2 Alignment,
	match float64, mismatch float64, gap float64, supergap float64) Alignment {
	mtx := ProgressiveAlignmentScoreTable(align1, align2, match, mismatch, gap, supergap)
	a := make(Alignment, len(align1)+len(align2))
	r, c := len(align1[0]), len(align2[0])
	numStrings1, numStrings2 := len(align1), len(align2)
	alignment := make(Alignment, numStrings1 + numStrings2)
	allRows := make([][]byte, numStrings1 + numStrings2)
	for i := range allRows {
		allRows[i] = make([]byte, 0)
	}
	

	for r > 0 && c > 0 {
		if mtx[r][c] == mtx[r][c-1] - supergap {
			for i := 0; i < numStrings1; i++ {
				allRows[i] = append([]byte{'-'}, allRows[i]...)
			}
			for j := 0; j < numStrings2; j++ {
				allRows[numStrings1 + j] = append([]byte{align2[j][c-1]}, allRows[numStrings1 + j]...)
			}
			c--
		} else if mtx[r][c] == mtx[r-1][c] - supergap {
			for i := 0; i < numStrings1; i++ {
				allRows[numStrings2 + i] = append([]byte{align1[i][c-1]}, allRows[numStrings2 + i]...)
			}
			for j := 0; j < numStrings2; j++ {
				allRows[j] = append([]byte{'-'}, allRows[j]...)
			}
			c--
		} else {

		}
	}

	for r > 0 {

	}

	for c > 0 {

	}

	for r > 0 || c > 0 {
		if r == 0 {
			for i := range align1 {
				a[i] = "-" + a[i]
			}
			for i := range align2 {
				a[i+len(align1)] = string(align2[i][c-1]) + a[i+len(align1)]
			}
			c--
		} else if c == 0 {
			for i := range align2 {
				a[i] = "-" + a[i]
			}

			for i := range align1 {
				a[i] = string(align1[i][c-1]) + a[i]
			}
			r--
		} else {
			n := mtx[r][c]
			switch n {
			case mtx[r][c-1] - supergap:
				for i := range align1 {
					a[i] = "-" + a[i]
				}

				for i := range align2 {
					a[i+len(align1)] = string(align2[i][c-1]) + a[i+len(align1)]
				}
				c--
			case mtx[r-1][c] - supergap:
				for i := range align2 {
					a[i+len(align1)] = "-" + a[i+len(align1)]
				}

				for i := range align1 {
					a[i] = string(align1[i][r-1]) + a[i]
				}
				r--
			case mtx[r-1][c-1] + SumPairsScore(align1, align2, r, c, match, mismatch, gap):
				for i := range align1 {
					a[i] = string(align1[i][r-1]) + a[i]
				}

				for i := range align2 {
					a[i+len(align1)] = string(align2[i][c-1]) + a[i+len(align1)]
				}
				r--
				c--
			default:
				panic("unexpected result")
			}
		}
	}
	return a
}

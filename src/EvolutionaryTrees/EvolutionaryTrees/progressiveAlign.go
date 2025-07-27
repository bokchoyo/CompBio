package main

import "math"

// ProgressiveAlign takes two (multiple) alignments as input and
// returns a multiple alignment corresponding to combining the two
// alignments according to the Clustal dynamic programming heuristic.

func ProgressiveAlign(align1 Alignment, align2 Alignment,
	match float64, mismatch float64, gap float64, supergap float64) Alignment {

	a := make(Alignment, len(align1)+len(align2))
	mtx := ProgressiveAlignmentScoreTable(align1, align2, match, mismatch, gap, supergap)
	r := len(align1[0])
	c := len(align2[0])

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
				a[i+len(align1)] = "-" + a[i+len(align1)]
			}
			for i := range align1 {
				a[i] = string(align1[i][r-1]) + a[i]
			}
			r--
		} else {
			n := mtx[r][c]
			if math.Abs(n-(mtx[r][c-1]-supergap)) < 1e-6 {
				for i := range align1 {
					a[i] = "-" + a[i]
				}
				for i := range align2 {
					a[i+len(align1)] = string(align2[i][c-1]) + a[i+len(align1)]
				}
				c--
			} else if math.Abs(n-(mtx[r-1][c]-supergap)) < 1e-6 {
				for i := range align2 {
					a[i+len(align1)] = "-" + a[i+len(align1)]
				}
				for i := range align1 {
					a[i] = string(align1[i][r-1]) + a[i]
				}
				r--
			} else {
				for i := range align1 {
					a[i] = string(align1[i][r-1]) + a[i]
				}
				for i := range align2 {
					a[i+len(align1)] = string(align2[i][c-1]) + a[i+len(align1)]
				}
				r--
				c--
			}
		}
	}
	return a
}

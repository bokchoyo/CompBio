package main

//ProgressiveAlignmentScoreTable takes two multiple alignments as well as a collection
//of Clustal scoring parameters. It returns a 2D matrix corresponding to
//the Clustal dynamic programming table for combining the two alignments heuristically
//into a single multiple alignment.
func ProgressiveAlignmentScoreTable(align1 Alignment, align2 Alignment,
	match float64, mismatch float64, gap float64, supergap float64) [][]float64 {
	
	mtx := make([][]float64, len(align1)+1)

	for i := range mtx {
		mtx[i] = make([]float64, len(align2)+1)
	}

    for i := 0; i <= len(align1[0]); i++ {
        mtx[i][0] -= float64(i) * supergap
    }

    for i := 0; i <= len(align2[0]); i++ {
        mtx[0][i] -= float64(i) * supergap
    }

	for r := 1; r <= len(align1[0]); r++ {
		for c := 1; c <= len(align2[0]); c++ {
			mtx[r][c] = MaxFloat(mtx[r-1][c-1] + SumPairsScore(align1, align2, r-1, c-1, match, mismatch, gap), mtx[r-1][c] - supergap, mtx[r][c-1] - supergap)
		}
	}

	return mtx
}
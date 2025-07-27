package main

import (
	"math"
	"strings"
)

// ProgressiveOverlapAlign aligns a new read against an existing alignment
// using the progressive alignment method.
// It performs an overlap alignment, in which we overlap a suffix of the existing alignment
// with a prefix of the new read, and we take the alignment achieving the maximum score
// over all such alignments.
// It computes the score of the alignment using the provided scoring parameters, where
// we score the current symbol against all elements of the indicated column in the alignment.
// match pairs receive the match reward; mismatch pairs receive the mismatch penalty;
// gaps receive the gap penalty.
func ProgressiveOverlapAlign(a Alignment, read string, match, mismatch, gap float64) Alignment {
	window := len(read)
	frozen, aSuf := SplitFrozenPrefix(a, window)
	profile := MakeProfile(aSuf)
	scoringMatrix := OverlapProfileScoringMatrix(profile, read, match, mismatch, gap)
	maxScore := float64(math.MinInt64)
	maxCol := -1

	for j := 1; j < len(scoringMatrix[0]); j++ {
		if scoringMatrix[len(scoringMatrix)-1][j] > maxScore {
			maxScore = scoringMatrix[len(scoringMatrix)-1][j]
			maxCol = j
		}
	}

	alignBytes := make([][]byte, len(aSuf)+1)

	for i := range alignBytes {
		alignBytes[i] = nil
	}

	row := len(scoringMatrix) - 1
	col := maxCol

	for row > 0 || col > 0 {
		if row == 0 {
			alignBytes[len(aSuf)] = prepend(alignBytes[len(aSuf)], read[col-1])
			for i := range aSuf {
				alignBytes[i] = prepend(alignBytes[i], '-')
			}
			col--
			continue
		}
		if col == 0 {
			for i := range aSuf {
				alignBytes[i] = prepend(alignBytes[i], aSuf[i][row-1])
			}
			alignBytes[len(aSuf)] = prepend(alignBytes[len(aSuf)], '-')
			row--
			continue
		}

		up := scoringMatrix[row-1][col] -
			gap*float64(NumberofNonGapSymbols(profile[row-1]))
		left := scoringMatrix[row][col-1] -
			gap*float64(NumberofNonGapSymbols(profile[row-1]))

		switch {
		case scoringMatrix[row][col] == up:
			for i := range aSuf {
				alignBytes[i] = prepend(alignBytes[i], aSuf[i][row-1])
			}
			alignBytes[len(aSuf)] = prepend(alignBytes[len(aSuf)], '-')
			row--
		case scoringMatrix[row][col] == left:
			for i := range aSuf {
				alignBytes[i] = prepend(alignBytes[i], '-')
			}
			alignBytes[len(aSuf)] = prepend(alignBytes[len(aSuf)], read[col-1])
			col--
		default:
			for i := range aSuf {
				alignBytes[i] = prepend(alignBytes[i], aSuf[i][row-1])
			}
			alignBytes[len(aSuf)] = prepend(alignBytes[len(aSuf)], read[col-1])
			row--
			col--
		}
	}

	for j := maxCol; j < len(read); j++ {
		alignBytes[len(aSuf)] = append(alignBytes[len(aSuf)], read[j])
		for i := range aSuf {
			alignBytes[i] = append(alignBytes[i], '-')
		}
	}

	suffixAligned := make(Alignment, len(alignBytes))
	for i := range alignBytes {
		suffixAligned[i] = string(alignBytes[i])
	}

	out := make(Alignment, len(suffixAligned))
	for i := 0; i < len(a); i++ {
		out[i] = frozen[i] + suffixAligned[i]
	}
	prefixGaps := strings.Repeat("-", len(frozen[0]))
	out[len(a)] = prefixGaps + suffixAligned[len(a)]

	return out
}

// SplitFrozenPrefix returns two things:
//  1. frozen[i]  — the left-hand part of row i that we will **never** touch
//  2. suffix     — the alignment consisting of just the last <window> columns
func SplitFrozenPrefix(a Alignment, window int) (frozen []string, suffix Alignment) {
	nCols := NumColumns(a)
	if nCols <= window {
		frozen = make([]string, len(a))
		suffix = a
		return
	}

	cut := nCols - window
	frozen = make([]string, len(a))
	suffix = make(Alignment, len(a))
	for i := range a {
		frozen[i] = a[i][:cut]
		suffix[i] = a[i][cut:]
	}
	return
}

func prepend(b []byte, c byte) []byte {
	b = append(b, 0)
	copy(b[1:], b[:len(b)-1])
	b[0] = c
	return b
}

func MakeSuffixProfile(a Alignment, w int) (Profile, int) {
	totalCols := NumColumns(a)
	if w >= totalCols {
		return MakeProfile(a), 0
	}
	start := totalCols - w
	suffix := make(Alignment, len(a))
	for i := range a {
		suffix[i] = a[i][start:]
	}
	return MakeProfile(suffix), start
}

func MakeProfile(a Alignment) Profile {
	AssertRectangular(a)
	numCols := NumColumns(a)
	profile := make(Profile, numCols)

	for col := 0; col < numCols; col++ {
		profile[col] = make(map[rune]float64)
		for row := 0; row < len(a); row++ {
			char := rune(a[row][col])
			if char != '-' {
				profile[col][char]++
			}
			if char == '-' && !LastGapSymbol(a[row], col) {
				profile[col]['-']++
			}
		}
	}

	return profile
}

func LastGapSymbol(row string, col int) bool {
	if col >= len(row) || row[col] != '-' {
		return false
	}
	for i := col + 1; i < len(row); i++ {
		if row[i] != '-' {
			return false
		}
	}
	return true
}

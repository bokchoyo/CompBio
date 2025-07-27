package main

// LocalAlignment takes two strings, along with match, mismatch, and gap scores.
// It returns a maximum score local alignment of the strings corresponding to these penalties.

func LocalAlignment(str1, str2 string, match, mismatch, gap float64) (Alignment, int, int, int, int) {
	var a Alignment
	var r, c int
	array := LocalScoreTable(str1, str2, match, mismatch, gap)
	max := 0.0
	end1 := 0
	end2 := 0

	for r := range array {
		for c := range array[r] {
			if array[r][c] > max {
				max = array[r][c]
				end1 = r
				end2 = c
			}
		}
	}

	r = end1
	c = end2

	for array[r][c] > 0 {
		n := array[r][c]
		switch n {
		case array[r][c-1] - gap:
			a[0] = "-" + a[0]
			a[1] = string(str2[c-1]) + a[1]
			c--
		case array[r-1][c] - gap:
			a[0] = string(str1[r-1]) + a[0]
			a[1] = "-" + a[1]
			r--
		case array[r-1][c-1] + match, array[r-1][c-1] - mismatch:
			a[0] = string(str1[r-1]) + a[0]
			a[1] = string(str2[c-1]) + a[1]
			r--
			c--
		default:
			panic("unexpected result")
		}
	}

	return a, r, end1, c, end2
}

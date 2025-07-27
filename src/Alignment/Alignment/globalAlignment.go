package main

type Alignment [2]string

// GlobalAlignment takes two strings, along with match, mismatch, and gap scores.
// It returns a maximum score global alignment of the strings corresponding to these penalties.
func GlobalAlignment(str1, str2 string, match, mismatch, gap float64) Alignment {
	var a Alignment
	array := GlobalScoreTable(str1, str2, match, mismatch, gap)
	r := len(str1)
	c := len(str2)
	for r > 0 || c > 0 {
		if r == 0 {
			a[0] = "-" + a[0]
			a[1] = string(str2[c-1]) + a[1]
			c--
		} else if c == 0 {
			a[0] = string(str1[r-1]) + a[0]
			a[1] = "-" + a[1]
			r--
		} else {
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

	}
	return a
}

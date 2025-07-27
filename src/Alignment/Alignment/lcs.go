package main

// LongestCommonSubsequence takes two strings as input.
// It returns a longest common subsequence of the two strings.
func LongestCommonSubsequence(str1, str2 string) string {
	if len(str1) == 0 || len(str2) == 0 {
		return ""
	}

	lcsMap := MakeLCSMap(str1, str2)
	r := len(str1)
	c := len(str2)
	lcs := ""
	for lcsMap[r][c] > 0 {
		match := 0
		if str1[r-1] == str2[c-1] {
			match = 1
		}
		d := lcsMap[r-1][c-1] + match
		u := lcsMap[r-1][c]
		l := lcsMap[r][c-1]

		if match == 1 {
			lcs += string(str1[r-1])
			r--
			c--
		} else if (d == u && u > l) || (d == l && l > u) || (d == u && d == l) {
			r--
			c--
		} else if u > l {
			r--
		} else if l > u {
			c--
		}
	}

	return Reverse(lcs)
}

func MakeLCSMap(str1, str2 string) [][]int {
	array := make([][]int, len(str1)+1)

	for i := range array {
		array[i] = make([]int, len(str2)+1)
	}

	for r := 1; r <= len(str1); r++ {
		for c := 1; c <= len(str2); c++ {
			if str1[r-1] == str2[c-1] {
				array[r][c] = Max3(array[r-1][c-1]+1, array[r][c-1], array[r-1][c])
			} else {
				array[r][c] = Max3(array[r-1][c-1], array[r][c-1], array[r-1][c])
			}
		}

	}

	return array
}

func Reverse(pattern string) string {
	rev := make([]byte, len(pattern))

	n := len(pattern)

	for i := range pattern {
		rev[i] = pattern[n-1-i]
	}

	return string(rev)
}

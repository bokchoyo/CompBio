package main

// EditDistance takes two strings as input. It returns the Levenshtein distance
// between the two strings; that is, the minimum number of substitutions, insertions, and deletions
// needed to transform one string into the other.
func EditDistance(str1, str2 string) int {
	array := make([][]int, len(str1)+1)

	for i := range array {
		array[i] = make([]int, len(str2)+1)
	}

    for i := 0; i <= len(str1); i++ {
        array[i][0] = i
    }

    for i := 0; i <= len(str2); i++ {
        array[0][i] = i
    }
    
	for r := 1; r <= len(str1); r++ {
		for c := 1; c <= len(str2); c++ {
			if str1[r-1] == str2[c-1] {
				array[r][c] = MinInts(array[r-1][c-1], array[r][c-1]+1, array[r-1][c]+1)
			} else {
				array[r][c] = MinInts(array[r-1][c-1]+1, array[r][c-1]+1, array[r-1][c]+1)
			}
		}
	}
    
    return array[len(str1)][len(str2)]

}

func MinInts(nums ...int) int {
	m := nums[0]

	for _, val := range nums {
		if val < m {
			m = val
		}
	}

	return m
}
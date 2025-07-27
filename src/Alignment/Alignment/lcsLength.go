package main

// LCSLength takes two strings as input. It returns the length of a longest common
// subsequence of the two strings.
func LCSLength(str1, str2 string) int {
	if len(str1) == 0 || len(str2) == 0 {
		return 0
	}

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

	return array[len(str1)][len(str2)]
}

func Max3(a, b, c int) int {
	if a > b && a > c {
		return a
	} else if b > c {
		return b
	} else {
		return c
	}
}

func recursiveSum(a int) int {
	if a == 1 {
		return 1
	}
	return a + recursiveSum(a-1)
}
func recursiveFactorial(a int) int {
	if a == 1 {
		return 1
	}
	return a * recursiveFactorial(a-1)
}
func recursiveFib(a int) int {
	if a == 1 || a == 0 {
		return 1
	}
	return recursiveFib(a-1) + recursiveFib(a-2)
}

/*

//redundant
func recursiveLongestCommonSubstring(dpVal int, i, j, s1, s2 string) int {

	if s1[i-1] == s1[j-1] {
		total++
	}

	return max(recursiveLongestCommonSubstring(dp[i][j-1], s1, s2), recursiveLongestCommonSubstring(dp[i-1][j], s1, s2))


} */
/*
func ManhattanProblem(rightWeights, downWeights
	][][]int) {
	dp := make([][]int, len(downWeights)+1)
	for i := range dp {
		dp[i] = make([]int, len(downWeights[])+1)
	}






}
*/
//break up into subroutines

// add subroutines
func LCSStorage(str1, str2 string) (string, string) {
	n, m := len(str1), len(str2)

	//Initialize matrices
	array := initializeIntMtx(n, m)
	result1 := initializeStrMtx(n, m)
	result2 := initializeStrMtx(n, m)

	//iterate through DP
	for r := 1; r <= n; r++ {
		for c := 1; c <= m; c++ {
			//if the strings are equal at indices, increase matches by 1, append char at index to both results
			if str1[r-1] == str2[c-1] {
				array[r][c] = array[r-1][c-1] + 1
				result1[r][c] = result1[r-1][c-1] + string(str1[r-1])
				result2[r][c] = result2[r-1][c-1] + string(str2[c-1])
			} else if array[r][c-1] > array[r-1][c] { //else if the strings are not equal and the dp left > dp above, append char to only result on top
				if array[r][c-1] == array[r-1][c] {
					array[r][c] = array[r][c-1]
					result1[r][c] = result1[r][c-1] + string(str2[c-1])
					result2[r][c] = result1[r][c-1] + string(str1[r-1])
				} else {
					array[r][c] = array[r][c-1]
					result1[r][c] = result1[r][c-1] + string(str2[c-1])
					result2[r][c] = result2[r][c-1] + "-"
				}
			} else { //else if dp above > dp left, append char to only one result on side
				array[r][c] = array[r-1][c]
				result1[r][c] = result1[r-1][c] + "-"
				result2[r][c] = result2[r-1][c] + string(str1[r-1])
			}
		}
	}
	return result1[n][m], result2[n][m]
}

func initializeIntMtx(n, m int) [][]int {
	array := make([][]int, n+1)
	for i := range array {
		array[i] = make([]int, m+1)
	}

	return array
}

func initializeStrMtx(n, m int) [][]string {
	array := make([][]string, n+1)
	for i := range array {
		array[i] = make([]string, m+1)
	}

	return array
}

func MaxInts(nums ...int) int {
	m := 0

	for i, val := range nums {
		if i == 0 || val > m {
			m = val
		}
	}

	return m
}

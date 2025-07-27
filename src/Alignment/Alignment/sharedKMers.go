package main

// CountSharedKmers takes two strings and an integer k. It returns the number of
// k-mers that are shared by the two strings.
func CountSharedKmers(text1, text2 string, k int) int {
	text1Map := map[string]int{}
	text2Map := map[string]int{}
	sum := 0

	for i := 0; i <= len(text1)-k; i++ {
		text := text1[i : i+k]
		text1Map[text]++
	}
	for i := 0; i <= len(text2)-k; i++ {
		text := text2[i : i+k]
		text2Map[text]++

	}
	for i, k := range text1Map {
		sum += Min(k, text2Map[i])
	}

	return sum
}

func Min(num1, num2 int) int {
	if num1 < num2 {
		return num1
	}
	return num2
}

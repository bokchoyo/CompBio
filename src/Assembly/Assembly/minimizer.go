package main

// Minimizer takes a string text and an integer k as input.
// It returns the k-mer of text that is lexicographically minimum.
func Minimizer(text string, k int) string {
	n := len(text)

	minimizer := text[0:k]
	for i := 0; i <= n-k; i++ {

		if minimizer > text[i:i+k] {
			minimizer = text[i : i+k]
		}

	}
	return minimizer
}
